package logic

import (
	"Fever_backend/dao/mysql"
	"Fever_backend/models"
	"go.uber.org/zap"
)

//GetMyHosDoctorList 获取本医院所有的医生
func GetMyHosDoctorList(username string, page int, size int) (data []*models.Doctor, err error) {
	doctors, err := mysql.GetMyHosDoctorList(username, page, size)
	if err != nil {
		zap.L().Error("mysql.GetMyHosDoctorList(username,page,size) failed", zap.Error(err))
		return
	}
	data = doctors
	return
}

//GetDoctorList  获取医生名单列表
func GetDoctorList(username string, page int, size int) (data []*models.Doctor, err error) {
	doctors, err := mysql.GetDoctorList(username, page, size)
	if err != nil {
		zap.L().Error("mysql.GetDoctorList(page,size) failed", zap.Error(err))
		return
	}
	data = doctors
	return
}

//AddDoctor  添加医生
func AddDoctor(up *models.UP, doctor *models.Doctor) (err error) {
	if err := mysql.AddDoctor(up, doctor); err != nil {
		zap.L().Error("mysql.AddDoctor(doctor) failed", zap.Error(err))
		return err
	}
	return
}

//ChangeDoctorDetail 修改医生信息
func ChangeDoctorDetail(userName string, doctor *models.Doc) error {
	return mysql.ChangeDoctorDetailByUserName(userName, doctor)
}

//DeleteDoctorDetail 删除医生
func DeleteDoctorDetail(username string) error {
	return mysql.DeleteDoctorDetailByUserName(username)
}

//UpdateMyMessage 修改医生管理员的信息
func UpdateMyMessage(username string, doctor *models.HospitalAdmin) error {
	return mysql.UpdateMyMessageByUserName(username, doctor)
}
