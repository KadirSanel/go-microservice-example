package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "host=127.0.0.1 user=postgres password=secret dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDb() *gorm.DB {
	return db
}
