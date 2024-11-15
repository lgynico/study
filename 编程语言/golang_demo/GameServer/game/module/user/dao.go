package user

import (
	"errors"
	"time"

	"github.com/lgynico/gameserver/db"
)

const (
	sqlQueryUserByUsername = `SELECT user_id, user_name, password, hero_avatar, curr_hp FROM t_user WHERE user_name = ?`
	sqlSaveOrUpdate        = `INSERT INTO t_user (user_name, password, hero_avatar, curr_hp, create_time) VALUES(?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE curr_hp = ?, last_login_time = ? `
)

func GetUserByUsername(username string) (*UserData, error) {
	if username == "" {
		return nil, errors.New("username is empty")
	}

	stmt, err := db.Mysql.Prepare(sqlQueryUserByUsername)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(username)
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, nil
	}

	data := UserData{}
	if err = rows.Scan(&data.UserID, &data.Username, &data.Password, &data.HeroAvatar, &data.CurrHp); err != nil {
		return nil, err
	}

	return &data, nil
}

func SaveOrUpdate(data *UserData) error {
	if data == nil {
		return nil
	}

	stmt, err := db.Mysql.Prepare(sqlSaveOrUpdate)
	if err != nil {
		return err
	}

	now := time.Now().Unix()
	result, err := stmt.Exec(data.Username, data.Password, data.HeroAvatar, data.CurrHp, now, data.CurrHp, now)
	if err != nil {
		return err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	data.UserID = userID

	return nil
}
