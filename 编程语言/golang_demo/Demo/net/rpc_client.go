package main

import (
	"fmt"
	"log"
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

type RpcClient struct {
	client *rpc.Client
}

func (c *RpcClient) Dial(addr string) error {
	client, err := rpc.DialHTTP("tcp", addr)
	if err != nil {
		return err
	}

	c.client = client
	return nil
}

func (c *RpcClient) Call(route string, arg, reply any) error {
	return c.client.Call(route, arg, reply)
}

func (c *RpcClient) Go(route string, arg, reply any) *rpc.Call {
	return c.client.Go(route, arg, reply, nil)
}

func main() {
	c := RpcClient{}
	if err := c.Dial(":1234"); err != nil {
		log.Fatal(err)
	}

	var result int
	if err := c.Call("MathService.Add", &Args{X: 1, Y: 2}, &result); err != nil {
		log.Fatal(err)
	}

	fmt.Println("result is", result)

	var res Result
	call := c.Go("MathService.Divide", &Args{X: 1, Y: 2}, &res)
	done := <-call.Done
	if done.Error != nil {
		log.Fatal(done.Error)
	}

	fmt.Println("result is", res.Quotient, res.Remainder)
}
