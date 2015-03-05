// DO NOT EDIT THIS FILE
// This file is generated from commands.go.in

// +build !generate

package memcached

type addCommand struct {
	key []byte

	value []byte

	expire uint32

	cas uint64

	quiet bool
}

func Add(key []byte, value []byte) *addCommand {
	return &addCommand{
		key:   key,
		value: value,
	}
}

func (self *addCommand) WithExpire(expire uint32) *addCommand {
	self.expire = expire
	return self
}

func (self *addCommand) WithCas(cas uint64) *addCommand {
	self.cas = cas
	return self
}

func (self *addCommand) WithQuiet(quiet bool) *addCommand {
	self.quiet = quiet
	return self
}

type replaceCommand struct {
	key []byte

	value []byte

	expire uint32

	cas uint64

	quiet bool
}

func Replace(key []byte, value []byte) *replaceCommand {
	return &replaceCommand{
		key:   key,
		value: value,
	}
}

func (self *replaceCommand) WithExpire(expire uint32) *replaceCommand {
	self.expire = expire
	return self
}

func (self *replaceCommand) WithCas(cas uint64) *replaceCommand {
	self.cas = cas
	return self
}

func (self *replaceCommand) WithQuiet(quiet bool) *replaceCommand {
	self.quiet = quiet
	return self
}

type getCommand struct {
	key []byte

	quiet bool
}

func Get(key []byte) *getCommand {
	return &getCommand{
		key: key,
	}
}

func (self *getCommand) WithQuiet(quiet bool) *getCommand {
	self.quiet = quiet
	return self
}

type incrementCommand struct {
	key []byte

	expire uint32

	cas uint64

	quiet bool

	amount uint64

	initial uint64
}

func Increment(key []byte, amount uint64, initial uint64) *incrementCommand {
	return &incrementCommand{
		key: key,

		amount: amount, initial: initial,
	}
}

func (self *incrementCommand) WithExpire(expire uint32) *incrementCommand {
	self.expire = expire
	return self
}

func (self *incrementCommand) WithCas(cas uint64) *incrementCommand {
	self.cas = cas
	return self
}

func (self *incrementCommand) WithQuiet(quiet bool) *incrementCommand {
	self.quiet = quiet
	return self
}

type decrementCommand struct {
	key []byte

	expire uint32

	cas uint64

	quiet bool

	amount uint64

	initial uint64
}

func Decrement(key []byte, amount uint64, initial uint64) *decrementCommand {
	return &decrementCommand{
		key: key,

		amount: amount, initial: initial,
	}
}

func (self *decrementCommand) WithExpire(expire uint32) *decrementCommand {
	self.expire = expire
	return self
}

func (self *decrementCommand) WithCas(cas uint64) *decrementCommand {
	self.cas = cas
	return self
}

func (self *decrementCommand) WithQuiet(quiet bool) *decrementCommand {
	self.quiet = quiet
	return self
}

type setCommand struct {
	key []byte

	value []byte

	expire uint32

	cas uint64

	quiet bool
}

func Set(key []byte, value []byte) *setCommand {
	return &setCommand{
		key:   key,
		value: value,
	}
}

func (self *setCommand) WithExpire(expire uint32) *setCommand {
	self.expire = expire
	return self
}

func (self *setCommand) WithCas(cas uint64) *setCommand {
	self.cas = cas
	return self
}

func (self *setCommand) WithQuiet(quiet bool) *setCommand {
	self.quiet = quiet
	return self
}

type deleteCommand struct {
	key []byte

	cas uint64

	quiet bool
}

func Delete(key []byte) *deleteCommand {
	return &deleteCommand{
		key: key,
	}
}

func (self *deleteCommand) WithCas(cas uint64) *deleteCommand {
	self.cas = cas
	return self
}

func (self *deleteCommand) WithQuiet(quiet bool) *deleteCommand {
	self.quiet = quiet
	return self
}

type appendCommand struct {
	key []byte

	value []byte

	expire uint32

	cas uint64

	quiet bool
}

func Append(key []byte, value []byte) *appendCommand {
	return &appendCommand{
		key:   key,
		value: value,
	}
}

func (self *appendCommand) WithExpire(expire uint32) *appendCommand {
	self.expire = expire
	return self
}

func (self *appendCommand) WithCas(cas uint64) *appendCommand {
	self.cas = cas
	return self
}

func (self *appendCommand) WithQuiet(quiet bool) *appendCommand {
	self.quiet = quiet
	return self
}

type prependCommand struct {
	key []byte

	value []byte

	expire uint32

	cas uint64

	quiet bool
}

func Prepend(key []byte, value []byte) *prependCommand {
	return &prependCommand{
		key:   key,
		value: value,
	}
}

func (self *prependCommand) WithExpire(expire uint32) *prependCommand {
	self.expire = expire
	return self
}

func (self *prependCommand) WithCas(cas uint64) *prependCommand {
	self.cas = cas
	return self
}

func (self *prependCommand) WithQuiet(quiet bool) *prependCommand {
	self.quiet = quiet
	return self
}

type flushCommand struct {
	key []byte

	quiet bool
}

func Flush() *flushCommand {
	return &flushCommand{}
}

func (self *flushCommand) WithQuiet(quiet bool) *flushCommand {
	self.quiet = quiet
	return self
}

type nopCommand struct {
	key []byte
}

func nop() *nopCommand {
	return &nopCommand{}
}

type quitCommand struct {
	key []byte

	quiet bool
}

func quit() *quitCommand {
	return &quitCommand{}
}

func (self *quitCommand) WithQuiet(quiet bool) *quitCommand {
	self.quiet = quiet
	return self
}

type versionCommand struct {
	key []byte
}

func version() *versionCommand {
	return &versionCommand{}
}

type statCommand struct {
	key []byte
}

func stat() *statCommand {
	return &statCommand{}
}

type getWithKeyCommand struct {
	key []byte

	quiet bool
}

func GetWithKey(key []byte) *getWithKeyCommand {
	return &getWithKeyCommand{
		key: key,
	}
}

func (self *getWithKeyCommand) WithQuiet(quiet bool) *getWithKeyCommand {
	self.quiet = quiet
	return self
}

type serverCommand interface {
	execute(server *MemcachedServer)
}
