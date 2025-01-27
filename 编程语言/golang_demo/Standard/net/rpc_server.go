package net

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"net/url"
)

type RpcServer struct {
	server *rpc.Server
}

func (s *RpcServer) RegisterToDiscovery(addr string) error {
	resp, err := http.PostForm(addr+"/register", url.Values{
		"node_type": []string{"rpc"},
		"addr":      []string{"127.0.0.1:1234"},
	})
	if err != nil {
		return err
	}

	fmt.Println(resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("register failed: %d", resp.StatusCode)
	}

	return nil
}

func (s *RpcServer) ServeHTTP(addr string) error {
	s.server = rpc.NewServer()
	s.server.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)
	// rpc.HandleHTTP()
	return http.ListenAndServe(addr, nil)
}

func (s *RpcServer) ServeTCP(addr string) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}

		// go rpc.ServeConn(conn)
		// go rpc.ServeCodec(NewJsonRpcCodec(conn))
		go s.server.ServeCodec(NewJsonRpcCodec(conn))
	}

	// rpc.Accept(listener)
	// return nil
}

func (s *RpcServer) Register(service any) error {
	return s.server.Register(service)
}

func (s *RpcServer) RegisterName(name string, service any) error {
	return s.server.RegisterName(name, service)
}
