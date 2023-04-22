package router

import (
	"school/middlewares"
	"school/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()
	//配置跨越
	router.Use(middlewares.Cors())

	v1 := router.Group("v1")
	course := v1.Group("course")
	{
		course.GET("/list", service.List)
		course.GET("/course_cw", service.GetCourseByIdAndData)
		course.GET("/course_id", service.GetCourseByWeek)
		course.POST("/new", service.New)
	}
	return router
}
