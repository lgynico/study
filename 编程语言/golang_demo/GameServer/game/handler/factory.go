package handler

import (
	"github.com/lgynico/gameserver/facade"
	"google.golang.org/protobuf/proto"
)

type CmdHandleFunc func(ctx facade.Context, cmd proto.Message) error

var codeToHandler = map[int32]CmdHandleFunc{}

func GetHandler(code int32) (CmdHandleFunc, bool) {
	handler, ok := codeToHandler[code]
	return handler, ok
}
