package core

import unsafe "unsafe"
import "bytes"
import "fmt"
import "sync"
import "sync/atomic"
import "hash/fnv"

type LruCache interface {
	// equivalent to SET operation
	Put(key []byte, value []byte, expire uint32) (oldval []byte, newcas uint64)

	// equivalent to COMPARE_AND_SET, COMPARE_AND_REPLACE
	CompareAndSet(key []byte, value []byte, cas uint64, expire uint32) (oldval []byte, newcas uint64)

	// equivalent to GET operation
	Get(key []byte) ([]byte, uint64) /* value and CAS token */

	// DELETE
	Remove(key []byte) []byte

	// COMPARE_AND_DELETE
	CompareAndRemove(key []byte, cas uint64) []byte

	// INCREMENT and DECREMENT
	Increment(key []byte, initial uint64, incr uint64, expire uint32) []byte
	Decrement(key []byte, initial uint64, decr uint64, expire uint32) []byte

	// REPLACE operation, which MUST fail if the item doesn't exist
	Replace(key []byte, value []byte, expire uint32) (success bool, newcas uint64)

	// equivalent to ADD operation, which MUST fail if the item already exists
	// COMPARE_AND_ADD always fails
	PutIfAbsent(key []byte, value []byte, expire uint32) []byte

	// APPEND
	Append(key []byte, value []byte, expire uint32) []byte

	// PREPEND
	Prepend(key []byte, value []byte, expire uint32) []byte

	Clear()
}

type entry_t struct {
	key    []byte
	hash   uint32
	next   *entry_t
	expire uint64
	value  []byte
	cas    uint64
}

type segment_t struct {
	lock sync.Mutex

	// final変数、ただしgolangではconstruct後のvisibilityが保証されていない TODO
	table []*entry_t

	// final変数、ただしgolangではconstruct後のvisibilityが保証されていない TODO
	mask uint32

	// このセグメントが持つエントリーの数
	// アクセスはlockで同期する必要がある
	count uint32
}

func (self *segment_t) put(key []byte, h uint32, value []byte, onlyIfAbsent bool) []byte {
	self.lock.Lock()
	defer self.lock.Unlock()

	index := h & self.mask
	ppent_ := (*unsafe.Pointer)(unsafe.Pointer(&self.table[index]))

	var oldvalue []byte

	head := (*entry_t)(atomic.LoadPointer(ppent_))
	for pent := head; ; {
		if pent != nil {
			if pent.hash == h && bytes.Compare(pent.key, key) == 0 {
				oldvalue = pent.value
				// これ、ConcurrentHashMapでただの代入で実行しているのは、javaのメモリモデルではポインタとかpritimive typeの代入は(doubleとlongは知らない)atomicだからだと思われる。
				// ただ、golangのメモリモデルはprimitive typeのアトミック性を保証していない。仕方がないので、golangではatomic.StorePointer()する。
				// というかそうしないとrace detectorに怒られた
				// pent.value = value
				atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&pent.value)), unsafe.Pointer(&value[0]))
				break
			}
			pent = pent.next
		} else {
			newcount := self.count + 1
			newhead := &entry_t{
				hash:  h,
				key:   key,
				value: value,
				next:  head,
			}
			// TODO: rehash

			// 上の.nextは、セグメントのheadにつなぐ前に必ず見えなければいけないので、
			// ConcurrentHashMapでは、StoreStoreバリアを実行してから書き込んでいる( putOrderedObject() )
			// ただし、golangにはatomic.LazyStorePointer()的なものはないので、パフォーマンス的に劣るものの、
			// releaseセマンティックを持つatomic.StorePointer()で代用する。
			atomic.StorePointer(ppent_, unsafe.Pointer(newhead))
			self.count = newcount
			break
		}
	}

	return oldvalue
}

// PutIfAbsent("hoge", "fuga").WithExpire(3).IfCasMatch(102)

type concurrentLruCache struct {
	segments       []*segment_t
	segmentMask    uint32
	segmentShift   uint32
	maxMemoryBytes uint64
}

func NewLruCache(maxMemoryBytes uint64) LruCache {
	concurrency := 24

	// use MSB (Most-Significat Bits) of the hash(key) to distribute keys into segments (i.e buckets).
	sshift := 0
	ssize := 1
	for ssize < concurrency {
		sshift++
		ssize = ssize << 1
	}

	self := &concurrentLruCache{
		segments:       make([]*segment_t, 32),
		segmentMask:    uint32(ssize - 1),
		segmentShift:   uint32(32 - sshift),
		maxMemoryBytes: maxMemoryBytes,
	}

	return self
}

func hash32(key []byte) uint32 {
	hasher := fnv.New32a()
	hasher.Write(key)
	return hasher.Sum32()
}

func (self *concurrentLruCache) Get(key []byte) ([]byte, uint64) {
	h := hash32(key)
	sindex := (h >> self.segmentShift) & self.segmentMask

	ppseg_ := (*unsafe.Pointer)(unsafe.Pointer(&self.segments[sindex]))
	pseg_ := atomic.LoadPointer(ppseg_)

	fmt.Printf("segments[hash(%v) >> %d & %d = %d] = %v\n", string(key), self.segmentShift, self.segmentMask, sindex, pseg_)
	if pseg_ == nil {
		return nil, 0
	}

	pseg := (*segment_t)(pseg_)
	ppent_ := (*unsafe.Pointer)(unsafe.Pointer(&pseg.table[h&pseg.mask]))

	for pent := (*entry_t)(atomic.LoadPointer(ppent_)); pent != nil; pent = pent.next {
		if pent.hash == h && bytes.Compare(key, pent.key) == 0 {
			return pent.value, pent.cas
		}
	}
	return nil, 0
}

func createSegment() *segment_t {
	capacity := 128
	return &segment_t{
		table: make([]*entry_t, capacity),
		mask:  uint32(capacity - 1),
	}
}

func ensureSegment(ppseg **segment_t) *segment_t {
	ppseg_ := (*unsafe.Pointer)(unsafe.Pointer(ppseg))
	pseg_ := atomic.LoadPointer(ppseg_)
	if pseg_ == nil {
		psegnew_ := unsafe.Pointer(createSegment())
		for {
			if pseg_ = atomic.LoadPointer(ppseg_); pseg_ != nil {
				break
			}
			if atomic.CompareAndSwapPointer(ppseg_, nil, psegnew_) {
				pseg_ = psegnew_
				break
			}
		}
	}
	return (*segment_t)(pseg_)
}

func (self *concurrentLruCache) Put(key []byte, value []byte, expire uint32) ([]byte, uint64) {
	h := hash32(key)
	sindex := (h >> self.segmentShift) & self.segmentMask

	ppseg := &self.segments[sindex]
	pseg_ := atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(ppseg)))
	if pseg_ == nil {
		pseg_ = unsafe.Pointer(ensureSegment(ppseg))
	}

	return (*segment_t)(pseg_).put(key, h, value, true), 0
}

func (self *concurrentLruCache) Remove(key []byte) []byte {
	// TODO
	return nil
}

func (self *concurrentLruCache) rehash() {
	// TODO
}

func (self *concurrentLruCache) PutIfAbsent(key []byte, value []byte, expire uint32) []byte {
	// TODO
	return nil
}

func (self *concurrentLruCache) Append(key []byte, value []byte, expire uint32) []byte {
	// TODO
	return nil
}

func (self *concurrentLruCache) Prepend(key []byte, value []byte, expire uint32) []byte {
	// TODO
	return nil
}

func (self *concurrentLruCache) Replace(key []byte, value []byte, expire uint32) (success bool, newcas uint64) {
	// TODO
	return false, 0
}

func (self *concurrentLruCache) Increment(key []byte, initial uint64, incr uint64, expire uint32) []byte {
	// TODO
	return nil
}

func (self *concurrentLruCache) Decrement(key []byte, initial uint64, decr uint64, expire uint32) []byte {
	// TODO
	return nil
}

func (self *concurrentLruCache) CompareAndRemove(key []byte, cas uint64) []byte {
	// TODO
	return nil
}

func (self *concurrentLruCache) CompareAndSet(key []byte, value []byte, cas uint64, expire uint32) ([]byte, uint64) {
	// TODO
	return nil, 0
}

func (self *concurrentLruCache) Clear() {
	// TODO
	return
}
