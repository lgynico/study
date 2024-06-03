package main

import (
	"log"
	"strconv"
	"zinx/ziface"
	"zinx/znet"
)

func main() {
	server := znet.NewServer()

	server.AddRouter(0, &PingRouter{})
	server.AddRouter(1, &HelloRouter{})

	server.SetOnConnectStart(onConnectStart)
	server.SetOnConnectStop(onConnectStop)

	server.Serve()
}

func onConnectStart(conn ziface.Connection) {
	log.Println("call on connect started....")

	conn.SetProperty("Name", "Nico")
	conn.SetProperty("Cookie", 123)

	conn.SendMessage(0, []byte("I am comming......"+strconv.Itoa(int(conn.GetConnID()))))
}

func onConnectStop(conn ziface.Connection) {
	log.Println("call on connect stoped......")

	if value, ok := conn.GetProperty("Name"); ok {
		log.Println("conn properties [Name]:", value)
	}

	if value, ok := conn.GetProperty("Cookie"); ok {
		log.Println("conn properties [Cookie]:", value)
	}
}

type PingRouter struct {
	znet.BaseRouter
}

func (p *PingRouter) Handle(req ziface.Request) {
	log.Println("Call PingRouter Handle")

	log.Printf("recv message (%d): %s\n", req.GetMsgID(), string(req.GetMsgData()))

	err := req.GetConnection().SendMessage(0, []byte("ping...ping...ping..."))
	if err != nil {
		log.Println("callback ping ping ping err:", err)
	}
}

type HelloRouter struct {
	znet.BaseRouter
}

func (p *HelloRouter) Handle(req ziface.Request) {
	log.Println("Call HelloRouter Handle")

	log.Printf("recv message (%d): %s\n", req.GetMsgID(), string(req.GetMsgData()))

	err := req.GetConnection().SendMessage(2, []byte("Hello Znix Server v0.6"))
	if err != nil {
		log.Println("callback hero err:", err)
	}
}
