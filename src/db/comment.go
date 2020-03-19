package db

import (
	"github.com/jinzhu/gorm"
	"scrips/src/model"
)

/**
查询纸条下评论
*/
func ScripComments(scripId int, userId int) []*model.CommentStruct {

	var comments []*model.CommentStruct

	DB.Table("comments").Select("comments.id , comments.user_id ,comments.point_number , comments.content,comments.created_at ,comment_points.is_point , users.alias , users.sex,users.head").Where("comments.scrip_id = ?", scripId).Joins("left join users on users.id = comments.user_id").Joins("left join comment_points on comment_points.comment_id = comments.id and comment_points.user_id = ?", userId).Order("comments.id desc").Limit(20).Scan(&comments)

	return comments
}

/**
写评论
*/
func WriteComment(scripId int, userId int, content string) {

	comment := model.Comment{ScripId: scripId, UserId: userId, Content: content}

	DB.Create(&comment)
}

/**
评论点赞
*/
func PointToComment(commentId int, userId int) {

	point := model.CommentPoint{CommentId: commentId, UserId: userId}
	resultPoint := model.CommentPoint{}
	result := DB.Where(&point).Find(&resultPoint).RecordNotFound()
	comment := model.Comment{}
	comment.ID = uint(commentId)
	if result {
		point.IsPoint = true
		DB.Create(&point)
		DB.Model(&comment).Update("point_number", gorm.Expr("point_number + 1"))
	} else {
		DB.Model(&point).Update("is_point", gorm.Expr("! is_point"))
		if resultPoint.IsPoint {
			DB.Model(&comment).Update("point_number", gorm.Expr("point_number - 1"))
		} else {
			DB.Model(&comment).Update("point_number", gorm.Expr("point_number + 1"))
		}
	}
}
