package core

import unsafe "unsafe"
import "sync/atomic"

type AtomicBoolean uint32

func asUint32(v bool) uint32 {
	if v {
		return 1
	} else {
		return 0
	}
}

func (self *AtomicBoolean) Set(v bool) {
	atomic.StoreUint32((*uint32)(unsafe.Pointer(self)), asUint32(v))
}

func (self *AtomicBoolean) Get() bool {
	return atomic.LoadUint32((*uint32)(unsafe.Pointer(self))) != 0
}
