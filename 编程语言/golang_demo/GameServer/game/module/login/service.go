package login

import (
	"github.com/lgynico/gameserver/game/module/user"
	"github.com/lgynico/gameserver/logger"
	"github.com/lgynico/gameserver/thread"
)

func Login(username string, password string) *thread.AsyncResult[user.UserData] {
	if username == "" || password == "" {
		return nil
	}

	result := &thread.AsyncResult[user.UserData]{}
	thread.Async.Submit(thread.Hash(username), func() {
		data, err := user.GetUserByUsername(username)
		if err != nil {
			logger.Error("login error: %+v", err)
			return
		}

		if data != nil {
			if data.Password != password {
				logger.Error("user auth failed: %s", username)
				return
			}
			if data.CurrHp <= 0 {
				data.CurrHp = 1000
			}
		} else {
			data = &user.UserData{
				Username:   username,
				Password:   password,
				HeroAvatar: "Hero_Hammer",
				CurrHp:     1000,
			}
		}

		if err := user.SaveOrUpdate(data); err != nil {
			logger.Error("login error: %+v", err)
			return
		}

		result.Complete(data)

	}, nil)

	return result
}
