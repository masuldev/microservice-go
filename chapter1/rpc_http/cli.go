package main

import (
	"fmt"

	"github.com/masuldev/micro-go/chapter1/rpc_http/client"
	"github.com/masuldev/micro-go/chapter1/rpc_http/server"
)

func main() {
	go server.StartServer()

	c := client.CreateClient()
	defer c.Close()

	reply := client.PerformRequest(c)
	fmt.Println(reply.Message)
}
