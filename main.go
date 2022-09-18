package main

import (
	"fmt"
	"net"

	register "github.com/Mukunth-arya/grcp/proto"

	servers "github.com/Mukunth-arya/grcp/handlres"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()

	gs := grpc.NewServer()

	servers := servers.LoggerValue(log)

	reflection.Register(gs)

	register.RegisterCgpa_CalculatorServer(gs, servers)

	l, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println("Something went wrong", err)
	}

	gs.Serve(l)

}
