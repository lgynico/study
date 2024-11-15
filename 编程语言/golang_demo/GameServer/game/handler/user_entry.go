package handler

import (
	"github.com/lgynico/gameserver/facade"
	"github.com/lgynico/gameserver/game/module/user"
	"github.com/lgynico/gameserver/game/network/broadcaster"
	"github.com/lgynico/gameserver/pb"
	"google.golang.org/protobuf/proto"
)

func init() {
	codeToHandler[int32(pb.MsgCode_USER_ENTRY_CMD.Number())] = userEntry
}

func userEntry(ctx facade.Context, _ proto.Message) error {

	data, ok := user.GetUser(ctx.UserID())
	if !ok {
		return nil
	}

	userEntryResult := pb.UserEntryResult{
		UserId:     uint32(data.UserID),
		UserName:   data.Username,
		HeroAvatar: data.HeroAvatar,
	}

	broadcaster.Broadcast(&userEntryResult)

	return nil
}
