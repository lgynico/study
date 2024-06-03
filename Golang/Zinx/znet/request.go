package znet

import "zinx/ziface"

type Request struct {
	conn    ziface.Connection
	message ziface.Message
}

func (p *Request) GetConnection() ziface.Connection {
	return p.conn
}

func (p *Request) GetMsgData() []byte {
	return p.message.Data()
}

func (p *Request) GetMsgID() int32 {
	return p.message.ID()
}
