package pb

import (
	"encoding/binary"
	"errors"
)

type InnerMsg struct {
	SessionID int32
	UserID    int64
	Data      []byte
}

func (p *InnerMsg) Unmarshal(data []byte) error {
	if len(data) < 12 {
		return errors.New("invalid data")
	}

	p.SessionID = int32(binary.BigEndian.Uint32(data[:4]))
	p.UserID = int64(binary.BigEndian.Uint64(data[4:12]))
	p.Data = data[12:]

	return nil
}

func (p *InnerMsg) Marshal() []byte {
	data := make([]byte, 12+len(p.Data))

	binary.BigEndian.PutUint32(data[0:4], uint32(p.SessionID))
	binary.BigEndian.PutUint64(data[4:12], uint64(p.UserID))
	copy(data[12:], p.Data)

	return data
}
