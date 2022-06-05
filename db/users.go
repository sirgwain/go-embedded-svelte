package db

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique" header:"Username"`
	Password string `json:"password"`
}

func GetUsers() *[]User {

	db = GetDB()
	users := []User{}
	db.Find(&users)

	return &users
}

func CreateUser(user User) {
	db = GetDB()
	db.Create(&user)
}

func FindUserById(id uint) *User {
	db = GetDB()
	user := User{}
	db.First(&user, id)

	return &user
}

func FindUserByUsername(username string) *User {
	db = GetDB()
	user := User{}
	db.Where("username = ?", username).First(&user)

	return &user
}
