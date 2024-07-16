package main

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

func Create(db *gorm.DB) {
	user := User{
		Username: "nico",
		Password: "123",
		Age:      30,
	}

	result := db.Create(&user)
	fmt.Printf("user.ID = %v, Error = %v, RowsAffected = %v\n", user.ID, result.Error, result.RowsAffected)
}

func CreateBatch(db *gorm.DB) {
	var (
		now   = time.Now()
		users = []*User{
			{Username: "nico123", Birthday: &now},
			{Username: "nico321", Password: "312"},
		}
		result = db.Create(users)
	)

	fmt.Printf("Error = %v, RowsAffected = %v\n", result.Error, result.RowsAffected)
	for _, user := range users {
		fmt.Printf("username: %s, id: %d\n", user.Username, user.ID)
	}
}
