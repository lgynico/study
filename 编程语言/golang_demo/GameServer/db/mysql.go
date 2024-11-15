package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lgynico/gameserver/logger"
)

var Mysql *sql.DB

func Init() {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/hero_story")
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(128)
	db.SetMaxIdleConns(16)
	db.SetConnMaxIdleTime(2 * time.Minute)

	if err := db.Ping(); err != nil {
		panic(err)
	}

	Mysql = db

	logger.Info("mysql init seccuess !!!")
}
