// +build !generate

package memcached

type serverCommand interface {
	execute(server *MemcachedServer)
}

func (self *MemcachedServer) Set(key []byte, value []byte, opts *SetOptions) (cas uint64, err *MemcachedError) {
	// fmt.Fprintf(os.Stderr, "PUT %v = %v [cas: %v]\n", string(key), string(value), opts.Cas)
	if opts != nil && opts.Cas != 0 {
		oldval, cas := self.cache.CompareAndSet(key, value, opts.Cas, opts.Expire)
		if oldval == nil {
			return 0, CasVersionMismatchError
		}
		return cas, nil
	} else {
		_, cas := self.cache.Put(key, value, opts.Expire)
		return cas, nil
	}
}

func (self *MemcachedServer) Add(key []byte, value []byte, opts *AddOptions) (cas uint64, err *MemcachedError) {
	return 0, nil
}

func (self *MemcachedServer) Replace(key []byte, value []byte, opts *ReplaceOptions) (cas uint64, err *MemcachedError) {
	if opts != nil && opts.Cas != 0 {
		oldval, cas := self.cache.CompareAndSet(key, value, opts.Cas, opts.Expire)
		if oldval == nil {
			return 0, CasVersionMismatchError
		}
		return cas, nil
	} else {
		success, cas := self.cache.Replace(key, value, opts.Expire)
		if !success {
			return 0, KeyNotFoundError
		}
		return cas, nil
	}
	return 0, nil
}

func (self *MemcachedServer) Get(key []byte, opts *GetOptions) (value []byte, flags uint32, cas uint64, err *MemcachedError) {
	rvalue, rcas := self.cache.Get(key)
	// fmt.Fprintf(os.Stderr, "GET %v = %v [cas: %v]\n", string(key), string(rvalue), rcas)
	return rvalue, 0, rcas, nil
}

func (self *MemcachedServer) Increment(key []byte, amount uint64, initial uint64, opts *IncrementOptions) (value []byte, cas uint64, err *MemcachedError) {
	if opts.Cas != 0 {
	} else {
	}
	return nil, 0, nil
}

func (self *MemcachedServer) Decrement(key []byte, amount uint64, initial uint64, opts *DecrementOptions) (value []byte, cas uint64, err *MemcachedError) {
	return nil, 0, nil
}

func (self *MemcachedServer) Delete(key []byte, opts *DeleteOptions) (err *MemcachedError) {
	if opts.Cas != 0 {
		self.cache.CompareAndRemove(key, opts.Cas)
	} else {
		self.cache.Remove(key)
	}
	return nil
}

func (self *MemcachedServer) Append(key []byte, value []byte, opts *AppendOptions) (err *MemcachedError) {
	return nil
}

func (self *MemcachedServer) Prepend(key []byte, value []byte, opts *PrependOptions) (err *MemcachedError) {
	return nil
}

func (self *MemcachedServer) Flush(opts *FlushOptions) (err *MemcachedError) {
	self.cache.Clear()
	return nil
}

func (self *MemcachedServer) Nop(opts *NopOptions) (err *MemcachedError) {
	return nil
}

func (self *MemcachedServer) Quit(opts *QuitOptions) (err *MemcachedError) {
	return nil
}

func (self *MemcachedServer) Stat(opts *StatOptions) (err *MemcachedError) {
	return nil
}

func (self *MemcachedServer) Version(opts *VersionOptions) (err *MemcachedError) {
	return nil
}

func (self *MemcachedServer) GetWithKey(key []byte, opts *GetWithKeyOptions) ([]byte, []byte, uint32, uint64, *MemcachedError) {
	rvalue, rcas := self.cache.Get(key)
	// fmt.Fprintf(os.Stderr, "GETK %v = %v [cas: %v]\n", string(key), string(rvalue), cas)
	return key, rvalue, 0, rcas, nil
}
