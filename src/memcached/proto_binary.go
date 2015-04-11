// +build !generate

package memcached

import "os"
import "io"
import "fmt"
import "net"
import "encoding/binary"

const MAGIC_REQUEST = 0x80
const MAGIC_RESPONSE = 0x81
const HEADER_BYTES = 24
const MAGIC_DEADBEEF = 0xdeadbeef

type binaryResponseHeader struct {
	magic           uint8
	opcode          uint8
	keyLength       uint16
	extrasLength    uint8
	dataType        uint8
	status          uint16
	totalBodyLength uint32
	opaque          uint32
	cas             uint64
}

func (header *binaryResponseHeader) write(buf []byte) {
	buf[0] = header.magic
	buf[1] = header.opcode
	binary.BigEndian.PutUint16(buf[2:], header.keyLength)
	buf[4] = header.extrasLength
	buf[5] = header.dataType
	binary.BigEndian.PutUint16(buf[6:], header.status)
	binary.BigEndian.PutUint32(buf[8:], header.totalBodyLength)
	binary.BigEndian.PutUint32(buf[12:], header.opaque)
	binary.BigEndian.PutUint64(buf[16:], header.cas)
}

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
	handlers []binaryRequestHandler
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

		fmt.Fprintf(os.Stderr, "header: %+v\n", headerBuf)

		bodyBuf := make([]byte, header.totalBodyLength)
		if _, err := io.ReadFull(conn, bodyBuf); err != nil {
			return err
		}

		fmt.Fprintf(os.Stderr, "body: %+v\n", bodyBuf)

		offset := 0
		extras := bodyBuf[offset : offset+int(header.extrasLength)]
		offset += int(header.extrasLength)
		key := bodyBuf[offset : offset+int(header.keyLength)]
		offset += int(header.keyLength)
		value := bodyBuf[offset:header.totalBodyLength]

		if int(header.opcode) >= len(self.handlers) {
			return fmt.Errorf("invalid opcode")
		}
		rheader, rkey, rvalue, rextras := self.handlers[header.opcode](server, &header, key, value, extras)
		fmt.Fprintf(os.Stderr, "response: header = %+v, key = %+v, value = %+v, extras = %+v\n", rheader, rkey, rvalue, rextras)

		var rbuf [HEADER_BYTES]byte
		rheader.write(rbuf[:])

		conn.Write(rbuf[:])
		if rkey != nil {
			conn.Write(rkey)
		}
		if rvalue != nil {
			conn.Write(rvalue)
		}
		if rextras != nil {
			conn.Write(rextras)
		}
	}
	return nil
}

func newBinaryProtocolHandler() protocolHandler {
	return &binaryProtocolHandler{
		handlers: binaryRequestHandlerTable(),
	}
}
