package db

import (
	"github.com/jinzhu/gorm"
	"scrips/src/model"
	"strings"
)

/**
写纸条
*/
func WriteScrips(userId int, content string, files []string) {

	for _, path := range files {
		DB.Create(&model.Image{Path: path})
	}

	scrips := &model.Scrips{UserId: userId, Content: content, Images: strings.Join(files, ",")}
	DB.Create(scrips)
}

/**
查询所有纸条
*/
func AllScrips(userId int) []*model.ResultScrips {
	var rscrips []*model.ResultScrips

	if userId != 0 {
		DB.Table("scrips").Order("scrips.id desc").Select("scrips.id, scrips.content,scrips.created_at, scrips.point_number,scrips.collect_number,scrips.images,users.alias,users.sex,users.head , collects.is_collect,collects.user_id , points.is_point").Joins("left join users on users.id = scrips.user_id").Joins("left join collects on collects.scrip_id = scrips.id and ISNULL(collects.deleted_at)  and collects.user_id = ?", userId).Joins("left join points on points.scrip_id = scrips.id and ISNULL(points.deleted_at) and points.user_id = ?", userId).Limit(20).Scan(&rscrips)
	} else {
		DB.Table("scrips").Order("scrips.id desc").Select("scrips.id, scrips.content,scrips.created_at, scrips.point_number,scrips.collect_number,scrips.images ,users.alias,users.sex,users.head").Joins("left join users on users.id = scrips.user_id").Limit(20).Find(&rscrips)
	}
	return rscrips
}

/**
点赞
*/
func AddPoint(userId int, scripId int) {

	point := model.Point{UserId: userId, ScripId: scripId}

	result := DB.Where(point).Find(&model.Point{}).RecordNotFound()

	scrips := model.Scrips{}
	scrips.ID = uint(scripId)

	if result {
		DB.Model(&scrips).Update("point_number", gorm.Expr("point_number + 1"))
		point.IsPoint = true
		DB.Create(&point)
	} else {
		DB.Model(&scrips).Update("point_number", gorm.Expr("point_number - 1"))
		point.IsPoint = false
		DB.Delete(&point)
	}
}

/**
收藏
*/
func AddCollect(userId int, scripId int) {

	collect := model.Collect{UserId: userId, ScripId: scripId}

	result := DB.Where(collect).Find(&model.Collect{}).RecordNotFound()

	scrips := model.Scrips{}
	scrips.ID = uint(scripId)
	if result {
		DB.Model(&scrips).Update("collect_number", gorm.Expr("collect_number + 1"))
		collect.IsCollect = true
		DB.Create(&collect)
	} else {
		DB.Model(&scrips).Update("collect_number", gorm.Expr("collect_number - 1"))
		collect.IsCollect = false
		DB.Delete(&collect)
	}
}

/**
我的收藏
*/
func MyCollect(userId int) []*model.ResultScrips {
	var rscrips []*model.ResultScrips
	DB.Table("collects").Select("collects.is_collect ,points.is_point , scrips.content , scrips.id , scrips.created_at ,scrips.images , scrips.content , scrips.updated_at, scrips.id , scrips.user_id , scrips.point_number , scrips.collect_number ").Where("collects.user_id = ? and ISNULL(collects.deleted_at) ", userId).Joins("left join scrips on scrips.id = collects.scrip_id AND collects.deleted_at IS NULL and collects.user_id = ?", userId).Joins("left join points on points.scrip_id = collects.scrip_id and points.deleted_at IS NULL and points.user_id = ?", userId).Limit(20).Scan(&rscrips)
	return rscrips
}
