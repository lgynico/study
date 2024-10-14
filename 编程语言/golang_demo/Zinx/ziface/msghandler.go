package ziface

type MsgHandler interface {
	HandleMsg(request Request)
	AddRouter(id int32, router Router)
	StartWorkerPool()
	SendMsgToTaskQueue(request Request)
}
