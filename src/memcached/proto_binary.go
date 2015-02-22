// +build !generate

package memcached

import "os"
import "io"
import "fmt"
import "net"
import "encoding/binary"

const MAGIC_REQUEST = 0x80
const HEADER_BYTES = 24
const MAGIC_DEADBEEF = 0xdeadbeef

type binaryRequestHeader struct {
	magic           uint8
	opcode          uint8
	keyLength       uint16
	extrasLength    uint8
	dataType        uint8
	reserved        uint16
	totalBodyLength uint32
	opaque          uint32
	cas             uint64
}

func (header *binaryRequestHeader) parse(buf []byte) {
	header.magic = buf[0]
	header.opcode = buf[1]
	header.keyLength = binary.BigEndian.Uint16(buf[2:])
	header.extrasLength = buf[4]
	header.dataType = buf[5]
	header.reserved = binary.BigEndian.Uint16(buf[6:])
	header.totalBodyLength = binary.BigEndian.Uint32(buf[8:])
	header.opaque = binary.BigEndian.Uint32(buf[12:])
	header.cas = binary.BigEndian.Uint64(buf[16:])
}

type binaryProtocolHandler struct {
	parsers []parseRequestFunc
}

func (self binaryProtocolHandler) handleConnection(conn net.Conn, server *MemcachedServer) error {
	var header binaryRequestHeader
	var headerBuf [HEADER_BYTES]byte
	for {
		if _, err := io.ReadFull(conn, headerBuf[:]); err != nil {
			fmt.Fprintf(os.Stderr, "error: %+v\n", err)
			return err
		}

		header.parse(headerBuf[:])
		if header.magic != MAGIC_REQUEST {
			return fmt.Errorf("invalid magic %+v", header.magic)
		}

		bodyBuf := make([]byte, header.totalBodyLength)
		if _, err := io.ReadFull(conn, bodyBuf); err != nil {
			return err
		}

		offset := 0
		extras := bodyBuf[offset : offset+int(header.extrasLength)]
		offset += int(header.extrasLength)
		key := bodyBuf[offset : offset+int(header.keyLength)]
		offset += int(header.keyLength)
		value := bodyBuf[offset : HEADER_BYTES+header.totalBodyLength]

		if int(header.opcode) >= len(self.parsers) {
			return fmt.Errorf("invalid opcode")
		}
		command, err := self.parsers[header.opcode](&header, key, value, extras)
		if err != nil {
			return err
		}

		server.Call(command)
	}
	return nil
}

func newBinaryProtocolHandler() protocolHandler {
	return &binaryProtocolHandler{
		parsers: parseRequestFuncTable(),
	}
}
