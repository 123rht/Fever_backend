package logic

import (
	"Fever_backend/dao/mysql"
	"Fever_backend/models"
	"go.uber.org/zap"
)

//创建医院管理
func AddHospital(up *models.UP, hospital *models.HospitalAdmin) (err error) {
	if err := mysql.AddHospital(up, hospital); err != nil {
		zap.L().Error("mysql.AddDoctor(doctor) failed", zap.Error(err))
		return err
	}
	return
}

//修改
func UpdateDetail(user string, hospitalAdmin *models.Hospital) error {
	return mysql.UpdateDetailByCard(user, hospitalAdmin)
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
