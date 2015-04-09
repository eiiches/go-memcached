// +build !generate

package memcached

type MemcachedError struct {
	code uint32
	msg  string
}

func (self MemcachedError) Error() string {
	return self.msg
}

func (self MemcachedError) ErrorCode() uint32 {
	return self.code
}

/* TODO: エラーコードはバイナリプロトコルのものなのでどこか別の場所に移動 */
var (
	KeyNotFoundError        = &MemcachedError{code: 0x01, msg: "Key not found"}
	KeyExistsError          = &MemcachedError{code: 0x02, msg: "Key exists"}
	CasVersionMismatchError = &MemcachedError{code: 0x02, msg: "CAS version mismatch"}
	ValueTooLargeError      = &MemcachedError{code: 0x03, msg: "Value too large"}
	InvalidArgumentsError   = &MemcachedError{code: 0x04, msg: "Invalid arguments"}
	ItemNotStoredError      = &MemcachedError{code: 0x05, msg: "Item not stored"}
	IncrOnNonNumericError   = &MemcachedError{code: 0x06, msg: "Incr/Decr on non-numeric value"}
	UnknownCommandError     = &MemcachedError{code: 0x81, msg: "Unknown command"}
	OutOfMemoryError        = &MemcachedError{code: 0x82, msg: "Out of memory"}
)
