package models

import (
	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	Desc   string `gorm:"size:255;not null;" json:"desc"`
	UserId int    `gorm:"size:255;not null;" json:"userId"`
}

func (u Task) Save() (Task, error) {
	err := DB.Create(&u).Error
	if err != nil {
		return Task{}, err
	}
	return u, nil
}
