package model

import (
	"encoding/json"
	"scrips/src/utils"
	"time"
)

var BaseSimble = "scrips"

type Token struct {
	Time   int64
	UserId int
	Symble string
}

/**
解析token
*/
func ParseToken(token string) *Token {
	tokenStr, _ := utils.Base64Decode(token)
	var t *Token
	err := json.Unmarshal([]byte(tokenStr), &t)
	if err != nil {
		print(err)
	}
	return t
}

/**
生成token
*/
func CreateToken(userId int) string {
	var token = Token{Time: time.Now().Unix(), UserId: userId, Symble: BaseSimble}
	result, _ := json.Marshal(token)
	return utils.Base64Encode(string(result))
}
