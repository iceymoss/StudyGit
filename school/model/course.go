package model

import (
	"gorm.io/gorm"
)

type CourseInfo struct {
	gorm.Model
	Name         string `gorm:"not null;type:varchar(20)"` //gorm为数据库字段约束
	CourseId     int    `gorm:"not null"`                  //课程id
	Teacher      string `gorm:"not null"`                  //课程老师
	StudentTotal int    //上课学生人数
	Time         string //时间
	ClassRoom    string //教室
	ClassRoomId  int    //教室id
	Week         string //多少周
	WeekNums     int    //周几
	IsBiWeek     int    //是否单双周
}

// UserTableName 指定表的名称
func (table *CourseInfo) UserTableName() string {
	return "course_info"
}
