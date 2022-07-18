package routers

import (
	"Fever_backend/controller"
	"Fever_backend/controller/doctor"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/fever/v1")
	v1.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "I'll love you till I die")
	})
	//登录
	v1.POST("/login", controller.LoginHandler)

	//登录验证token
	v1.Use(controller.JWTAuthMiddleware())
	{
		//修改密码
		v1.POST("/change_password", controller.ChangePasswordHandler)
	}

	//医生名单相关接口
	v2 := r.Group("/doctor/v1")

	//登录验证token
	v2.Use(controller.JWTAuthMiddleware())
	{
		//查看所有医生的名单
		v2.GET("/information_all", doctor.DoctorListHandler)
		//添加医生
		v2.POST("/add_doc", doctor.AddDoctorHandler)
		//修改医生的信息
		v2.POST("/change_doc", doctor.ChangeDoctorHandler)
		//删除医生
		v2.POST("/delete_doc", doctor.DeleteDoctorHandler)
		//修改当前医生管理员的信息
		v2.POST("/update_myself", doctor.UpdateMyMessage)
	}

	//医生名单相关接口
	v3 := r.Group("/county/v1")

	//登录验证token
	v3.Use(controller.JWTAuthMiddleware())
	{
		//查看所有区县的名单
		v3.GET("/district_all", controller.DistrictListHandler)
		//修改区县的信息
		v3.POST("/change_con", controller.ChangeCountyHandler)
		//删除区县的信息
		v3.POST("/delete_con", controller.DeleteCountyHandler)
		//通过区县名查找区县的所有医院
		v3.POST("/find_con_his", controller.FindCountyHandler)

	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
