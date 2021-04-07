package server

import (
	"fmt"
	"github.com/Grivn/go-practice/api"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Arith int

func (t *Arith) Process(args string, reply *string) error {
	fmt.Printf("%s\n", args)

	str := "hi"
	reply = &str

	return nil
}

type serverImpl struct {}

func (s *serverImpl) Start() {
	l, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	arith := new(Arith)

	err = rpc.Register(arith)
	if err != nil {
		return
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}

		// 注意这一行
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

func (s *serverImpl) Stop() {}

func New() api.Basic {
	return &serverImpl{}
}
