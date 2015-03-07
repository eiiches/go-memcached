// +build !generate

package memcached

import "net"

type protocolHandler interface {
	handleConnection(conn net.Conn, server *MemcachedServer) error
}
