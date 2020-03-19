package model

import "github.com/jinzhu/gorm"

type SmsModel struct {
	gorm.Model
	Code  string `json:"code"`
	Email string `json:"email"`
}
