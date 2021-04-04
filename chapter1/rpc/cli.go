package main

import (
	"fmt"

	"github.com/masuldev/micro-go/chapter1/rpc/client"
	"github.com/masuldev/micro-go/chapter1/rpc/server"
)

func main() {
	go server.StartServer()

	c := client.CreateClient()
	defer c.Close()

	reply := client.PerformRequest(c)
	fmt.Println(reply.Message)
}