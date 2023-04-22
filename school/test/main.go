package main

import (
	"fmt"
	"net/http"
)

func main() {
	res, err := http.Get("http://127.0.0.1:8080/v1/course/list")
	if err != nil {
		panic(err)
	}

	str := make([]byte, 2000)

	res.Body.Read((str))
	fmt.Println(string(str))

	//dsn := "root:Qq/2013XiaoKUang@tcp(127.0.0.1:3306)/school?charset=utf8mb4&parseTime=True&loc=Local"
	//
	//var err error
	//global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	panic(err)
	//}

	//err = global.DB.AutoMigrate(&model.CourseInfo{})
	//if err != nil {
	//	panic(err)
	//}

	//for i := 1; i < 10; i++ {
	//	dao.CreateCourse(&model.CourseInfo{
	//		Name:         fmt.Sprintf("数据库原理%d", i),
	//		CourseId:     i,
	//		Teacher:      "荣老师",
	//		StudentTotal: 45,
	//		Time:         "第1-2节",
	//		ClassRoom:    "D7尚美楼308",
	//		ClassRoomId:  i / 2,
	//		Week:         "第1-16周",
	//		IsBiWeek:     i % 2,
	//		WeekNums:     i / 2,
	//	})
	//}

	//res := dao.GetCourseList()
	//for _, v := range *res {
	//	fmt.Println(v)
	//}
	//
	//res := dao.GetCourse(1, 1)
	//for _, v := range *res {
	//	fmt.Println(v)
	//}

	//res := dao.GetCourseByRoom(3)
	//fmt.Println(res)

	//err = global.DB.AutoMigrate(&model.CourseInfo{})
	//if err != nil {
	//	panic(err)
	//}

}
