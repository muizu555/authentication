package models

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
}

func (u User) Save() (User, error) {
	//fmt.Println(u)
	fmt.Println(DB, "ほげ")
	err := DB.Create(&u) //ここでエラーが出てる
	//fmt.Println(u)
	if err != nil {
		return User{}, err.Error
	}
	return u, nil
}

// 自動的に呼び出される
func (u *User) BeforeSave() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)

	u.Username = strings.ToLower(u.Username)

	return nil
}

func (u User) PrepareOutput() User {
	u.Password = ""
	return u
}
