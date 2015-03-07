// +build !generate

package memcached

import "net"

type MemcachedClient struct {
	conn net.Conn
}

func NewMemcachedClient(proto string, laddr string) (*MemcachedClient, error) {
	conn, err := net.Dial(proto, laddr)
	if err != nil {
		return nil, err
	}

	return &MemcachedClient{
		conn: conn,
	}, nil
}

func (self *MemcachedClient) Close() error {
	return self.conn.Close()
}

func (self *MemcachedClient) Call(command clientCommand) error {
}
