package rpcx

import (
	"context"
	"log"

	"github.com/smallnest/rpcx/server"
)

type Server struct {
	server.Server
	name string
}

// func (p *Server) Register(name string, service any) {

// }

func (p *Server) Service(ctx context.Context, args *ArithArgs, reply *ArithReply) error {
	log.Printf("Service by %s", p.name)
	reply.C = args.A * args.B
	return nil
}
