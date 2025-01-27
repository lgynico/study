package rpcx

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/server"
)

func TestServer(t *testing.T) {
	s := server.NewServer()
	s.RegisterName("Arith", &Arith{}, "")
	s.Serve("tcp", "localhost:8972")
}

func TestMultiServer(t *testing.T) {
	for i := 0; i < 3; i++ {
		go func(i int) {
			s := &Server{
				Server: *server.NewServer(),
				name:   fmt.Sprintf("server_%d", i),
			}
			s.RegisterName("Server", s, "")
			s.Serve("tcp", fmt.Sprintf("localhost:897%d", i))
		}(i + 1)
	}

	d, err := client.NewMultipleServersDiscovery([]*client.KVPair{
		{Key: "tcp@localhost:8971", Value: "weight=1"},
		{Key: "tcp@localhost:8972", Value: "weight=1"},
		{Key: "tcp@localhost:8973", Value: "weight=2"},
	})

	if err != nil {
		log.Fatalf("failed to NewMultipleServersDiscovery: %v", err)
	}

	c := client.NewXClient("Server", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer c.Close()

	for i := 0; i < 10; i++ {
		var (
			req   = &ArithArgs{7, 8}
			reply = &ArithReply{}
			err   = c.Call(context.Background(), "Service", req, reply)
		)

		if err != nil {
			log.Fatalf("failed to call: %v", err)
		}
	}
}

func TestClient(t *testing.T) {
	c := client.NewClient(client.DefaultOption)
	if err := c.Connect("tcp", "localhost:8972"); err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer c.Close()

	var (
		args  = &ArithArgs{7, 8}
		reply = &ArithReply{}
	)

	if err := c.Call(context.Background(), "Arith", "Mul", args, reply); err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("Mul: %d * %d = %d", args.A, args.B, reply.C)

	call := c.Go(context.Background(), "Arith", "Add", args, reply, nil)
	replyCall := <-call.Done
	if replyCall.Error != nil {
		log.Fatalf("failed to call: %v", replyCall.Error)
	} else {
		log.Printf("Add: %d + %d = %d", args.A, args.B, reply.C)
	}
}

func TestXClient(t *testing.T) {
	d, err := client.NewPeer2PeerDiscovery("tcp@localhost:8972", "")
	if err != nil {
		log.Fatalf("failed to NewPeer2PeerDiscovery: %v", err)
	}

	c := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer c.Close()

	var (
		args  = &ArithArgs{7, 8}
		reply = &ArithReply{}
	)

	if err = c.Call(context.Background(), "Mul", args, reply); err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("Mul: %d * %d = %d", args.A, args.B, reply.C)

	call, err := c.Go(context.Background(), "Add", args, reply, nil)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	replyCall := <-call.Done
	if replyCall.Error != nil {
		log.Fatalf("failed to call: %v", replyCall.Error)
	} else {
		log.Printf("Add: %d + %d = %d", args.A, args.B, reply.C)
	}
}
