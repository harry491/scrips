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
写纸条
*/
func PublishScrips(c *gin.Context) {
	token := c.Query("token")
	content := c.Query("content")

	var tokenModel = model.ParseToken(token)

	if tokenModel == nil || tokenModel.UserId == 0 {
		tokenModel = &model.Token{UserId: 0}
	}

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

	db.WriteScrips(tokenModel.UserId, content, paths)

	c.JSON(200, gin.H{
		"msg":  "成功",
		"code": "200",
	})
}

/**
所有纸条
*/
func AllScrips(c *gin.Context) {
	token := c.Query("token")
	var tokenModel = model.ParseToken(token)
	var scrips []*model.ResultScrips
	if tokenModel == nil {
		scrips = db.AllScrips(0)
	} else {
		scrips = db.AllScrips(tokenModel.UserId)
	}

	c.JSON(200, gin.H{
		"msg":    "成功",
		"code":   "200",
		"scrips": scrips,
	})
}

/**
点赞
*/
func AddPoint(c *gin.Context) {
	token := c.Query("token")
	scripId := c.Query("scripId")
	var tokenModel = model.ParseToken(token)
	if tokenModel == nil || scripId == "" {
		c.JSON(200, gin.H{
			"msg":  "缺少必要参数",
			"code": "10007",
		})
		return
	}

	scripIdInt, _ := strconv.Atoi(scripId)

	db.AddPoint(tokenModel.UserId, scripIdInt)
	c.JSON(200, gin.H{
		"msg":  "成功",
		"code": "200",
	})
}

/**
评论
*/
func AddCollect(c *gin.Context) {
	token := c.Query("token")
	scripId := c.Query("scripId")
	var tokenModel = model.ParseToken(token)
	if tokenModel == nil || scripId == "" {
		c.JSON(200, gin.H{
			"msg":  "缺少必要参数",
			"code": "10007",
		})
		return
	}

	scripIdInt, _ := strconv.Atoi(scripId)

	db.AddCollect(tokenModel.UserId, scripIdInt)
	c.JSON(200, gin.H{
		"msg":  "成功",
		"code": "200",
	})
}

/**
收藏
*/
func AllCollectScrips(c *gin.Context) {
	token := c.Query("token")
	var tokenModel = model.ParseToken(token)
	var scrips []*model.ResultScrips
	if tokenModel == nil {
		c.JSON(200, gin.H{
			"msg":    "用户未登录",
			"code":   "10009",
		})
		return
	} else {
		scrips = db.MyCollect(tokenModel.UserId)
	}

	c.JSON(200, gin.H{
		"msg":    "成功",
		"code":   "200",
		"scrips": scrips,
	})
}
