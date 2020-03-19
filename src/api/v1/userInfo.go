package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrips/src/db"
	"scrips/src/model"
	"strconv"
	"time"
)

/**
获取用户信息
*/
func GetUserInfo(c *gin.Context) {
	token := c.Query("token")
	var tokenModel = model.ParseToken(token)
	if tokenModel == nil {
		c.JSON(200, gin.H{
			"msg":  "缺少必要参数",
			"code": "10007",
		})
		return
	}

	user := db.SearchUserById(tokenModel.UserId);

	c.JSON(200, gin.H{
		"msg":  "登录成功",
		"code": "200",
		"user": user,
	})
}

/**
编辑用户信息
*/
func EditUserInfo(c *gin.Context) {
	token := c.Query("token")
	name := c.Query("name")
	sex := c.Query("sex")
	area := c.Query("area")

	form, err := c.MultipartForm()
	var paths []string
	if err == nil {
		files := form.File["files[]"]

		for index, file := range files {
			/// 图片保存
			path := fmt.Sprintf("src/static/images/%d_%d.png", time.Now().UnixNano(), index)
			paths = append(paths, path)
			if file != nil {
				err := c.SaveUploadedFile(file, path)
				if err != nil {
					c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
					return
				}
			}
		}
	}

	var tokenModel = model.ParseToken(token)
	if tokenModel == nil {
		c.JSON(200, gin.H{
			"msg":  "缺少必要参数",
			"code": "10007",
		})
		return
	}

	user := db.SearchUserById(tokenModel.UserId);

	if name != "" {
		user.Alias = name
	}

	if sex != ""  {
		sexInt ,_ := strconv.Atoi(sex)
		user.Sex = sexInt
	}

	if area != "" {
		user.Area = area
	}

	if len(paths) != 0 {
		user.Head = paths[0]
	}

	db.SaveUserInfo(user)

	c.JSON(200, gin.H{
		"msg":  "登录成功",
		"code": "200",
		"user": user,
	})
}
