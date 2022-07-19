package logic

import (
	"Fever_backend/dao/mysql"
	"Fever_backend/models"
	"go.uber.org/zap"
)

//创建医院管理
/*
func CreateHospitalAdmin(hospitalAdmin *models.HospitalAdmin) (err error) {
	card, err := snowflake.GetID()
	if err != nil {
		zap.L().Error("snowflake.GetID() failed", zap.Error(err))
		return
	}
	hospitalAdmin.Card = card
	if err := mysql.CreatePost(hospitalAdmin); err != nil {
		zap.L().Error("mysql.CreatePost(&post) failed", zap.Error(err))
		return err
	}

	return
}*/

//修改
func UpdateDetail(card int64, hospitalAdmin *models.HospitalAdmin) error {
	return mysql.UpdateDetailByCard(card, hospitalAdmin)
}

/*func GetDetail(head string) (*models.HospitalAdmin, error) {
	return mysql.GetDetailByHead(head)
}*/

func ShowDoctor(hospital string) ([]models.Doctor, error) {
	return mysql.ShowDoctorByHospital(hospital)
}

//分页
func GetDetailList(page, size int) (data []*models.HospitalAdmin, err error) {
	details, err := mysql.GetAllList(page, size)
	if err != nil {
		zap.L().Error("mysql.GetPostList(page,size) failed", zap.Error(err))
		return
	}
	data = details
	return
}

//删除
func DeleteDetail(UserName string) error {
	return mysql.DeleteDetailByCard(UserName)
}
