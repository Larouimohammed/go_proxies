package proxy

import "net"

type Proxier interface {
	Listent_And_Accept(port string) error
	Handle_conn(conn net.Conn) error
	Forward_to(url string) error
}

