package client

import (
	"fmt"
	"github.com/Grivn/go-practice/api"
	"log"
	"net/rpc/jsonrpc"
)

type clientImpl struct {}
func (c *clientImpl) Start() {
	client, err := jsonrpc.Dial("tcp", "10.1.42.91:9999")
	if err != nil {
		log.Fatal("dial error:", err)
	}

	args := "hello!"
	var reply string
	err = client.Call("Arith.Process", args, &reply)
	if err != nil {
		log.Fatal("Multiply error:", err)
	}
	fmt.Printf("%s\n", reply)
}

func (c *clientImpl) Stop() {}

func New() api.Basic {
	return &clientImpl{}
}
