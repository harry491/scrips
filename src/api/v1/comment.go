package v1

import (
	"github.com/gin-gonic/gin"
	"scrips/src/db"
	"scrips/src/model"
	"strconv"
)

/**
所有评论
*/
func AllComments(c *gin.Context) {
	scripId := c.Query("scripId")
	token := c.Query("token")

	if scripId == "" {
		c.JSON(200, gin.H{
			"msg":  "未知的纸条",
			"code": "10010",
		})
		return
	}

	var tokenModel = model.ParseToken(token)
	var comments []*model.CommentStruct
	if tokenModel == nil {
		c.JSON(200, gin.H{
			"msg":  "用户未登录",
			"code": "10009",
		})
		return
	} else {
		scripIdInt, _ := strconv.Atoi(scripId)
		comments = db.ScripComments(scripIdInt, tokenModel.UserId)
	}

	c.JSON(200, gin.H{
		"msg":      "成功",
		"code":     "200",
		"comments": comments,
	})
}

/**
写评论
*/
func AddComment(c *gin.Context) {
	content := c.Query("content")
	scripId := c.Query("scripId")
	token := c.Query("token")

	if content == "" {
		c.JSON(200, gin.H{
			"msg":  "评论为空",
			"code": "10010",
		})
		return
	}

	if scripId == "" {
		c.JSON(200, gin.H{
			"msg":  "未知的纸条",
			"code": "10010",
		})
		return
	}

	var tokenModel = model.ParseToken(token)
	if tokenModel == nil {
		c.JSON(200, gin.H{
			"msg":  "用户未登录",
			"code": "10009",
		})
		return
	} else {
		scripIdInt, _ := strconv.Atoi(scripId)
		db.WriteComment(scripIdInt, tokenModel.UserId, content)
	}

	c.JSON(200, gin.H{
		"msg":  "成功",
		"code": "200",
	})
}

/**
为评论点赞
*/
func PointComment(c *gin.Context) {
	commentId := c.Query("commentId")
	token := c.Query("token")

	if commentId == "" {
		c.JSON(200, gin.H{
			"msg":  "未知的评论",
			"code": "10010",
		})
		return
	}

	var tokenModel = model.ParseToken(token)
	if tokenModel == nil {
		c.JSON(200, gin.H{
			"msg":  "用户未登录",
			"code": "10009",
		})
		return
	} else {
		commentIdInt, _ := strconv.Atoi(commentId)
		db.PointToComment(commentIdInt, tokenModel.UserId)
	}

	c.JSON(200, gin.H{
		"msg":  "成功",
		"code": "1",
	})
}
