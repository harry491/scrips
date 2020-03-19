package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"scrips/src/api/sms"
	"scrips/src/api/v1"
	"scrips/src/db"
	"scrips/src/model"
	"time"
)

func BeginService() {
	engine := gin.Default()

	/// 静态文件
	engine.Static("/src/static", "/Users/zhangnan/go/src/scrips/src/static")
	engine.NoRoute(func(c *gin.Context) {
		//返回404状态码
		c.JSON(404, gin.H{
			"time": time.Now().String(),
			"msg":  "error",
		})
	})

	db.InitTables()

	InitRouter(engine)

	err := engine.Run(":80")
	if err != nil {
		log.Fatal("listen port 80 error")
	}
}

func InitRouter(engine *gin.Engine) {

	/// base

	/// 验证码
	engine.Any("/sendEmailCode", sms.SendEmail)

	/// 通用请求 api
	groupV1 := engine.Group("/v1")
	{
		groupV1.Any("/register", v1.Register)
		groupV1.Any("/findPassword" , v1.EditPassword)
		groupV1.Any("/login", v1.Login)

	}

	/// 登录请求
	groupV2 := engine.Group("/v1")
	{
		groupV2.Any("/scrips", v1.AllScrips)
		groupV2.Any("/writeScrip", v1.PublishScrips)
		groupV2.Any("/addPoint", v1.AddPoint)
		groupV2.Any("/addCollect", v1.AddCollect)
		groupV2.Any("/myCollect", v1.AllCollectScrips)
		groupV2.Any("/addComment", v1.AddComment)
		groupV2.Any("/pointComment", v1.PointComment)
		groupV2.Any("/allComments", v1.AllComments)
		groupV2.Any("/getUserInfo", v1.GetUserInfo)
		groupV2.Any("/editUserInfo", v1.EditUserInfo)


	}
}

func verify(c *gin.Context) {

	var method = c.Request.Method
	var token string = ""
	if method == "GET" {
		token = c.Query("token")
	} else if method == "POST" {
		token = c.PostForm("token")
	}

	tokenModel := model.ParseToken(token)

	now := time.Now().Unix()

	if token == "" ||
		tokenModel == nil ||
		tokenModel.UserId == 0 ||
		tokenModel.Time > now ||
		tokenModel.Time < (now-20*60) {
		c.JSON(500, gin.H{
			"msg": "illegal request",
		})
		c.Abort()
	}
}
