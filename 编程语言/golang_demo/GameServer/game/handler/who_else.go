package handler

import (
	"github.com/lgynico/gameserver/facade"
	"github.com/lgynico/gameserver/game/module/user"
	"github.com/lgynico/gameserver/game/network/broadcaster"
	"github.com/lgynico/gameserver/pb"
	"google.golang.org/protobuf/proto"
)

func init() {
	codeToHandler[int32(pb.MsgCode_WHO_ELSE_IS_HERE_CMD.Number())] = whoElse
}

func whoElse(ctx facade.Context, _ proto.Message) error {

	result := pb.WhoElseIsHereResult{}
	for _, userData := range user.GetUsers() {
		result.UserInfo = append(result.UserInfo, &pb.WhoElseIsHereResult_UserInfo{
			UserId:     uint32(userData.UserID),
			UserName:   userData.Username,
			HeroAvatar: userData.HeroAvatar,
			MoveState:  userData.GetMoveStatePb(),
		})
	}

	broadcaster.Broadcast(&result)

	return nil
}
