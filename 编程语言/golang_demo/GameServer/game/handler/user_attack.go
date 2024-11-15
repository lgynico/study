package handler

import (
	"github.com/lgynico/gameserver/facade"
	"github.com/lgynico/gameserver/game/dbsave"
	"github.com/lgynico/gameserver/game/module/user"
	"github.com/lgynico/gameserver/game/network/broadcaster"
	"github.com/lgynico/gameserver/logger"
	"github.com/lgynico/gameserver/pb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func init() {
	codeToHandler[int32(pb.MsgCode_USER_ATTK_CMD.Number())] = userAttack
}

func userAttack(ctx facade.Context, msg proto.Message) error {
	cmd := pb.UserAttkCmd{}
	msg.ProtoReflect().Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		cmd.ProtoReflect().Set(fd, v)
		return true
	})

	targetUser, ok := user.GetUser(int64(cmd.TargetUserId))
	if !ok {
		return nil
	}

	logger.Info("attack: %d -> %d", ctx.UserID(), cmd.TargetUserId)

	attackResult := pb.UserAttkResult{
		AttkUserId:   uint32(ctx.UserID()),
		TargetUserId: cmd.TargetUserId,
	}

	broadcaster.Broadcast(&attackResult)

	subHp := uint32(10)
	subHpResult := pb.UserSubtractHpResult{
		TargetUserId: cmd.TargetUserId,
		SubtractHp:   subHp,
	}

	broadcaster.Broadcast(&subHpResult)

	targetUser.CurrHp -= int32(subHp)

	dbsave.DBSaver.Save(dbsave.NewLazyRecord(targetUser))

	return nil
}
