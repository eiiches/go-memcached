// +build !generate

package memcached

import "encoding/binary"

const BINARY_MAGIC_REQUEST = 0x80
const BINARY_MAGIC_RESPONSE = 0x81
const BINARY_HEADER_BYTES = 24

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

func (header *binaryResponseHeader) read(buf []byte) {
	header.magic = buf[0]
	header.opcode = buf[1]
	header.keyLength = binary.BigEndian.Uint16(buf[2:])
	header.extrasLength = buf[4]
	header.dataType = buf[5]
	header.status = binary.BigEndian.Uint16(buf[6:])
	header.totalBodyLength = binary.BigEndian.Uint32(buf[8:])
	header.opaque = binary.BigEndian.Uint32(buf[12:])
	header.cas = binary.BigEndian.Uint64(buf[16:])
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

func (header *binaryRequestHeader) read(buf []byte) {
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

func (header *binaryRequestHeader) write(buf []byte) {
	buf[0] = header.magic
	buf[1] = header.opcode
	binary.BigEndian.PutUint16(buf[2:], header.keyLength)
	buf[4] = header.extrasLength
	buf[5] = header.dataType
	binary.BigEndian.PutUint16(buf[6:], header.reserved)
	binary.BigEndian.PutUint32(buf[8:], header.totalBodyLength)
	binary.BigEndian.PutUint32(buf[12:], header.opaque)
	binary.BigEndian.PutUint64(buf[16:], header.cas)
}
