package facade

import "google.golang.org/protobuf/proto"

type Context interface {
	SetUserID(userID int64)
	UserID() int64
	SessionID() int32
	SendMessage(msg proto.Message)
	SendCode(code int32)
	Start()
	Stop()
}
