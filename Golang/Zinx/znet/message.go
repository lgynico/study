package znet

import "zinx/ziface"

type Message struct {
	id      int32
	dataLen int32
	data    []byte
}

func NewMessage(id int32, data []byte) ziface.Message {
	return &Message{
		id:      id,
		dataLen: int32(len(data)),
		data:    data,
	}
}

func (p *Message) ID() int32 { return p.id }

func (p *Message) DataLen() int32 { return p.dataLen }

func (p *Message) Data() []byte { return p.data }

func (p *Message) SetID(id int32) { p.id = id }

func (p *Message) SetDataLen(len int32) { p.dataLen = len }

func (p *Message) SetData(data []byte) { p.data = data }
