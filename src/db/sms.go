package db

import (
	"scrips/src/model"
)

/**
存储验证码
*/
func SaveEmailCode(email string, code string) {
	if DB.HasTable(&model.SmsModel{}) == false {
		DB.CreateTable(&model.SmsModel{})
	}

	DB.Create(&model.SmsModel{Email: email, Code: code})
}

/**
验证验证码
*/

func SearchEmailCode(email string) *model.SmsModel {

	smsModel := &model.SmsModel{Email: email}

	DB.Last(smsModel)

	return smsModel
}

/**
删除验证码
*/
func DeleteEmailCode(email string) {

	smsModel := &model.SmsModel{Email: email}

	DB.Delete(smsModel)

}
