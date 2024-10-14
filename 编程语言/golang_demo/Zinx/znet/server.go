package znet

import (
	"fmt"
	"log"
	"net"
	"zinx/utils"
	"zinx/ziface"
)

type Server struct {
	Name        string
	IPVersion   string
	IP          string
	Port        int32
	MsgHandler  ziface.MsgHandler
	connManager ziface.ConnManager

	connStartHook func(ziface.Connection)
	connStopHook  func(ziface.Connection)
}

func NewServer() ziface.Server {
	utils.GlobalObject.Reload()

	s := &Server{
		Name:        utils.GlobalObject.Name,
		IPVersion:   "tcp4",
		IP:          utils.GlobalObject.Host,
		Port:        utils.GlobalObject.Port,
		MsgHandler:  NewMsgHandler(),
		connManager: NewConnManager(),
	}
	// utils.GlobalObject.Server = s

	return s
}

func (p *Server) Start() {
	log.Printf("[START] Server listen at %s:%d is starting\n", p.IP, p.Port)
	log.Printf("[Zinx] Version: %s, MaxConn: %d, MaxPacketSize: %d\n",
		utils.GlobalObject.Version, utils.GlobalObject.MaxConn, utils.GlobalObject.MaxPacketSize)

	p.MsgHandler.StartWorkerPool()

	addr, err := net.ResolveTCPAddr(p.IPVersion, fmt.Sprintf("%s:%d", p.IP, p.Port))
	if err != nil {
		log.Println("resolve tcp addr fail:", err)
		return
	}

	listener, err := net.ListenTCP(p.IPVersion, addr)
	if err != nil {
		log.Println("listen fail:", err)
		return
	}

	log.Printf("start Zinx server %s success, now listenning...\n", p.Name)

	var connID = int32(1)

	go func() {
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				log.Println("accept fail:", err)
				continue
			}

			// TODO: max connect control
			// TODO: bind connect handler

			newConn := NewConnection(p, conn, connID, p.MsgHandler)
			go newConn.Start()

			connID++
		}
	}()

}

func (p *Server) Stop() {
	log.Println("[STOP] Zinx server", p.Name)

	// TODO: clean connect and resource
	p.connManager.CleanConn()
}

func (p *Server) Serve() {
	p.Start()

	// TODO: handle after server start

	select {}
}

func (p *Server) AddRouter(id int32, router ziface.Router) {
	p.MsgHandler.AddRouter(id, router)
}

func (p *Server) ConnManager() ziface.ConnManager {
	return p.connManager
}

func (p *Server) SetOnConnectStart(hook func(ziface.Connection)) {
	p.connStartHook = hook
}

func (p *Server) SetOnConnectStop(hook func(ziface.Connection)) {
	p.connStopHook = hook
}

func (p *Server) OnConnectStart(conn ziface.Connection) {
	if p.connStartHook != nil {
		log.Println("call OnConnectStart")
		p.connStartHook(conn)
	}
}

func (p *Server) OnConnectStop(conn ziface.Connection) {
	if p.connStopHook != nil {
		log.Println("call OnConnectStop")
		p.connStopHook(conn)
	}
}
