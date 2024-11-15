package tools

import "github.com/lgynico/gameserver/logger"

func Try(code func(), catch func(error)) {
	defer func() {
		if err := recover(); err != nil {
			if catch != nil {
				catch(err.(error))
			} else {
				logger.Error("catch error: %v", err)
			}
		}
	}()

	code()
}
