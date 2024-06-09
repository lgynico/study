package iface

type Connection interface {
	GetConnID() int32
	Write([]byte) error
}
