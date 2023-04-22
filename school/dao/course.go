package dao

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"log"
	"school/global"
	"school/model"
)

// GetCourseList 查询课程
func GetCourseList() *[]model.CourseInfo {
	var courses []model.CourseInfo

	res := global.DB.Find(&courses)
	if res.RowsAffected == 0 {
		zap.S().Infof("没有课程")
		return nil
	}
	return &courses
}

// GetCourse 查询某教室一天的课程
func GetCourse(id int, weekNum int) *[]model.CourseInfo {
	var courses []model.CourseInfo
	fmt.Println("id:", id, "week:", weekNum)
	if res := global.DB.Where("class_room_id= ? and week_nums= ?", id, weekNum).Find(&courses); res.RowsAffected == 0 {
		zap.S().Infof("没有课程")
		return nil
	}
	return &courses
}

// GetCourseByRoom 查询教室整周的课程
func GetCourseByRoom(id int) *[]model.CourseInfo {
	var courses []model.CourseInfo
	if res := global.DB.Where("class_room_id= ?", id).Find(&courses); res.RowsAffected == 0 {
		zap.S().Infof("没有课程")
		return nil
	}
	return &courses
}

func CreateCourse(info *model.CourseInfo) (int, error) {
	var course model.CourseInfo
	tx := global.DB.Where("class_room_id= ? and time= ? and week_nums= ?", info.ClassRoomId, info.Time, info.WeekNums).Find(&course)
	if tx.RowsAffected > 0 {
		zap.S().Infof("已有课程存在")
		return 0, errors.New("已有课程存在")
	}
	res := global.DB.Create(&info)
	if res.Error != nil {
		log.Fatalln("create course error")
		return 0, errors.New("create course error")
	}
	return int(info.ID), nil
}
