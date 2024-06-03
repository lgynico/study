package ziface

type Packet interface {
	HeaderLen() int32
	Pack(message Message) ([]byte, error)
	Unpack(data []byte) (Message, error)
}
