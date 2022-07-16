package routers

import (
	"Fever_backend/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/fever/v1")
	v1.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "I'll love you till I die")
	})
	v1.POST("/login", controller.LoginHandler)
	//登录使用
	v1.Use(controller.JWTAuthMiddleware())
	{
		v1.POST("/change_password", controller.ChangePasswordHandler)
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
