package znet

import "zinx/ziface"

type BaseRouter struct {
}

func (p *BaseRouter) PreHandle(req ziface.Request)  {}
func (p *BaseRouter) Handle(req ziface.Request)     {}
func (p *BaseRouter) PostHandle(req ziface.Request) {}
