// DO NOT EDIT THIS FILE
// This file is generated from proto_binary_commands.go.in

// +build !generate

package memcached

import "os"
import "fmt"
import "encoding/binary"

type binaryRequestHandler func(cli Memcached, header *binaryRequestHeader, key []byte, value []byte, extras []byte) (rheader *binaryResponseHeader, rkey []byte, rvalue []byte, rextras []byte)

func binaryRequestHandlerTable() []binaryRequestHandler {
	return []binaryRequestHandler{
		handleGetRequest,         // 0
		handleSetRequest,         // 1
		handleAddRequest,         // 2
		handleReplaceRequest,     // 3
		handleDeleteRequest,      // 4
		handleIncrementRequest,   // 5
		handleDecrementRequest,   // 6
		handleQuitRequest,        // 7
		handleFlushRequest,       // 8
		handleGetQRequest,        // 9
		handleNopRequest,         // 10
		handleVersionRequest,     // 11
		handleGetWithKeyRequest,  // 12
		handleGetWithKeyQRequest, // 13
		handleAppendRequest,      // 14
		handlePrependRequest,     // 15
		handleStatRequest,        // 16
		handleSetQRequest,        // 17
		handleAddQRequest,        // 18
		handleReplaceQRequest,    // 19
		handleDeleteQRequest,     // 20
		handleIncrementQRequest,  // 21
		handleDecrementQRequest,  // 22
		handleQuitQRequest,       // 23
		handleFlushQRequest,      // 24
		handleAppendQRequest,     // 25
		handlePrependQRequest,    // 26
	}
}

func binaryErrorResponse(header *binaryRequestHeader, err *MemcachedError) (*binaryResponseHeader, []byte, []byte, []byte) {
	return &binaryResponseHeader{
		magic:  MAGIC_RESPONSE,
		opaque: header.opaque,
		opcode: header.opcode,
		status: err.ErrorCode(),
	}, nil, nil, nil
}

func handleGetRequest(cli Memcached, header *binaryRequestHeader, key []byte, value []byte, extras []byte) (*binaryResponseHeader, []byte, []byte, []byte) {
	if len(key) == 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("Get MUST have key"))
	}

	if len(value) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("Get MUST NOT have value"))
	}

	if len(extras) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("Get MUST NOT have extras"))
	}

	opts := &GetOptions{}
	rvalue, rflags, rcas, rerr := cli.Get(key, opts)

	if rerr != nil {
		return binaryErrorResponse(header, rerr)
	}

	var extralen int

	var extrabuf [4]byte

	binary.BigEndian.PutUint32(extrabuf[4:], rflags)

	extralen = 4

	rheader := &binaryResponseHeader{
		magic:  MAGIC_RESPONSE,
		opaque: header.opaque,
		opcode: header.opcode,

		status:          0,
		totalBodyLength: uint32(extralen + len(rvalue)),
		extrasLength:    uint8(extralen),
		cas:             rcas,
	}

	return rheader, nil, rvalue, nil

}

func handleSetRequest(cli Memcached, header *binaryRequestHeader, key []byte, value []byte, extras []byte) (*binaryResponseHeader, []byte, []byte, []byte) {
	if len(key) == 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("Set MUST have key"))
	}

	if len(value) == 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("Set MUST have value"))
	}

	if len(extras) != 8 {
		return binaryErrorResponse(header, newInvalidArgumentsError("Set MUST have extras of exactly 8 bytes"))
	}

	magic := binary.BigEndian.Uint32(extras[0:])
	if magic != MAGIC_DEADBEEF {
		fmt.Fprintf(os.Stderr, "Invalid magic for Set: %+v\n", magic)
		// return nil, fmt.Errorf("Invalid magic for Set: %+v", magic)
	}

	opts := &SetOptions{
		Expire: binary.BigEndian.Uint32(extras[4:]),
		Cas:    header.cas,
	}
	rcas, rerr := cli.Set(key, value, opts)

	if rerr != nil {
		return binaryErrorResponse(header, rerr)
	}

	var extralen int

	rheader := &binaryResponseHeader{
		magic:  MAGIC_RESPONSE,
		opaque: header.opaque,
		opcode: header.opcode,

		status:          0,
		totalBodyLength: uint32(extralen),
		extrasLength:    uint8(extralen),
		cas:             rcas,
	}

	return rheader, nil, nil, nil

}

func handleAddRequest(cli Memcached, header *binaryRequestHeader, key []byte, value []byte, extras []byte) (*binaryResponseHeader, []byte, []byte, []byte) {
	if len(key) == 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("Add MUST have key"))
	}

	if len(value) == 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("Add MUST have value"))
	}

	if len(extras) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("Add MUST NOT have extras"))
	}

	magic := binary.BigEndian.Uint32(extras[0:])
	if magic != MAGIC_DEADBEEF {
		fmt.Fprintf(os.Stderr, "Invalid magic for Add: %+v\n", magic)
		// return nil, fmt.Errorf("Invalid magic for Add: %+v", magic)
	}

	opts := &AddOptions{
		Expire: binary.BigEndian.Uint32(extras[4:]),
		Cas:    header.cas,
	}
	rcas, rerr := cli.Add(key, value, opts)

	if rerr != nil {
		return binaryErrorResponse(header, rerr)
	}

	var extralen int

	rheader := &binaryResponseHeader{
		magic:  MAGIC_RESPONSE,
		opaque: header.opaque,
		opcode: header.opcode,

		status:          0,
		totalBodyLength: uint32(extralen),
		extrasLength:    uint8(extralen),
		cas:             rcas,
	}

	return rheader, nil, nil, nil

}

func handleReplaceRequest(cli Memcached, header *binaryRequestHeader, key []byte, value []byte, extras []byte) (*binaryResponseHeader, []byte, []byte, []byte) {
	if len(key) == 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("Replace MUST have key"))
	}

	if len(value) == 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("Replace MUST have value"))
	}

	if len(extras) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("Replace MUST NOT have extras"))
	}

	magic := binary.BigEndian.Uint32(extras[0:])
	if magic != MAGIC_DEADBEEF {
		fmt.Fprintf(os.Stderr, "Invalid magic for Replace: %+v\n", magic)
		// return nil, fmt.Errorf("Invalid magic for Replace: %+v", magic)
	}

	opts := &ReplaceOptions{
		Expire: binary.BigEndian.Uint32(extras[4:]),
		Cas:    header.cas,
	}
	rcas, rerr := cli.Replace(key, value, opts)

	if rerr != nil {
		return binaryErrorResponse(header, rerr)
	}

	var extralen int

	rheader := &binaryResponseHeader{
		magic:  MAGIC_RESPONSE,
		opaque: header.opaque,
		opcode: header.opcode,

		status:          0,
		totalBodyLength: uint32(extralen),
		extrasLength:    uint8(extralen),
		cas:             rcas,
	}

	return rheader, nil, nil, nil

}

func handleDeleteRequest(cli Memcached, header *binaryRequestHeader, key []byte, value []byte, extras []byte) (*binaryResponseHeader, []byte, []byte, []byte) {
	if len(key) == 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("Delete MUST have key"))
	}

	if len(value) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("Delete MUST NOT have value"))
	}

	if len(extras) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("Delete MUST NOT have extras"))
	}

	opts := &DeleteOptions{

		Cas: header.cas,
	}
	rerr := cli.Delete(key, opts)

	if rerr != nil {
		return binaryErrorResponse(header, rerr)
	}

	var extralen int

	rheader := &binaryResponseHeader{
		magic:  MAGIC_RESPONSE,
		opaque: header.opaque,
		opcode: header.opcode,

		status:          0,
		totalBodyLength: uint32(extralen),
		extrasLength:    uint8(extralen),
	}

	return rheader, nil, nil, nil

}

func handleIncrementRequest(cli Memcached, header *binaryRequestHeader, key []byte, value []byte, extras []byte) (*binaryResponseHeader, []byte, []byte, []byte) {
	if len(key) == 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("Increment MUST have key"))
	}

	if len(value) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("Increment MUST NOT have value"))
	}

	if len(extras) != 20 {
		return binaryErrorResponse(header, newInvalidArgumentsError("Increment MUST have extras of exactly 20 bytes"))
	}

	amount := binary.BigEndian.Uint64(extras[0:])

	initial := binary.BigEndian.Uint64(extras[8:])

	opts := &IncrementOptions{
		Expire: binary.BigEndian.Uint32(extras[16:]),
		Cas:    header.cas,
	}
	rvalue, rcas, rerr := cli.Increment(key, amount, initial, opts)

	if rerr != nil {
		return binaryErrorResponse(header, rerr)
	}

	var extralen int

	rheader := &binaryResponseHeader{
		magic:  MAGIC_RESPONSE,
		opaque: header.opaque,
		opcode: header.opcode,

		status:          0,
		totalBodyLength: uint32(extralen + len(rvalue)),
		extrasLength:    uint8(extralen),
		cas:             rcas,
	}

	return rheader, nil, rvalue, nil

}

func handleDecrementRequest(cli Memcached, header *binaryRequestHeader, key []byte, value []byte, extras []byte) (*binaryResponseHeader, []byte, []byte, []byte) {
	if len(key) == 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("Decrement MUST have key"))
	}

	if len(value) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("Decrement MUST NOT have value"))
	}

	if len(extras) != 20 {
		return binaryErrorResponse(header, newInvalidArgumentsError("Decrement MUST have extras of exactly 20 bytes"))
	}

	amount := binary.BigEndian.Uint64(extras[0:])

	initial := binary.BigEndian.Uint64(extras[8:])

	opts := &DecrementOptions{
		Expire: binary.BigEndian.Uint32(extras[16:]),
		Cas:    header.cas,
	}
	rvalue, rcas, rerr := cli.Decrement(key, amount, initial, opts)

	if rerr != nil {
		return binaryErrorResponse(header, rerr)
	}

	var extralen int

	rheader := &binaryResponseHeader{
		magic:  MAGIC_RESPONSE,
		opaque: header.opaque,
		opcode: header.opcode,

		status:          0,
		totalBodyLength: uint32(extralen + len(rvalue)),
		extrasLength:    uint8(extralen),
		cas:             rcas,
	}

	return rheader, nil, rvalue, nil

}

func handleQuitRequest(cli Memcached, header *binaryRequestHeader, key []byte, value []byte, extras []byte) (*binaryResponseHeader, []byte, []byte, []byte) {
	if len(key) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("Quit MUST NOT have key"))
	}

	if len(value) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("Quit MUST NOT have value"))
	}

	if len(extras) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("Quit MUST NOT have extras"))
	}

	opts := &QuitOptions{}
	rerr := cli.Quit(opts)

	if rerr != nil {
		return binaryErrorResponse(header, rerr)
	}

	var extralen int

	rheader := &binaryResponseHeader{
		magic:  MAGIC_RESPONSE,
		opaque: header.opaque,
		opcode: header.opcode,

		status:          0,
		totalBodyLength: uint32(extralen),
		extrasLength:    uint8(extralen),
	}

	return rheader, nil, nil, nil

}

func handleFlushRequest(cli Memcached, header *binaryRequestHeader, key []byte, value []byte, extras []byte) (*binaryResponseHeader, []byte, []byte, []byte) {
	if len(key) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("Flush MUST NOT have key"))
	}

	if len(value) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("Flush MUST NOT have value"))
	}

	if len(extras) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("Flush MUST NOT have extras"))
	}

	opts := &FlushOptions{}
	rerr := cli.Flush(opts)

	if rerr != nil {
		return binaryErrorResponse(header, rerr)
	}

	var extralen int

	rheader := &binaryResponseHeader{
		magic:  MAGIC_RESPONSE,
		opaque: header.opaque,
		opcode: header.opcode,

		status:          0,
		totalBodyLength: uint32(extralen),
		extrasLength:    uint8(extralen),
	}

	return rheader, nil, nil, nil

}

func handleGetQRequest(cli Memcached, header *binaryRequestHeader, key []byte, value []byte, extras []byte) (*binaryResponseHeader, []byte, []byte, []byte) {
	if len(key) == 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("GetQ MUST have key"))
	}

	if len(value) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("GetQ MUST NOT have value"))
	}

	if len(extras) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("GetQ MUST NOT have extras"))
	}

	opts := &GetOptions{

		Quiet: true,
	}
	rvalue, rflags, rcas, rerr := cli.Get(key, opts)

	if rerr != nil {
		return binaryErrorResponse(header, rerr)
	}

	var extralen int

	var extrabuf [4]byte

	binary.BigEndian.PutUint32(extrabuf[4:], rflags)

	extralen = 4

	rheader := &binaryResponseHeader{
		magic:  MAGIC_RESPONSE,
		opaque: header.opaque,
		opcode: header.opcode,

		status:          0,
		totalBodyLength: uint32(extralen + len(rvalue)),
		extrasLength:    uint8(extralen),
		cas:             rcas,
	}

	return rheader, nil, rvalue, nil

}

func handleNopRequest(cli Memcached, header *binaryRequestHeader, key []byte, value []byte, extras []byte) (*binaryResponseHeader, []byte, []byte, []byte) {
	if len(key) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("Nop MUST NOT have key"))
	}

	if len(value) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("Nop MUST NOT have value"))
	}

	if len(extras) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("Nop MUST NOT have extras"))
	}

	opts := &NopOptions{}
	rerr := cli.Nop(opts)

	if rerr != nil {
		return binaryErrorResponse(header, rerr)
	}

	var extralen int

	rheader := &binaryResponseHeader{
		magic:  MAGIC_RESPONSE,
		opaque: header.opaque,
		opcode: header.opcode,

		status:          0,
		totalBodyLength: uint32(extralen),
		extrasLength:    uint8(extralen),
	}

	return rheader, nil, nil, nil

}

func handleVersionRequest(cli Memcached, header *binaryRequestHeader, key []byte, value []byte, extras []byte) (*binaryResponseHeader, []byte, []byte, []byte) {
	if len(key) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("Version MUST NOT have key"))
	}

	if len(value) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("Version MUST NOT have value"))
	}

	if len(extras) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("Version MUST NOT have extras"))
	}

	opts := &VersionOptions{}
	rerr := cli.Version(opts)

	if rerr != nil {
		return binaryErrorResponse(header, rerr)
	}

	var extralen int

	rheader := &binaryResponseHeader{
		magic:  MAGIC_RESPONSE,
		opaque: header.opaque,
		opcode: header.opcode,

		status:          0,
		totalBodyLength: uint32(extralen),
		extrasLength:    uint8(extralen),
	}

	return rheader, nil, nil, nil

}

func handleGetWithKeyRequest(cli Memcached, header *binaryRequestHeader, key []byte, value []byte, extras []byte) (*binaryResponseHeader, []byte, []byte, []byte) {
	if len(key) == 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("GetWithKey MUST have key"))
	}

	if len(value) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("GetWithKey MUST NOT have value"))
	}

	if len(extras) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("GetWithKey MUST NOT have extras"))
	}

	opts := &GetWithKeyOptions{}
	rkey, rvalue, rflags, rcas, rerr := cli.GetWithKey(key, opts)

	if rerr != nil {
		return binaryErrorResponse(header, rerr)
	}

	var extralen int

	var extrabuf [4]byte

	binary.BigEndian.PutUint32(extrabuf[4:], rflags)

	extralen = 4

	rheader := &binaryResponseHeader{
		magic:           MAGIC_RESPONSE,
		opaque:          header.opaque,
		opcode:          header.opcode,
		keyLength:       uint16(len(rkey)),
		status:          0,
		totalBodyLength: uint32(extralen + len(rkey) + len(rvalue)),
		extrasLength:    uint8(extralen),
		cas:             rcas,
	}

	return rheader, rkey, rvalue, nil

}

func handleGetWithKeyQRequest(cli Memcached, header *binaryRequestHeader, key []byte, value []byte, extras []byte) (*binaryResponseHeader, []byte, []byte, []byte) {
	if len(key) == 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("GetWithKeyQ MUST have key"))
	}

	if len(value) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("GetWithKeyQ MUST NOT have value"))
	}

	if len(extras) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("GetWithKeyQ MUST NOT have extras"))
	}

	opts := &GetWithKeyOptions{

		Quiet: true,
	}
	rkey, rvalue, rflags, rcas, rerr := cli.GetWithKey(key, opts)

	if rerr != nil {
		return binaryErrorResponse(header, rerr)
	}

	var extralen int

	var extrabuf [4]byte

	binary.BigEndian.PutUint32(extrabuf[4:], rflags)

	extralen = 4

	rheader := &binaryResponseHeader{
		magic:           MAGIC_RESPONSE,
		opaque:          header.opaque,
		opcode:          header.opcode,
		keyLength:       uint16(len(rkey)),
		status:          0,
		totalBodyLength: uint32(extralen + len(rkey) + len(rvalue)),
		extrasLength:    uint8(extralen),
		cas:             rcas,
	}

	return rheader, rkey, rvalue, nil

}

func handleAppendRequest(cli Memcached, header *binaryRequestHeader, key []byte, value []byte, extras []byte) (*binaryResponseHeader, []byte, []byte, []byte) {
	if len(key) == 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("Append MUST have key"))
	}

	if len(value) == 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("Append MUST have value"))
	}

	if len(extras) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("Append MUST NOT have extras"))
	}

	opts := &AppendOptions{

		Cas: header.cas,
	}
	rerr := cli.Append(key, value, opts)

	if rerr != nil {
		return binaryErrorResponse(header, rerr)
	}

	var extralen int

	rheader := &binaryResponseHeader{
		magic:  MAGIC_RESPONSE,
		opaque: header.opaque,
		opcode: header.opcode,

		status:          0,
		totalBodyLength: uint32(extralen),
		extrasLength:    uint8(extralen),
	}

	return rheader, nil, nil, nil

}

func handlePrependRequest(cli Memcached, header *binaryRequestHeader, key []byte, value []byte, extras []byte) (*binaryResponseHeader, []byte, []byte, []byte) {
	if len(key) == 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("Prepend MUST have key"))
	}

	if len(value) == 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("Prepend MUST have value"))
	}

	if len(extras) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("Prepend MUST NOT have extras"))
	}

	opts := &PrependOptions{

		Cas: header.cas,
	}
	rerr := cli.Prepend(key, value, opts)

	if rerr != nil {
		return binaryErrorResponse(header, rerr)
	}

	var extralen int

	rheader := &binaryResponseHeader{
		magic:  MAGIC_RESPONSE,
		opaque: header.opaque,
		opcode: header.opcode,

		status:          0,
		totalBodyLength: uint32(extralen),
		extrasLength:    uint8(extralen),
	}

	return rheader, nil, nil, nil

}

func handleStatRequest(cli Memcached, header *binaryRequestHeader, key []byte, value []byte, extras []byte) (*binaryResponseHeader, []byte, []byte, []byte) {
	if len(key) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("Stat MUST NOT have key"))
	}

	if len(value) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("Stat MUST NOT have value"))
	}

	if len(extras) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("Stat MUST NOT have extras"))
	}

	opts := &StatOptions{}
	rerr := cli.Stat(opts)

	if rerr != nil {
		return binaryErrorResponse(header, rerr)
	}

	var extralen int

	rheader := &binaryResponseHeader{
		magic:  MAGIC_RESPONSE,
		opaque: header.opaque,
		opcode: header.opcode,

		status:          0,
		totalBodyLength: uint32(extralen),
		extrasLength:    uint8(extralen),
	}

	return rheader, nil, nil, nil

}

func handleSetQRequest(cli Memcached, header *binaryRequestHeader, key []byte, value []byte, extras []byte) (*binaryResponseHeader, []byte, []byte, []byte) {
	if len(key) == 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("SetQ MUST have key"))
	}

	if len(value) == 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("SetQ MUST have value"))
	}

	if len(extras) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("SetQ MUST NOT have extras"))
	}

	opts := &SetOptions{
		Expire: binary.BigEndian.Uint32(extras[4:]),
		Cas:    header.cas,

		Quiet: true,
	}
	rcas, rerr := cli.Set(key, value, opts)

	if rerr != nil {
		return binaryErrorResponse(header, rerr)
	}

	var extralen int

	rheader := &binaryResponseHeader{
		magic:  MAGIC_RESPONSE,
		opaque: header.opaque,
		opcode: header.opcode,

		status:          0,
		totalBodyLength: uint32(extralen),
		extrasLength:    uint8(extralen),
		cas:             rcas,
	}

	return rheader, nil, nil, nil

}

func handleAddQRequest(cli Memcached, header *binaryRequestHeader, key []byte, value []byte, extras []byte) (*binaryResponseHeader, []byte, []byte, []byte) {
	if len(key) == 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("AddQ MUST have key"))
	}

	if len(value) == 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("AddQ MUST have value"))
	}

	if len(extras) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("AddQ MUST NOT have extras"))
	}

	magic := binary.BigEndian.Uint32(extras[0:])
	if magic != MAGIC_DEADBEEF {
		fmt.Fprintf(os.Stderr, "Invalid magic for AddQ: %+v\n", magic)
		// return nil, fmt.Errorf("Invalid magic for AddQ: %+v", magic)
	}

	opts := &AddOptions{
		Expire: binary.BigEndian.Uint32(extras[4:]),
		Cas:    header.cas,

		Quiet: true,
	}
	rcas, rerr := cli.Add(key, value, opts)

	if rerr != nil {
		return binaryErrorResponse(header, rerr)
	}

	var extralen int

	rheader := &binaryResponseHeader{
		magic:  MAGIC_RESPONSE,
		opaque: header.opaque,
		opcode: header.opcode,

		status:          0,
		totalBodyLength: uint32(extralen),
		extrasLength:    uint8(extralen),
		cas:             rcas,
	}

	return rheader, nil, nil, nil

}

func handleReplaceQRequest(cli Memcached, header *binaryRequestHeader, key []byte, value []byte, extras []byte) (*binaryResponseHeader, []byte, []byte, []byte) {
	if len(key) == 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("ReplaceQ MUST have key"))
	}

	if len(value) == 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("ReplaceQ MUST have value"))
	}

	if len(extras) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("ReplaceQ MUST NOT have extras"))
	}

	magic := binary.BigEndian.Uint32(extras[0:])
	if magic != MAGIC_DEADBEEF {
		fmt.Fprintf(os.Stderr, "Invalid magic for ReplaceQ: %+v\n", magic)
		// return nil, fmt.Errorf("Invalid magic for ReplaceQ: %+v", magic)
	}

	opts := &ReplaceOptions{
		Expire: binary.BigEndian.Uint32(extras[4:]),
		Cas:    header.cas,

		Quiet: true,
	}
	rcas, rerr := cli.Replace(key, value, opts)

	if rerr != nil {
		return binaryErrorResponse(header, rerr)
	}

	var extralen int

	rheader := &binaryResponseHeader{
		magic:  MAGIC_RESPONSE,
		opaque: header.opaque,
		opcode: header.opcode,

		status:          0,
		totalBodyLength: uint32(extralen),
		extrasLength:    uint8(extralen),
		cas:             rcas,
	}

	return rheader, nil, nil, nil

}

func handleDeleteQRequest(cli Memcached, header *binaryRequestHeader, key []byte, value []byte, extras []byte) (*binaryResponseHeader, []byte, []byte, []byte) {
	if len(key) == 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("DeleteQ MUST have key"))
	}

	if len(value) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("DeleteQ MUST NOT have value"))
	}

	if len(extras) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("DeleteQ MUST NOT have extras"))
	}

	opts := &DeleteOptions{

		Cas: header.cas,

		Quiet: true,
	}
	rerr := cli.Delete(key, opts)

	if rerr != nil {
		return binaryErrorResponse(header, rerr)
	}

	var extralen int

	rheader := &binaryResponseHeader{
		magic:  MAGIC_RESPONSE,
		opaque: header.opaque,
		opcode: header.opcode,

		status:          0,
		totalBodyLength: uint32(extralen),
		extrasLength:    uint8(extralen),
	}

	return rheader, nil, nil, nil

}

func handleIncrementQRequest(cli Memcached, header *binaryRequestHeader, key []byte, value []byte, extras []byte) (*binaryResponseHeader, []byte, []byte, []byte) {
	if len(key) == 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("IncrementQ MUST have key"))
	}

	if len(value) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("IncrementQ MUST NOT have value"))
	}

	if len(extras) != 20 {
		return binaryErrorResponse(header, newInvalidArgumentsError("IncrementQ MUST have extras of exactly 20 bytes"))
	}

	amount := binary.BigEndian.Uint64(extras[0:])

	initial := binary.BigEndian.Uint64(extras[8:])

	opts := &IncrementOptions{
		Expire: binary.BigEndian.Uint32(extras[16:]),
		Cas:    header.cas,

		Quiet: true,
	}
	rvalue, rcas, rerr := cli.Increment(key, amount, initial, opts)

	if rerr != nil {
		return binaryErrorResponse(header, rerr)
	}

	var extralen int

	rheader := &binaryResponseHeader{
		magic:  MAGIC_RESPONSE,
		opaque: header.opaque,
		opcode: header.opcode,

		status:          0,
		totalBodyLength: uint32(extralen + len(rvalue)),
		extrasLength:    uint8(extralen),
		cas:             rcas,
	}

	return rheader, nil, rvalue, nil

}

func handleDecrementQRequest(cli Memcached, header *binaryRequestHeader, key []byte, value []byte, extras []byte) (*binaryResponseHeader, []byte, []byte, []byte) {
	if len(key) == 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("DecrementQ MUST have key"))
	}

	if len(value) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("DecrementQ MUST NOT have value"))
	}

	if len(extras) != 20 {
		return binaryErrorResponse(header, newInvalidArgumentsError("DecrementQ MUST have extras of exactly 20 bytes"))
	}

	amount := binary.BigEndian.Uint64(extras[0:])

	initial := binary.BigEndian.Uint64(extras[8:])

	opts := &DecrementOptions{
		Expire: binary.BigEndian.Uint32(extras[16:]),
		Cas:    header.cas,

		Quiet: true,
	}
	rvalue, rcas, rerr := cli.Decrement(key, amount, initial, opts)

	if rerr != nil {
		return binaryErrorResponse(header, rerr)
	}

	var extralen int

	rheader := &binaryResponseHeader{
		magic:  MAGIC_RESPONSE,
		opaque: header.opaque,
		opcode: header.opcode,

		status:          0,
		totalBodyLength: uint32(extralen + len(rvalue)),
		extrasLength:    uint8(extralen),
		cas:             rcas,
	}

	return rheader, nil, rvalue, nil

}

func handleQuitQRequest(cli Memcached, header *binaryRequestHeader, key []byte, value []byte, extras []byte) (*binaryResponseHeader, []byte, []byte, []byte) {
	if len(key) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("QuitQ MUST NOT have key"))
	}

	if len(value) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("QuitQ MUST NOT have value"))
	}

	if len(extras) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("QuitQ MUST NOT have extras"))
	}

	opts := &QuitOptions{

		Quiet: true,
	}
	rerr := cli.Quit(opts)

	if rerr != nil {
		return binaryErrorResponse(header, rerr)
	}

	var extralen int

	rheader := &binaryResponseHeader{
		magic:  MAGIC_RESPONSE,
		opaque: header.opaque,
		opcode: header.opcode,

		status:          0,
		totalBodyLength: uint32(extralen),
		extrasLength:    uint8(extralen),
	}

	return rheader, nil, nil, nil

}

func handleFlushQRequest(cli Memcached, header *binaryRequestHeader, key []byte, value []byte, extras []byte) (*binaryResponseHeader, []byte, []byte, []byte) {
	if len(key) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("FlushQ MUST NOT have key"))
	}

	if len(value) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("FlushQ MUST NOT have value"))
	}

	if len(extras) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("FlushQ MUST NOT have extras"))
	}

	opts := &FlushOptions{

		Quiet: true,
	}
	rerr := cli.Flush(opts)

	if rerr != nil {
		return binaryErrorResponse(header, rerr)
	}

	var extralen int

	rheader := &binaryResponseHeader{
		magic:  MAGIC_RESPONSE,
		opaque: header.opaque,
		opcode: header.opcode,

		status:          0,
		totalBodyLength: uint32(extralen),
		extrasLength:    uint8(extralen),
	}

	return rheader, nil, nil, nil

}

func handleAppendQRequest(cli Memcached, header *binaryRequestHeader, key []byte, value []byte, extras []byte) (*binaryResponseHeader, []byte, []byte, []byte) {
	if len(key) == 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("AppendQ MUST have key"))
	}

	if len(value) == 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("AppendQ MUST have value"))
	}

	if len(extras) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("AppendQ MUST NOT have extras"))
	}

	opts := &AppendOptions{

		Cas: header.cas,

		Quiet: true,
	}
	rerr := cli.Append(key, value, opts)

	if rerr != nil {
		return binaryErrorResponse(header, rerr)
	}

	var extralen int

	rheader := &binaryResponseHeader{
		magic:  MAGIC_RESPONSE,
		opaque: header.opaque,
		opcode: header.opcode,

		status:          0,
		totalBodyLength: uint32(extralen),
		extrasLength:    uint8(extralen),
	}

	return rheader, nil, nil, nil

}

func handlePrependQRequest(cli Memcached, header *binaryRequestHeader, key []byte, value []byte, extras []byte) (*binaryResponseHeader, []byte, []byte, []byte) {
	if len(key) == 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("PrependQ MUST have key"))
	}

	if len(value) == 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("PrependQ MUST have value"))
	}

	if len(extras) > 0 {
		return binaryErrorResponse(header, newInvalidArgumentsError("PrependQ MUST NOT have extras"))
	}

	opts := &PrependOptions{

		Cas: header.cas,

		Quiet: true,
	}
	rerr := cli.Prepend(key, value, opts)

	if rerr != nil {
		return binaryErrorResponse(header, rerr)
	}

	var extralen int

	rheader := &binaryResponseHeader{
		magic:  MAGIC_RESPONSE,
		opaque: header.opaque,
		opcode: header.opcode,

		status:          0,
		totalBodyLength: uint32(extralen),
		extrasLength:    uint8(extralen),
	}

	return rheader, nil, nil, nil

}
