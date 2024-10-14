package ziface

type Request interface {
	GetConnection() Connection
	GetMsgData() []byte
	GetMsgID() int32
}
