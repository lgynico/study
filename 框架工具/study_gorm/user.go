package main

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string
	Password     string
	Age          int32
	Email        *string
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivedAt    sql.NullTime
}
