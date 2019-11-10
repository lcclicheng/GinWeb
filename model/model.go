package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

type Stu struct {
	gorm.Model
	Name     string
	PassWord string
}

var GlobalDB *gorm.DB

func InitModel() {
	//打开数据库
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/ihome")
	if err != nil {
		fmt.Println("连接数据库失败")
		return
	}
	//连接池设置
	db.DB().SetMaxIdleConns(20)//最大空闲数量
	db.DB().SetMaxOpenConns(30)//最大打开数量
	db.DB().SetConnMaxLifetime(60 * 30)//最大生命周期

	//db.SingularTable(true)  创建单数表
	GlobalDB = db

	//自动迁移
	db.AutoMigrate(new(Stu))
}
