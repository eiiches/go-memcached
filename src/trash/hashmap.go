package concurrent

import unsafe "unsafe"
import "bytes"
import "fmt"
import "sync"
import "sync/atomic"
import "hash/fnv"

// import "sync/atomic"

type entry_t struct {
	hash  uint32
	key   []byte
	value []byte
	next  *entry_t
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

type ConcurrentHashMap struct {
	segments     []*segment_t
	segmentMask  uint32
	segmentShift uint32
}

func NewConcurrentHashMap() *ConcurrentHashMap {
	concurrency := 24

	// use MSB (Most-Significat Bits) of the hash(key) to distribute keys into segments (i.e buckets).
	sshift := 0
	ssize := 1
	for ssize < concurrency {
		sshift++
		ssize = ssize << 1
	}

	self := &ConcurrentHashMap{
		segments:     make([]*segment_t, 32),
		segmentMask:  uint32(ssize - 1),
		segmentShift: uint32(32 - sshift),
	}

	return self
}

func hash32(key []byte) uint32 {
	hasher := fnv.New32a()
	hasher.Write(key)
	return hasher.Sum32()
}

func (self *ConcurrentHashMap) Get(key []byte) []byte {
	h := hash32(key)
	sindex := (h >> self.segmentShift) & self.segmentMask

	ppseg_ := (*unsafe.Pointer)(unsafe.Pointer(&self.segments[sindex]))
	pseg_ := atomic.LoadPointer(ppseg_)

	fmt.Printf("segments[hash(%v) >> %d & %d = %d] = %v\n", string(key), self.segmentShift, self.segmentMask, sindex, pseg_)
	if pseg_ == nil {
		return nil
	}

	pseg := (*segment_t)(pseg_)
	ppent_ := (*unsafe.Pointer)(unsafe.Pointer(&pseg.table[h&pseg.mask]))

	for pent := (*entry_t)(atomic.LoadPointer(ppent_)); pent != nil; pent = pent.next {
		if pent.hash == h && bytes.Compare(key, pent.key) == 0 {
			return pent.value
		}
	}
	return nil
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

func (self *ConcurrentHashMap) Put(key []byte, value []byte) []byte {
	h := hash32(key)
	sindex := (h >> self.segmentShift) & self.segmentMask

	ppseg := &self.segments[sindex]
	pseg_ := atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(ppseg)))
	if pseg_ == nil {
		pseg_ = unsafe.Pointer(ensureSegment(ppseg))
	}

	return (*segment_t)(pseg_).put(key, h, value, true)
}

func (self *ConcurrentHashMap) Remove(key []byte, value []byte) []byte {
	// TODO
	return nil
}

func (self *ConcurrentHashMap) rehash() {
	// TODO
}
