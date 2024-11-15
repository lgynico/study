package handler

import (
	"time"

	"github.com/lgynico/gameserver/facade"
	"github.com/lgynico/gameserver/game/module/user"
	"github.com/lgynico/gameserver/game/network/broadcaster"
	"github.com/lgynico/gameserver/pb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func init() {
	codeToHandler[int32(pb.MsgCode_USER_MOVE_TO_CMD.Number())] = moveTo
}

func moveTo(ctx facade.Context, msg proto.Message) error {
	cmd := pb.UserMoveToCmd{}
	msg.ProtoReflect().Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		cmd.ProtoReflect().Set(fd, v)
		return true
	})

	now := uint64(time.Now().UnixMilli())
	if data, ok := user.GetUser(ctx.UserID()); ok {
		data.MoveState = &user.MoveState{
			FromPosX:      cmd.MoveFromPosX,
			FromPosY:      cmd.MoveFromPosY,
			ToPosX:        cmd.MoveToPosX,
			ToPosY:        cmd.MoveToPosY,
			MoveStartTime: now,
		}
	}

	result := pb.UserMoveToResult{
		MoveUserId:    uint32(ctx.UserID()),
		MoveFromPosX:  cmd.MoveFromPosX,
		MoveFromPosY:  cmd.MoveFromPosY,
		MoveToPosX:    cmd.MoveToPosX,
		MoveToPosY:    cmd.MoveToPosY,
		MoveStartTime: now,
	}

	broadcaster.Broadcast(&result)

	return nil
}
