package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/rpc"
)

type Args struct {
	X int
	Y int
}

type Result struct {
	Quotient  int
	Remainder int
}

type MathService struct{}

func (s *MathService) Add(args *Args, result *int) error {
	*result = args.X + args.Y
	return nil
}

func (s *MathService) Divide(args *Args, result *Result) error {
	if args.Y == 0 {
		return errors.New("division by zero")
	}

	result.Quotient = args.X / args.Y
	result.Remainder = args.X % args.Y
	return nil
}

type RpcServer struct {
}

func (s *RpcServer) Start(addr string) error {
	rpc.HandleHTTP()
	return http.ListenAndServe(addr, nil)
}

func (s *RpcServer) Register(service any) {
	rpc.Register(service)
}

func main() {
	server := &RpcServer{}
	server.Register(&MathService{})

	fmt.Println("rpc server start")
	err := server.Start(":1234")
	if err != nil {
		panic(err)
	}
}
