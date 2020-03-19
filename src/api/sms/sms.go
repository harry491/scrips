package sms

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"scrips/src/db"
	"scrips/src/utils"
)

type EmailForm struct {
	Account    string `form:"account"`
	Code     string `form:"code"`
	Password string `form:"password"`
}

var letters = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

/**
获取随机数
 */
func randEmailVerifyCode(length int) string  {
	b := make([]rune , length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

/**
发送验证码
*/

func SendEmail(c *gin.Context) {

	var email EmailForm

	err := c.ShouldBind(&email)

	if err != nil {
		print(err)
	}

	if utils.CheckEmail(email.Account) == false {
		c.JSON(200, gin.H{
			"msg":   "邮箱格式错误",
			"code":  "10000",
			"email": email,
		})
		return
	}

	code := randEmailVerifyCode(6)

	/// 发送邮件
	go utils.SendEmailTo(email.Account, code)

	db.SaveEmailCode(email.Account, code)

	c.JSON(200, gin.H{
		"msg":  "发送成功",
		"code": "200",
	})
}
