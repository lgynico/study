package user

import (
	"fmt"
	"time"

	"github.com/lgynico/gameserver/pb"
	"github.com/lgynico/gameserver/thread"
)

type (
	UserData struct {
		UserID        int64     `db:"user_id"`
		Username      string    `db:"user_name"`
		Password      string    `db:"password"`
		HeroAvatar    string    `db:"hero_avatar"`
		CurrHp        int32     `db:"curr_hp"`
		CreateTime    time.Time `db:"create_time"`
		LastLoginTime time.Time `db:"last_login_time"`
		MoveState     *MoveState
	}

	MoveState struct {
		FromPosX      float32
		FromPosY      float32
		ToPosX        float32
		ToPosY        float32
		MoveStartTime uint64
	}
)

func (p *UserData) GetMoveStatePb() *pb.WhoElseIsHereResult_UserInfo_MoveState {
	if p.MoveState == nil {
		return nil
	}

	return &pb.WhoElseIsHereResult_UserInfo_MoveState{
		FromPosX:  p.MoveState.FromPosX,
		FromPosY:  p.MoveState.FromPosY,
		ToPosX:    p.MoveState.ToPosX,
		ToPosY:    p.MoveState.ToPosY,
		StartTime: p.MoveState.MoveStartTime,
	}
}

func (p *UserData) Save() {
	thread.Async.Submit(p.UserID, func() { _ = SaveOrUpdate(p) }, nil)
}

func (p *UserData) Key() string {
	return fmt.Sprintf("user_%d", p.UserID)
}
