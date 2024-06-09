package iface

type ConnManager interface {
	Add(conn Connection)
	Remove(conn Connection)
	Clear()
}
