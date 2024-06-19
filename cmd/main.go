package main

import (
	"log"

	"github.com/Larouimohammed/go_proxies/proxies/direct_proxy"
	"github.com/Larouimohammed/go_proxies/proxies/forwarding_proxy"
	"github.com/Larouimohammed/go_proxies/server"
)

const (
	Forward_Proxy_Addr = "localhost:8080"
	Direct_Proxy_Addr  = "localhost:8081"
	Server_Addr     = "localhost:8082"
)

func main() {
	forward_proxy := forwarding_proxy.NewForwardProxy(Forward_Proxy_Addr, Direct_Proxy_Addr)
	direct_proxy := direct_proxy.NewDirectProxy(Direct_Proxy_Addr, Server_Addr)
	server := server.NewServer(Server_Addr)
	go forward_proxy.Listent_And_Accept(Forward_Proxy_Addr)
	go direct_proxy.Listent_And_Accept(Direct_Proxy_Addr)
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
