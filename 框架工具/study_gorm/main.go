package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := getDB()
	autoMigrate(db)

	Create(db)
	CreateBatch(db)

}

func autoMigrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
}

func getDB() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}

	return db
}
