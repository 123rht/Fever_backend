package controller

import (
	"Fever_backend/dao/mysql"
	"Fever_backend/logic"
	"Fever_backend/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

//创建医院管理
func CreateHospitalAdmin(c *gin.Context) {
	var hospitalAdmin models.HospitalAdmin
	var user models.User
	if err := c.ShouldBindJSON(&hospitalAdmin); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}

	err := mysql.CreatePost(&models.User{
		User_name: hospitalAdmin.UserName,
		Role:      "药店管理者",
		Password:  hospitalAdmin.Password,
	}, &models.HospitalAdmin{
		Credit:   hospitalAdmin.Credit,
		Phone:    hospitalAdmin.Phone,
		ID:       hospitalAdmin.ID,
		District: hospitalAdmin.District,
		Hospital: hospitalAdmin.Hospital,
		Head:     hospitalAdmin.Head,
		Address:  hospitalAdmin.Address,
	})
	user.User_name = hospitalAdmin.UserName
	user.Password = hospitalAdmin.Password
	if err != nil {
		zap.L().Error("logic.CreateHospitalAdmin failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

func CreateHospital(c *gin.Context) {
	var hospitalAdmin models.HospitalAdmin

	if err := c.ShouldBindJSON(&hospitalAdmin); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}

	err := mysql.CreatePostSecond(&models.User{
		User_name: hospitalAdmin.UserName,
		Role:      "药店管理员",
		Password:  1234456,
	})
	if err != nil {
		zap.L().Error("logic.CreateHospitalAdmin failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

//获取所有医院管理
func DetailHandle(c *gin.Context) {
	data, err := logic.GetDetailList(getPageInfo(c))
	if err != nil {
		zap.L().Error("logic.GetBlogList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)
}

//修改医院管理信息
func UpdateDetailHandle(c *gin.Context) {
	idStr := c.Param("card")
	fmt.Println(idStr)
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		return
	}
	var hospitalAdmin models.HospitalAdmin
	if err := c.ShouldBindJSON(&hospitalAdmin); err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}

	err = logic.UpdateDetail(id, &hospitalAdmin)
	if err != nil {
		zap.L().Error("mysql.UpdateDetail() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

//删除医院管理信息
func DeleteDetailHandle(c *gin.Context) {
	idStr := c.Param("user_name")
	/*fmt.Println(idStr)
	id, err := strconv.Parse(idStr)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		return
	}*/

	err := logic.DeleteDetail(idStr)
	if err != nil {
		zap.L().Error("mysql.DeleteDetail() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

func ShowDoctorHandle(c *gin.Context) {
	hospital := c.Param("hospital")
	fmt.Println(hospital)

	data, err := logic.ShowDoctor(hospital)
	if err != nil {
		zap.L().Error("mysql.GetUserDetail() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
