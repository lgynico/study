package handler

import (
	"fmt"

	"github.com/lgynico/gameserver/facade"
	"github.com/lgynico/gameserver/game/module/login"
	"github.com/lgynico/gameserver/game/module/user"
	"github.com/lgynico/gameserver/logger"
	"github.com/lgynico/gameserver/pb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func init() {
	codeToHandler[int32(pb.MsgCode_USER_LOGIN_CMD.Number())] = userLogin
}

func userLogin(ctx facade.Context, cmd proto.Message) error {
	userLoginCmd := pb.UserLoginCmd{}
	cmd.ProtoReflect().Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		userLoginCmd.ProtoReflect().Set(fd, v)
		return true
	})

	logger.Info("receive: username=%s, password=%s", userLoginCmd.UserName, userLoginCmd.Password)

	result := login.Login(userLoginCmd.UserName, userLoginCmd.Password)
	if result == nil {
		logger.Error("user not exists: %v", userLoginCmd.UserName)
		return fmt.Errorf("user not exists: %v", userLoginCmd.UserName)
	}

	result.OnComplete(func(data *user.UserData) {
		ctx.SetUserID(data.UserID)
		user.AddToGroup(data)

		result := pb.UserLoginResult{
			UserId:     uint32(data.UserID),
			UserName:   data.Username,
			HeroAvatar: data.HeroAvatar,
		}

		ctx.SendMessage(&result)
	})

	return nil
}
