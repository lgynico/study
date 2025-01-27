package rpcx

import "context"

type Arith struct {
}

func (p *Arith) Mul(ctx context.Context, args *ArithArgs, reply *ArithReply) error {
	reply.C = args.A * args.B
	return nil
}

func (p *Arith) Add(ctx context.Context, args *ArithArgs, reply *ArithReply) error {
	reply.C = args.A + args.B
	return nil
}
