package znet

import (
	"fmt"
	"log"
	"zinx/utils"
	"zinx/ziface"
)

type MsgHandler struct {
	apis         map[int32]ziface.Router
	workPoolSize int32
	taskQueue    []chan ziface.Request
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		apis:         make(map[int32]ziface.Router),
		workPoolSize: utils.GlobalObject.WorkerPoolSize,
		taskQueue:    make([]chan ziface.Request, utils.GlobalObject.WorkerPoolSize),
	}
}

func (p *MsgHandler) HandleMsg(request ziface.Request) {
	handler, ok := p.apis[request.GetMsgID()]
	if !ok {
		log.Printf("message handler not found: %d\n", request.GetMsgID())
		return
	}

	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}

func (p *MsgHandler) AddRouter(id int32, router ziface.Router) {
	if _, ok := p.apis[id]; ok {
		panic(fmt.Sprintf("exist handler: %d", id))
	}

	p.apis[id] = router
	log.Println("add router success:", id)
}

func (p *MsgHandler) StartWorker(workerID int32, taskQueue chan ziface.Request) {
	log.Printf("worker %d is started\n", workerID)

	for request := range taskQueue {
		p.HandleMsg(request)
	}
}

func (p *MsgHandler) StartWorkerPool() {
	for i := int32(0); i < p.workPoolSize; i++ {
		p.taskQueue[i] = make(chan ziface.Request, utils.GlobalObject.TaskQueueLen)
		go p.StartWorker(i, p.taskQueue[i])
	}
}

func (p *MsgHandler) SendMsgToTaskQueue(request ziface.Request) {
	workerID := request.GetConnection().GetConnID() % p.workPoolSize
	log.Printf("add task[connID=%d, msgID=%d] to worker %d\n", request.GetConnection().GetConnID(), request.GetMsgID(), workerID)
	p.taskQueue[workerID] <- request
}
