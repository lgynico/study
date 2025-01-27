package net

import (
	"net"
	"net/rpc"
)

type RpcClient struct {
	client *rpc.Client
}

func (c *RpcClient) DialHTTP(addr string) error {
	client, err := rpc.DialHTTP("tcp", addr)
	if err != nil {
		return err
	}

	c.client = client
	return nil
}

func (c *RpcClient) DialTCP(addr string) error {
	client, err := rpc.Dial("tcp", addr)
	if err != nil {
		return err
	}

	c.client = client
	return nil
}

func (c *RpcClient) DialTCPWithCodec(addr string) error {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}

	c.client = rpc.NewClientWithCodec(NewJsonClientCodec(conn))
	return nil
}

func (c *RpcClient) Call(route string, arg, reply any) error {
	return c.client.Call(route, arg, reply)
}

func (c *RpcClient) Go(route string, arg, reply any) *rpc.Call {
	return c.client.Go(route, arg, reply, nil)
}
