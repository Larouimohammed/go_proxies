package direct_proxy

import (
	"fmt"
	"net"
)

type DirectProxy struct {
	port string
	url  string
}

func NewDirectProxy(port string, url string) *DirectProxy {

	return &DirectProxy{
		port: port,

		url:  url}

}
func (dp *DirectProxy) Listent_And_Accept(port string) error {

	fmt.Printf("Direct Proxy Start Listening On Port : %s\n",port)

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
		go dp.Handle_conn(conn)
	}
}
func (dp *DirectProxy) Handle_conn(conn net.Conn) {

	defer conn.Close()
	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("Data Received: %s\n", buffer[:n])
	}
}

func (dp *DirectProxy) Forward_to(url string) {

}
