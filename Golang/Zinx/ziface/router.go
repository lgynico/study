package ziface

type Router interface {
	PreHandle(Request)
	Handle(Request)
	PostHandle(Request)
}
