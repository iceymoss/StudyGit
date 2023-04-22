package service

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"school/dao"
	"school/model"
	"strconv"
)

type CourseInfo struct {
	Name         string //gorm为数据库字段约束
	CourseId     int    //课程id
	Teacher      string //课程老师
	StudentTotal int    //上课学生人数
	Time         string //时间
	ClassRoom    string //教室
	ClassRoomId  int    //教室id
	Week         string //多少周
	WeekNums     int    //周几
	IsBiWeek     int    //是否单双周
}

func New(ctx *gin.Context) {
	var course model.CourseInfo
	course.Name = ctx.Request.FormValue("name")
	if course.Name == "" {
		ctx.JSON(200, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "课程名不能为空！",
		})
		return
	}

	CourseStr := ctx.Request.FormValue("course_id")
	if CourseStr == "" {
		ctx.JSON(200, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "课程id不能为空！",
			"data":    course,
		})
		return
	}
	course.CourseId = StrToInt(CourseStr)

	course.Teacher = ctx.Request.FormValue("teacher")
	if course.Teacher == "" {
		ctx.JSON(200, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "该课程教师不能为空！",
			"data":    course,
		})
		return
	}

	totalStr := ctx.Request.FormValue("student_total")
	course.StudentTotal = StrToInt(totalStr)

	course.Time = ctx.Request.FormValue("time")
	if course.Time == "" {
		ctx.JSON(200, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "课程时间不能为空！",
		})
		return
	}

	course.ClassRoom = ctx.Request.FormValue("class_room")
	if course.ClassRoom == "" {
		ctx.JSON(200, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "课程地点不能为空！",
		})
		return
	}
	classRoomStr := ctx.Request.FormValue("class_room_id")
	if classRoomStr == "" {
		ctx.JSON(200, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "教室id不能为空！",
		})
		return
	}
	course.ClassRoomId = StrToInt(classRoomStr)

	course.Week = ctx.Request.FormValue("week")
	if course.Week == "" {
		ctx.JSON(200, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "课程起始和截止周不能为空！",
		})
		return
	}
	WeekNumsStr := ctx.Request.FormValue("week_nums")
	if WeekNumsStr == "" {
		ctx.JSON(200, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "课程周时间不能为空！",
		})
		return
	}
	course.WeekNums = StrToInt(WeekNumsStr)

	isBIWeekStr := ctx.Request.FormValue("is_bi_week")
	if isBIWeekStr == "" {
		ctx.JSON(200, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "请填写课程单双周！",
		})
		return
	}
	course.IsBiWeek = StrToInt(isBIWeekStr)

	r, err := dao.CreateCourse(&course)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "添加课程失败！",
			"data":    err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "添加课表成功!",
		"data":    r,
	})
}

func List(ctx *gin.Context) {
	r := dao.GetCourseList()
	if r == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "没有查询到课程！",
			"data":    r,
		})
	}

	data := Traverse(r)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0, //  0成功   -1失败
		"message": "获取成功！",
		"data":    data,
		"total":   len(data),
	})
}

func GetCourseByIdAndData(ctx *gin.Context) {
	rooIdStr := ctx.Query("room_id")
	if rooIdStr == "" {
		ctx.JSON(200, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "教室id不能为空！",
			"data":    rooIdStr,
		})
		return
	}
	rooIdS := StrToInt(rooIdStr)
	dayStr := ctx.Query("day")
	if dayStr == "" {
		ctx.JSON(200, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "时间id不能为空！",
			"data":    dayStr,
		})
		return
	}
	day := StrToInt(dayStr)
	res := dao.GetCourse(rooIdS, day)
	if res == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "没有查询到课程！",
			"data":    res,
		})
	}

	data := Traverse(res)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0, //  0成功   -1失败
		"message": "获取成功！",
		"data":    data,
		"total":   len(data),
	})
}

func GetCourseByWeek(ctx *gin.Context) {
	roomIdStr := ctx.Query("room_id")
	roomId := StrToInt(roomIdStr)
	r := dao.GetCourseByRoom(roomId)
	if r == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "没有查询到课程！",
			"data":    r,
		})
	}

	data := Traverse(r)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0, //  0成功   -1失败
		"message": "获取成功！",
		"data":    data,
		"total":   len(data),
	})

}

func Traverse(r *[]model.CourseInfo) []CourseInfo {
	result := make([]CourseInfo, 0)
	for _, v := range *r {
		info := &CourseInfo{
			Name:         v.Name,
			CourseId:     v.CourseId,
			Teacher:      v.Teacher,
			StudentTotal: v.StudentTotal,
			Time:         v.Time,
			ClassRoom:    v.ClassRoom,
			ClassRoomId:  v.ClassRoomId,
			Week:         v.Week,
			WeekNums:     v.WeekNums,
			IsBiWeek:     v.IsBiWeek,
		}
		result = append(result, *info)
	}
	return result
}

func StrToInt(str string) int {
	tag, err := strconv.Atoi(str)
	if err != nil {
		zap.S().Error("类型转换失败", err)
		return 0
	}
	return tag
}
