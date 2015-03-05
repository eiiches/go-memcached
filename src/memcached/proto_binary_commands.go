// DO NOT EDIT THIS FILE
// This file is generated from proto_binary_commands.go.in

// +build !generate

package memcached

import "fmt"
import "encoding/binary"

type parseRequestFunc func(header *binaryRequestHeader, key []byte, value []byte, extras []byte) (serverCommand, error)

func parseRequestFuncTable() []parseRequestFunc {
	return []parseRequestFunc{
		parseGetRequest,         // 0
		parseSetRequest,         // 1
		parseAddRequest,         // 2
		parseReplaceRequest,     // 3
		parseDeleteRequest,      // 4
		parseIncrementRequest,   // 5
		parseDecrementRequest,   // 6
		parseQuitRequest,        // 7
		parseFlushRequest,       // 8
		parseGetQRequest,        // 9
		parseNopRequest,         // 10
		parseVersionRequest,     // 11
		parseGetWithKeyRequest,  // 12
		parseGetWithKeyQRequest, // 13
		parseAppendRequest,      // 14
		parsePrependRequest,     // 15
		parseStatRequest,        // 16
		parseSetQRequest,        // 17
		parseAddQRequest,        // 18
		parseReplaceQRequest,    // 19
		parseDeleteQRequest,     // 20
		parseIncrementQRequest,  // 21
		parseDecrementQRequest,  // 22
		parseQuitQRequest,       // 23
		parseFlushQRequest,      // 24
		parseAppendQRequest,     // 25
		parsePrependQRequest,    // 26
	}
}

func parseGetRequest(header *binaryRequestHeader, key []byte, value []byte, extras []byte) (serverCommand, error) {
	if len(key) == 0 {
		return nil, fmt.Errorf("Get MUST have key")
	}

	if len(value) > 0 {
		return nil, fmt.Errorf("Get MUST NOT have value")
	}

	if len(extras) > 0 {
		return nil, fmt.Errorf("Get MUST NOT have extras")
	}

	command := Get(key)

	return command, nil
}

func parseSetRequest(header *binaryRequestHeader, key []byte, value []byte, extras []byte) (serverCommand, error) {
	if len(key) == 0 {
		return nil, fmt.Errorf("Set MUST have key")
	}

	if len(value) == 0 {
		return nil, fmt.Errorf("Set MUST have value")
	}

	if len(extras) > 0 {
		return nil, fmt.Errorf("Set MUST NOT have extras")
	}

	if binary.BigEndian.Uint32(extras[0:]) != MAGIC_DEADBEEF {
		return nil, fmt.Errorf("Invalid magic for Set")
	}

	command := Set(key, value)

	command.WithExpire(binary.BigEndian.Uint32(extras[4:]))

	command.WithCas(header.cas)

	return command, nil
}

func parseAddRequest(header *binaryRequestHeader, key []byte, value []byte, extras []byte) (serverCommand, error) {
	if len(key) == 0 {
		return nil, fmt.Errorf("Add MUST have key")
	}

	if len(value) == 0 {
		return nil, fmt.Errorf("Add MUST have value")
	}

	if len(extras) > 0 {
		return nil, fmt.Errorf("Add MUST NOT have extras")
	}

	if binary.BigEndian.Uint32(extras[0:]) != MAGIC_DEADBEEF {
		return nil, fmt.Errorf("Invalid magic for Add")
	}

	command := Add(key, value)

	command.WithExpire(binary.BigEndian.Uint32(extras[4:]))

	command.WithCas(header.cas)

	return command, nil
}

func parseReplaceRequest(header *binaryRequestHeader, key []byte, value []byte, extras []byte) (serverCommand, error) {
	if len(key) == 0 {
		return nil, fmt.Errorf("Replace MUST have key")
	}

	if len(value) == 0 {
		return nil, fmt.Errorf("Replace MUST have value")
	}

	if len(extras) > 0 {
		return nil, fmt.Errorf("Replace MUST NOT have extras")
	}

	if binary.BigEndian.Uint32(extras[0:]) != MAGIC_DEADBEEF {
		return nil, fmt.Errorf("Invalid magic for Replace")
	}

	command := Replace(key, value)

	command.WithExpire(binary.BigEndian.Uint32(extras[4:]))

	command.WithCas(header.cas)

	return command, nil
}

func parseDeleteRequest(header *binaryRequestHeader, key []byte, value []byte, extras []byte) (serverCommand, error) {
	if len(key) == 0 {
		return nil, fmt.Errorf("Delete MUST have key")
	}

	if len(value) > 0 {
		return nil, fmt.Errorf("Delete MUST NOT have value")
	}

	if len(extras) > 0 {
		return nil, fmt.Errorf("Delete MUST NOT have extras")
	}

	command := Delete(key)

	command.WithCas(header.cas)

	return command, nil
}

func parseIncrementRequest(header *binaryRequestHeader, key []byte, value []byte, extras []byte) (serverCommand, error) {
	if len(key) == 0 {
		return nil, fmt.Errorf("Increment MUST have key")
	}

	if len(value) > 0 {
		return nil, fmt.Errorf("Increment MUST NOT have value")
	}

	if len(extras) != 20 {
		return nil, fmt.Errorf("Increment MUST have extras of exactly 20 bytes")
	}

	amount := binary.BigEndian.Uint64(extras[0:])

	initial := binary.BigEndian.Uint64(extras[8:])

	command := Increment(key, amount, initial)

	command.WithExpire(binary.BigEndian.Uint32(extras[16:]))

	command.WithCas(header.cas)

	return command, nil
}

func parseDecrementRequest(header *binaryRequestHeader, key []byte, value []byte, extras []byte) (serverCommand, error) {
	if len(key) == 0 {
		return nil, fmt.Errorf("Decrement MUST have key")
	}

	if len(value) > 0 {
		return nil, fmt.Errorf("Decrement MUST NOT have value")
	}

	if len(extras) != 20 {
		return nil, fmt.Errorf("Decrement MUST have extras of exactly 20 bytes")
	}

	amount := binary.BigEndian.Uint64(extras[0:])

	initial := binary.BigEndian.Uint64(extras[8:])

	command := Decrement(key, amount, initial)

	command.WithExpire(binary.BigEndian.Uint32(extras[16:]))

	command.WithCas(header.cas)

	return command, nil
}

func parseQuitRequest(header *binaryRequestHeader, key []byte, value []byte, extras []byte) (serverCommand, error) {
	if len(key) > 0 {
		return nil, fmt.Errorf("Quit MUST NOT have key")
	}

	if len(value) > 0 {
		return nil, fmt.Errorf("Quit MUST NOT have value")
	}

	if len(extras) > 0 {
		return nil, fmt.Errorf("Quit MUST NOT have extras")
	}

	command := quit()

	return command, nil
}

func parseFlushRequest(header *binaryRequestHeader, key []byte, value []byte, extras []byte) (serverCommand, error) {
	if len(key) > 0 {
		return nil, fmt.Errorf("Flush MUST NOT have key")
	}

	if len(value) > 0 {
		return nil, fmt.Errorf("Flush MUST NOT have value")
	}

	if len(extras) > 0 {
		return nil, fmt.Errorf("Flush MUST NOT have extras")
	}

	command := Flush()

	return command, nil
}

func parseGetQRequest(header *binaryRequestHeader, key []byte, value []byte, extras []byte) (serverCommand, error) {
	if len(key) == 0 {
		return nil, fmt.Errorf("GetQ MUST have key")
	}

	if len(value) > 0 {
		return nil, fmt.Errorf("GetQ MUST NOT have value")
	}

	if len(extras) > 0 {
		return nil, fmt.Errorf("GetQ MUST NOT have extras")
	}

	command := Get(key)

	command.WithQuiet(true)

	return command, nil
}

func parseNopRequest(header *binaryRequestHeader, key []byte, value []byte, extras []byte) (serverCommand, error) {
	if len(key) > 0 {
		return nil, fmt.Errorf("Nop MUST NOT have key")
	}

	if len(value) > 0 {
		return nil, fmt.Errorf("Nop MUST NOT have value")
	}

	if len(extras) > 0 {
		return nil, fmt.Errorf("Nop MUST NOT have extras")
	}

	command := nop()

	return command, nil
}

func parseVersionRequest(header *binaryRequestHeader, key []byte, value []byte, extras []byte) (serverCommand, error) {
	if len(key) > 0 {
		return nil, fmt.Errorf("Version MUST NOT have key")
	}

	if len(value) > 0 {
		return nil, fmt.Errorf("Version MUST NOT have value")
	}

	if len(extras) > 0 {
		return nil, fmt.Errorf("Version MUST NOT have extras")
	}

	command := version()

	return command, nil
}

func parseGetWithKeyRequest(header *binaryRequestHeader, key []byte, value []byte, extras []byte) (serverCommand, error) {
	if len(key) == 0 {
		return nil, fmt.Errorf("GetWithKey MUST have key")
	}

	if len(value) > 0 {
		return nil, fmt.Errorf("GetWithKey MUST NOT have value")
	}

	if len(extras) > 0 {
		return nil, fmt.Errorf("GetWithKey MUST NOT have extras")
	}

	command := GetWithKey(key)

	return command, nil
}

func parseGetWithKeyQRequest(header *binaryRequestHeader, key []byte, value []byte, extras []byte) (serverCommand, error) {
	if len(key) == 0 {
		return nil, fmt.Errorf("GetWithKeyQ MUST have key")
	}

	if len(value) > 0 {
		return nil, fmt.Errorf("GetWithKeyQ MUST NOT have value")
	}

	if len(extras) > 0 {
		return nil, fmt.Errorf("GetWithKeyQ MUST NOT have extras")
	}

	command := GetWithKey(key)

	command.WithQuiet(true)

	return command, nil
}

func parseAppendRequest(header *binaryRequestHeader, key []byte, value []byte, extras []byte) (serverCommand, error) {
	if len(key) == 0 {
		return nil, fmt.Errorf("Append MUST have key")
	}

	if len(value) == 0 {
		return nil, fmt.Errorf("Append MUST have value")
	}

	if len(extras) > 0 {
		return nil, fmt.Errorf("Append MUST NOT have extras")
	}

	command := Append(key, value)

	return command, nil
}

func parsePrependRequest(header *binaryRequestHeader, key []byte, value []byte, extras []byte) (serverCommand, error) {
	if len(key) == 0 {
		return nil, fmt.Errorf("Prepend MUST have key")
	}

	if len(value) == 0 {
		return nil, fmt.Errorf("Prepend MUST have value")
	}

	if len(extras) > 0 {
		return nil, fmt.Errorf("Prepend MUST NOT have extras")
	}

	command := Prepend(key, value)

	return command, nil
}

func parseStatRequest(header *binaryRequestHeader, key []byte, value []byte, extras []byte) (serverCommand, error) {
	if len(key) > 0 {
		return nil, fmt.Errorf("Stat MUST NOT have key")
	}

	if len(value) > 0 {
		return nil, fmt.Errorf("Stat MUST NOT have value")
	}

	if len(extras) > 0 {
		return nil, fmt.Errorf("Stat MUST NOT have extras")
	}

	command := stat()

	return command, nil
}

func parseSetQRequest(header *binaryRequestHeader, key []byte, value []byte, extras []byte) (serverCommand, error) {
	if len(key) == 0 {
		return nil, fmt.Errorf("SetQ MUST have key")
	}

	if len(value) == 0 {
		return nil, fmt.Errorf("SetQ MUST have value")
	}

	if len(extras) > 0 {
		return nil, fmt.Errorf("SetQ MUST NOT have extras")
	}

	command := Set(key, value)

	command.WithQuiet(true)

	return command, nil
}

func parseAddQRequest(header *binaryRequestHeader, key []byte, value []byte, extras []byte) (serverCommand, error) {
	if len(key) == 0 {
		return nil, fmt.Errorf("AddQ MUST have key")
	}

	if len(value) == 0 {
		return nil, fmt.Errorf("AddQ MUST have value")
	}

	if len(extras) > 0 {
		return nil, fmt.Errorf("AddQ MUST NOT have extras")
	}

	if binary.BigEndian.Uint32(extras[0:]) != MAGIC_DEADBEEF {
		return nil, fmt.Errorf("Invalid magic for AddQ")
	}

	command := Add(key, value)

	command.WithExpire(binary.BigEndian.Uint32(extras[4:]))

	command.WithCas(header.cas)

	command.WithQuiet(true)

	return command, nil
}

func parseReplaceQRequest(header *binaryRequestHeader, key []byte, value []byte, extras []byte) (serverCommand, error) {
	if len(key) == 0 {
		return nil, fmt.Errorf("ReplaceQ MUST have key")
	}

	if len(value) == 0 {
		return nil, fmt.Errorf("ReplaceQ MUST have value")
	}

	if len(extras) > 0 {
		return nil, fmt.Errorf("ReplaceQ MUST NOT have extras")
	}

	if binary.BigEndian.Uint32(extras[0:]) != MAGIC_DEADBEEF {
		return nil, fmt.Errorf("Invalid magic for ReplaceQ")
	}

	command := Replace(key, value)

	command.WithExpire(binary.BigEndian.Uint32(extras[4:]))

	command.WithCas(header.cas)

	command.WithQuiet(true)

	return command, nil
}

func parseDeleteQRequest(header *binaryRequestHeader, key []byte, value []byte, extras []byte) (serverCommand, error) {
	if len(key) == 0 {
		return nil, fmt.Errorf("DeleteQ MUST have key")
	}

	if len(value) > 0 {
		return nil, fmt.Errorf("DeleteQ MUST NOT have value")
	}

	if len(extras) > 0 {
		return nil, fmt.Errorf("DeleteQ MUST NOT have extras")
	}

	command := Delete(key)

	command.WithQuiet(true)

	return command, nil
}

func parseIncrementQRequest(header *binaryRequestHeader, key []byte, value []byte, extras []byte) (serverCommand, error) {
	if len(key) == 0 {
		return nil, fmt.Errorf("IncrementQ MUST have key")
	}

	if len(value) > 0 {
		return nil, fmt.Errorf("IncrementQ MUST NOT have value")
	}

	if len(extras) != 20 {
		return nil, fmt.Errorf("IncrementQ MUST have extras of exactly 20 bytes")
	}

	amount := binary.BigEndian.Uint64(extras[0:])

	initial := binary.BigEndian.Uint64(extras[8:])

	command := Increment(key, amount, initial)

	command.WithCas(header.cas)

	command.WithQuiet(true)

	return command, nil
}

func parseDecrementQRequest(header *binaryRequestHeader, key []byte, value []byte, extras []byte) (serverCommand, error) {
	if len(key) == 0 {
		return nil, fmt.Errorf("DecrementQ MUST have key")
	}

	if len(value) > 0 {
		return nil, fmt.Errorf("DecrementQ MUST NOT have value")
	}

	if len(extras) != 20 {
		return nil, fmt.Errorf("DecrementQ MUST have extras of exactly 20 bytes")
	}

	amount := binary.BigEndian.Uint64(extras[0:])

	initial := binary.BigEndian.Uint64(extras[8:])

	command := Decrement(key, amount, initial)

	command.WithCas(header.cas)

	command.WithQuiet(true)

	return command, nil
}

func parseQuitQRequest(header *binaryRequestHeader, key []byte, value []byte, extras []byte) (serverCommand, error) {
	if len(key) > 0 {
		return nil, fmt.Errorf("QuitQ MUST NOT have key")
	}

	if len(value) > 0 {
		return nil, fmt.Errorf("QuitQ MUST NOT have value")
	}

	if len(extras) > 0 {
		return nil, fmt.Errorf("QuitQ MUST NOT have extras")
	}

	command := quit()

	command.WithQuiet(true)

	return command, nil
}

func parseFlushQRequest(header *binaryRequestHeader, key []byte, value []byte, extras []byte) (serverCommand, error) {
	if len(key) > 0 {
		return nil, fmt.Errorf("FlushQ MUST NOT have key")
	}

	if len(value) > 0 {
		return nil, fmt.Errorf("FlushQ MUST NOT have value")
	}

	if len(extras) > 0 {
		return nil, fmt.Errorf("FlushQ MUST NOT have extras")
	}

	command := Flush()

	command.WithQuiet(true)

	return command, nil
}

func parseAppendQRequest(header *binaryRequestHeader, key []byte, value []byte, extras []byte) (serverCommand, error) {
	if len(key) == 0 {
		return nil, fmt.Errorf("AppendQ MUST have key")
	}

	if len(value) == 0 {
		return nil, fmt.Errorf("AppendQ MUST have value")
	}

	if len(extras) > 0 {
		return nil, fmt.Errorf("AppendQ MUST NOT have extras")
	}

	command := Append(key, value)

	command.WithQuiet(true)

	return command, nil
}

func parsePrependQRequest(header *binaryRequestHeader, key []byte, value []byte, extras []byte) (serverCommand, error) {
	if len(key) == 0 {
		return nil, fmt.Errorf("PrependQ MUST have key")
	}

	if len(value) == 0 {
		return nil, fmt.Errorf("PrependQ MUST have value")
	}

	if len(extras) > 0 {
		return nil, fmt.Errorf("PrependQ MUST NOT have extras")
	}

	command := Prepend(key, value)

	command.WithQuiet(true)

	return command, nil
}
