package infrastructure

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GetDB return db variable
func GetDB() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/books?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	return db
}
