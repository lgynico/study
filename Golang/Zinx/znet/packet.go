package znet

import (
	"bytes"
	"encoding/binary"
	"zinx/ziface"
)

type Packet struct {
}

func (p *Packet) HeaderLen() int32 {
	// int32(dataLen) + int32(id)
	return 8
}

func (p *Packet) Pack(message ziface.Message) ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	if err := binary.Write(buffer, binary.LittleEndian, message.DataLen()); err != nil {
		return nil, err
	}

	if err := binary.Write(buffer, binary.LittleEndian, message.ID()); err != nil {
		return nil, err
	}

	if err := binary.Write(buffer, binary.LittleEndian, message.Data()); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (p *Packet) Unpack(data []byte) (ziface.Message, error) {
	var (
		reader  = bytes.NewReader(data)
		message = &Message{}
	)

	if err := binary.Read(reader, binary.LittleEndian, &message.dataLen); err != nil {
		return nil, err
	}

	if err := binary.Read(reader, binary.LittleEndian, &message.id); err != nil {
		return nil, err
	}

	// message.data = make([]byte, message.dataLen)
	return message, nil
}
