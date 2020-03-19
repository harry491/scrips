package v1

import (
	"github.com/gin-gonic/gin"
	"scrips/src/api/sms"
	"scrips/src/db"
	"scrips/src/model"
	"scrips/src/utils"
)


/**
注册
*/
func Register(c *gin.Context) {

	var request sms.EmailForm

	c.ShouldBind(&request)

	email := request.Account
	code := request.Code
	password := request.Password

	if utils.CheckEmail(email) == false {
		c.JSON(200, gin.H{
			"msg":  "邮箱格式错误",
			"code": "10000",
		})
		return
	}

	if len(password) < 6 {
		c.JSON(200, gin.H{
			"msg":  "请输入正确密码格式(6位以上)",
			"code": "10001",
		})
		return
	}

	checkUser := db.SearchUser(email)

	if checkUser != nil {
		c.JSON(200, gin.H{
			"code": "10005",
			"msg":  "用户已存在",
		})
		return
	}

	/// 验证码验证
	smsCode := db.SearchEmailCode(email)

	if smsCode.Code != code {
		c.JSON(200, gin.H{
			"msg":  "验证码错误",
			"code": "10002",
		})
		return
	}

	db.DeleteEmailCode(email)

	/// 创建用户
	user := &model.User{Email: email, Password: utils.Base64Encode(password)}

	db.CreateUser(user)

	c.JSON(200, gin.H{
		"code": "200",
		"msg":  "注册成功",
	})
}

/**
找回密码
 */
func EditPassword(c *gin.Context)  {
	var request sms.EmailForm

	c.ShouldBind(&request)

	email := request.Account
	password := request.Password
	if utils.CheckEmail(email) == false {
		c.JSON(200, gin.H{
			"msg":  "邮箱格式错误",
			"code": "10000",
		})
		return
	}

	if len(password) < 6 {
		c.JSON(200, gin.H{
			"msg":  "请输入正确密码格式(6位以上)",
			"code": "10001",
		})
		return
	}

	user := db.SearchUser(email)
	if user == nil {
		c.JSON(200, gin.H{
			"msg":  "用户不存在",
			"code": "10006",
		})
		return
	}

	base64Password := utils.Base64Encode(password)

	db.EditPassword(email , base64Password)

	c.JSON(200, gin.H{
		"msg":  "修改成功",
		"code": "200",
	})
}

/**
登录
*/
func Login(c *gin.Context) {
	var request sms.EmailForm

	c.ShouldBind(&request)

	email := request.Account
	password := request.Password
	if utils.CheckEmail(email) == false {
		c.JSON(200, gin.H{
			"msg":  "邮箱格式错误",
			"code": "10000",
		})
		return
	}

	if len(password) < 6 {
		c.JSON(200, gin.H{
			"msg":  "请输入正确密码格式(6位以上)",
			"code": "10001",
		})
		return
	}

	user := db.SearchUser(email)
	if user == nil {
		c.JSON(200, gin.H{
			"msg":  "用户不存在",
			"code": "10006",
		})
		return
	}

	base64Password := utils.Base64Encode(password)

	if user.Password != base64Password {
		c.JSON(200, gin.H{
			"msg":  "用户名或密码不正确",
			"code": "10007",
		})
		return
	}

	/// 生成用户令牌
	token := model.CreateToken(int(user.ID))
	user.Token = token

	c.JSON(200, gin.H{
		"msg":  "登录成功",
		"code": "200",
		"user": user,
	})
}
