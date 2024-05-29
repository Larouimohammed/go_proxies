package forwarding_proxy

import (
	"fmt"
	"net"
)

func Forward() {
	fmt.Println("Forwarding Proxy Starting....")

	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return

	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {

	defer conn.Close()
}
