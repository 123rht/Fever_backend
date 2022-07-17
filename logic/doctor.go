package logic

import (
	"Fever_backend/dao/mysql"
	"Fever_backend/models"
	"go.uber.org/zap"
)

//GetDoctorList  获取医生名单列表
func GetDoctorList(page, size int) (data []*models.Doctor, err error) {
	doctors, err := mysql.GetDoctorList(page, size)
	if err != nil {
		zap.L().Error("mysql.GetDoctorList(page,size) failed", zap.Error(err))
		return
	}
	data = doctors
	return
}

//AddDoctor  添加医生
func AddDoctor(doctor *models.Doctor) (err error) {
	if err := mysql.AddDoctor(doctor); err != nil {
		zap.L().Error("mysql.AddDoctor(&post) failed", zap.Error(err))
		return err
	}
	return
}
