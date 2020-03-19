package utils

import (
	"encoding/base64"
	"regexp"
)

func CheckEmail(email string) bool {
	//匹配电子邮箱
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

/**
  base64加密
*/
func Base64Encode(src string) string {
	return base64.StdEncoding.EncodeToString([]byte(src))
}

/**
  base64解密
*/
func Base64Decode(src string) (string, error) {
	bytes, err := base64.StdEncoding.DecodeString(src)
	return string(bytes), err
}


