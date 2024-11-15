package handler

import (
	"github.com/lgynico/gameserver/facade"
	"github.com/lgynico/gameserver/game/dbsave"
	"github.com/lgynico/gameserver/game/module/user"
	"github.com/lgynico/gameserver/game/network/broadcaster"
	"github.com/lgynico/gameserver/pb"
)

func OnUserQuit(ctx facade.Context) {
	defer user.RemoveFromGroup(ctx.UserID())

	result := pb.UserQuitResult{
		QuitUserId: uint32(ctx.UserID()),
	}

	broadcaster.Broadcast(&result)

	if userData, ok := user.GetUser(ctx.UserID()); ok {
		if saver, ok := dbsave.DBSaver.GetSaver(userData.Key()); ok {
			dbsave.DBSaver.Discard(saver.Key())
			saver.Save()
		}
	}

}
