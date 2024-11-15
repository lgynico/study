package broadcaster

import (
	"sync"

	"github.com/lgynico/gameserver/facade"
	"google.golang.org/protobuf/proto"
)

var group sync.Map

func Add(ctx facade.Context) {
	if ctx == nil {
		return
	}

	group.Store(ctx.SessionID(), ctx)
}

func Remove(sessionID int32) {
	group.Delete(sessionID)
}

func Broadcast(msg proto.Message) {
	group.Range(func(key, value any) bool {
		if ctx, ok := value.(facade.Context); ok {
			ctx.SendMessage(msg)
		}
		return true
	})
}
