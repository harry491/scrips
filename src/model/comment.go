package model

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	UserId      int
	ScripId     int
	Content     string
	PointNumber int
	IsPoint     bool
}

type CommentPoint struct {
	gorm.Model
	CommentId int
	UserId    int
	IsPoint   bool
}

type CommentStruct struct {
	gorm.Model
	UserId      int
	Alias       string
	Head        string
	Sex         int
	ScripId     int
	Content     string
	PointNumber int
	IsPoint     bool
}
