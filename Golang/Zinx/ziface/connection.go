package ziface

import "net"

type Connection interface {
	Start()
	Stop()

	GetTCPConnection() *net.TCPConn
	GetConnID() int32
	RemoteAddr() net.Addr

	SendMessage(id int32, data []byte) error
	SendBufMessage(id int32, data []byte) error

	SetProperty(key string, value any)
	GetProperty(key string) (any, bool)
	RemoveProperty(key string)
}
