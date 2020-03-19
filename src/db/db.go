package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"scrips/src/model"
	"time"
)

const (
	USERNAME = "root"
	PASSWORD = "123456"
	NETWORK  = "tcp"
	SERVER   = "localhost"
	PORT     = 3306
	DATABASE = "dev_db"
)

var DB *gorm.DB

/**
打开数据库
*/

func OpenGormDb() *gorm.DB {
	dns := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?parseTime=true", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	db, err := gorm.Open("mysql", dns)
	//defer db.Close()
	if err != nil {
		fmt.Println("open error:%v\n", err)
	}

	db.DB().SetConnMaxLifetime(100 * time.Second)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetMaxIdleConns(16)

	return db
}

/**
初始化tables
 */
func InitTables()  {

	DB = OpenGormDb()

	if DB.HasTable(&model.SmsModel{}) == false {
		DB.CreateTable(&model.SmsModel{})
	}

	if DB.HasTable(&model.User{}) == false {
		DB.CreateTable(&model.User{})
	}

	if DB.HasTable(&model.Scrips{}) == false {
		DB.CreateTable(&model.Scrips{})
	}

	if DB.HasTable(&model.Point{}) == false {
		DB.CreateTable(&model.Point{})
	}

	if DB.HasTable(&model.Collect{}) == false {
		DB.CreateTable(&model.Collect{})
	}

	if DB.HasTable(&model.Image{}) == false {
		DB.CreateTable(&model.Image{})
	}

	if DB.HasTable(&model.Comment{}) == false {
		DB.CreateTable(&model.Comment{})
	}

	if DB.HasTable(&model.CommentPoint{}) == false {
		DB.CreateTable(&model.CommentPoint{})
	}
}
