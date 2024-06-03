package znet

import (
	"log"
	"sync"
	"sync/atomic"
	"zinx/ziface"
)

type ConnManager struct {
	connections sync.Map
	size        atomic.Int32
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: sync.Map{},
	}
}

func (p *ConnManager) Add(conn ziface.Connection) {
	p.connections.Store(conn.GetConnID(), conn)
	p.size.Add(1)

	log.Printf("add connection %d to manager (size=%d)\n", conn.GetConnID(), p.Len())
}

func (p *ConnManager) Remove(conn ziface.Connection) {
	p.connections.Delete(conn.GetConnID())
	p.size.Add(-1)

	log.Printf("remove connection %d from manager (size=%d)\n", conn.GetConnID(), p.Len())
}

func (p *ConnManager) CleanConn() {
	p.connections.Range(func(key, value any) bool {
		conn := value.(ziface.Connection)
		conn.Stop()
		p.connections.Delete(key)

		return true
	})
	p.size.Store(0)

	log.Printf("clean connections from manager (size=%d)\n", p.Len())
}

func (p *ConnManager) Get(connID int32) (ziface.Connection, bool) {
	if value, ok := p.connections.Load(connID); ok {
		return value.(ziface.Connection), true
	}

	return nil, false
}

func (p *ConnManager) Len() int {
	return int(p.size.Load())
}
