package main

import (
	"log"

	"github.com/Larouimohammed/go_proxies/proxies/direct_proxy"
	"github.com/Larouimohammed/go_proxies/proxies/forwarding_proxy"
	"github.com/Larouimohammed/go_proxies/server"
)

const (
	Forward_Proxy_Port = "localhost:8080"
	Forward_Proxy_URL  = "localhost:8081"
	Direct_Proxy_Port  = "localhost:8081"
	Server_Address     = "localhost:8082"
)

func main() {
	forward_proxy := forwarding_proxy.NewForwardProxy(Forward_Proxy_Port, Forward_Proxy_URL)
	direct_proxy := direct_proxy.NewDirectProxy(Direct_Proxy_Port, Server_Address)
	server := server.NewServer(Server_Address)
	go forward_proxy.Listent_And_Accept(Forward_Proxy_Port)
	go direct_proxy.Listent_And_Accept(Direct_Proxy_Port)
	if err := server.Serve(); err != nil {
		log.Fatal(err)
	}

}

//TODO:
//implment forwardind using net.io packages 
//implement https between curl client and gin server
// goroutine and thread optimization 
// env load into config struct 
// redis data base
// interfaces implementation
// write server with http package
