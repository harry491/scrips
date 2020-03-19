package model

import "github.com/jinzhu/gorm"

type ResultScrips struct {
	gorm.Model
	UserId        int
	Content       string
	PointNumber   int
	CollectNumber int
	Tag           int
	IsPoint       bool
	IsCollect     bool
	Alias         string
	Sex           int
	Head          string
	Images        string
}

type Scrips struct {
	gorm.Model
	UserId        int
	Content       string
	PointNumber   int
	CollectNumber int
	Tag           int
	Images        string
}

type Point struct {
	gorm.Model
	UserId  int
	IsPoint bool
	ScripId int
}

type Collect struct {
	gorm.Model
	UserId    int
	IsCollect bool
	ScripId   int
}
