package initialize

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"school/global"
	"school/model"
)

func DB() {
	dsn := "root:Qq/2013XiaoKUang@tcp(127.0.0.1:3306)/school?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = global.DB.AutoMigrate(&model.CourseInfo{})
	if err != nil {
		panic(err)
	}
}
