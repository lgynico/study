package ziface

type Server interface {
	Start()
	Stop()
	Serve()

	AddRouter(int32, Router)
	ConnManager() ConnManager

	SetOnConnectStart(func(Connection))
	SetOnConnectStop(func(Connection))

	OnConnectStart(Connection)
	OnConnectStop(Connection)
}
