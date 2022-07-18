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
		v2.POST("add_doc", doctor.AddDoctorHandler)
		//修改医生的信息
		v2.POST("change_doc", doctor.ChangeDoctorHandler)
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
