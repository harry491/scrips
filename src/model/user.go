package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"not null;unique"` // 设置字段为非空并唯一
	Password string
	Alias    string `gorm:"default:'小纸条'"`
	Sex      int
	Head     string
	Token    string `gorm:"-"`
	Area     string
}
