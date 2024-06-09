package iface

type Connector interface {
	Start()
	Stop()

	ConnManager() ConnManager
}
