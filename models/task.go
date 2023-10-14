package models

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	Desc   string `gorm:"size:255;not null;" json:"desc"`
	UserId string `gorm:"size:255;not null;" json:"userId"`
}
