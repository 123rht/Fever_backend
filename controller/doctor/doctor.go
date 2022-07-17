package doctor

import (
	"Fever_backend/logic"
	"Fever_backend/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//DoctorListHandler 获取医生名单的处理函数
func DoctorListHandler(c *gin.Context) {
	//获取数据
	data, err := logic.GetDoctorList(getPageInfo(c))
	if err != nil {
		zap.L().Error("Logic.GetDoctorList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回响应
	ResponseSuccess(c, data)
}

//AddDoctorHandler  添加医生
func AddDoctorHandler(c *gin.Context) {
	//获取参数
	var doctor models.Doctor
	if err := c.ShouldBindJSON(&doctor); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		return
	}

	err := logic.AddDoctor(&doctor)
	if err != nil {
		zap.L().Error("logic.AddDoctor failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
