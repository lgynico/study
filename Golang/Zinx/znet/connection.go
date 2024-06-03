package znet

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
	"net"
	"sync"
	"zinx/utils"
	"zinx/ziface"
)

type Connection struct {
	conn     *net.TCPConn
	server   ziface.Server
	connID   int32
	isClosed bool

	msgHandler ziface.MsgHandler
	exitChan   chan bool

	msgChan    chan []byte
	msgBufChan chan []byte

	properties sync.Map
}

func NewConnection(server ziface.Server, conn *net.TCPConn, connID int32, msgHandler ziface.MsgHandler) *Connection {
	c := &Connection{
		conn:       conn,
		connID:     connID,
		msgHandler: msgHandler,
		isClosed:   false,
		exitChan:   make(chan bool),
		msgChan:    make(chan []byte),
		msgBufChan: make(chan []byte, utils.GlobalObject.MaxMsgBufChan),
		properties: sync.Map{},
		server:     server,
	}

	c.server.ConnManager().Add(c)

	return c
}

func (p *Connection) Start() {
	go p.StartReader()
	go p.StartWriter()

	p.server.OnConnectStart(p)

	<-p.exitChan
	// for {
	// 	select {
	// 	case <-p.exitChan:
	// 		return
	// 	}
	// }
}

func (p *Connection) StartReader() {
	log.Println("reader goroutine is running...")
	defer log.Println(p.RemoteAddr().String(), "conn reader exit!")
	defer p.Stop()

	for {
		var (
			packet = &Packet{}
			buffer = make([]byte, packet.HeaderLen())
		)

		if _, err := io.ReadFull(p.GetTCPConnection(), buffer); err != nil {
			log.Println("read message fail:", err)
			p.exitChan <- true
			return
		}

		message, err := packet.Unpack(buffer)
		if err != nil {
			log.Println("unpack message fail:", err)
			p.exitChan <- true
			return
		}

		if utils.GlobalObject.MaxPacketSize > 0 && message.DataLen() > utils.GlobalObject.MaxPacketSize {
			log.Println("recv too large data")
			p.exitChan <- true
			return
		}

		if message.DataLen() > 0 {
			buffer = make([]byte, message.DataLen())
			if _, err := io.ReadFull(p.GetTCPConnection(), buffer); err != nil {
				log.Println("read message fail:", err)
				p.exitChan <- true
				return
			}

			reader := bytes.NewReader(buffer)
			data := make([]byte, message.DataLen())
			if err := binary.Read(reader, binary.LittleEndian, data); err != nil {
				log.Println("read message fail:", err)
				p.exitChan <- true
				return
			}

			message.SetData(data)
		}

		req := Request{
			conn:    p,
			message: message,
		}

		if utils.GlobalObject.WorkerPoolSize > 0 {
			p.msgHandler.SendMsgToTaskQueue(&req)
		} else {
			go p.msgHandler.HandleMsg(&req)
		}
	}
}

func (p *Connection) StartWriter() {
	log.Println("writer goroutine is running...")
	defer log.Println(p.RemoteAddr().String(), "conn writer exit!")

	for {
		select {
		case data := <-p.msgChan:
			if _, err := p.conn.Write(data); err != nil {
				log.Println("write message fail:", err)
				return
			}
		case data, ok := <-p.msgBufChan:
			if ok {
				if _, err := p.conn.Write(data); err != nil {
					log.Println("write message fail:", err)
					return
				}
			} else {
				log.Println("channel closed")
			}
		case <-p.exitChan:
			return
		}
	}
}

func (p *Connection) Stop() {
	if p.isClosed {
		return
	}

	p.isClosed = true

	// TODO: connect close callback
	p.server.ConnManager().Remove(p)
	p.server.OnConnectStop(p)

	p.conn.Close()
	p.exitChan <- true
	close(p.exitChan)
}

func (p *Connection) SendMessage(id int32, data []byte) error {
	message := NewMessage(id, data)
	packet := &Packet{}
	pack, err := packet.Pack(message)
	if err != nil {
		return err
	}

	p.msgChan <- pack
	return nil
}

func (p *Connection) SendBufMessage(id int32, data []byte) error {
	message := NewMessage(id, data)
	packet := &Packet{}
	pack, err := packet.Pack(message)
	if err != nil {
		return err
	}

	p.msgBufChan <- pack
	return nil
}

func (p *Connection) GetTCPConnection() *net.TCPConn {
	return p.conn
}

func (p *Connection) GetConnID() int32 {
	return p.connID
}

func (p *Connection) RemoteAddr() net.Addr {
	return p.conn.RemoteAddr()
}

func (p *Connection) SetProperty(key string, value any) {
	p.properties.Store(key, value)
}

func (p *Connection) GetProperty(key string) (any, bool) {
	return p.properties.Load(key)
}

func (p *Connection) RemoveProperty(key string) {
	p.properties.Delete(key)
}
