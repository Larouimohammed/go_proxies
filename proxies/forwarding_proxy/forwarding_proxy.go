package forwarding_proxy

import (
	"fmt"
	"net"
)

type ForwardProxy struct {
	port string
	url  string
}

func NewForwardProxy(port string, url string) *ForwardProxy {

	return &ForwardProxy{
		port: port,

		url: url}

}
func (fp *ForwardProxy) Listent_And_Accept(port string) error {
	fmt.Printf("Forward Proxy Start Listening On Port : %s\n", port)
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Error:", err)
		return err

	}

	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		go fp.Handle_conn(conn)
	}

}

func (fp *ForwardProxy) Handle_conn(conn net.Conn) error {

	defer conn.Close()
	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}

		fmt.Printf("Data Received: %s\n", buffer[:n])
	}
	
}

func (fp *ForwardProxy) Forward_to(url string) error {
	return nil

}
