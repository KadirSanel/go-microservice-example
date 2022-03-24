package models

import (
	"github.com/KadirSanel/go-microservice-example/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Name     string `gorm:""json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func init() {
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&User{})
}

func (u *User) CreateUser() *User {
	db.Create(&u)
	return u
}

func GetAllUser() []User {
	var Users []User
	db.Find(&Users)
	return Users
}

func GetUserById(Id int64) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("ID=?", Id).Find(&getUser)
	return &getUser, db
}

func DeleteUser(Id int64) User {
	var user User
	db.Where("ID=?", Id).Delete(user)
	return user
}
