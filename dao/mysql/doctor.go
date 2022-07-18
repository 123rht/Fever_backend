package mysql

import (
	"Fever_backend/models"
	"go.uber.org/zap"
	"time"
)

//查找所有的医生  并分页
func GetDoctorList(username string, page, size int) (doctors []*models.Doctor, err error) {
	// 查看所有的文章  并分页
	db.Table("doctors").Not("username = ?", username).Limit(size).Offset((page - 1) * size).Find(&doctors)
	return
}

//AddDoctor 添加医生
func AddDoctor(doctor *models.Doctor) (err error) {
	u := db.Table("doctors").Where("username = ?", doctor.Username).Find(doctor)
	if u.RowsAffected > 0 {
		// 用户已存在
		return ErrorUserExit
	}
	// 把博客插入数据库
	db.Table("doctors").Create(map[string]interface{}{
		"hospital": doctor.Hospital, "id_number": doctor.IDNumber, "phone_number": doctor.PhoneNumber, "realname": doctor.Realname, "username": doctor.Username, "created_at": time.Now(),
	})
	if err != nil {
		zap.L().Error("add  doctor failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}
	return
}

//ChangeDoctorDetailByUserName  修改医生
func ChangeDoctorDetailByUserName(doctor *models.Doctor) (err error) {
	db.Table("doctors").Where("username = ?", doctor.Username).Updates(map[string]interface{}{"hospital": doctor.Hospital, "id_number": doctor.IDNumber, "phone_number": doctor.PhoneNumber, "realname": doctor.Realname, "updated_at": time.Now()})
	return err
}

//DeleteDoctorDetailByUserName 删除医生
func DeleteDoctorDetailByUserName(username string) (err error) {
	db.Table("doctors").Where("username = ?", username).Delete(username)
	return err
}

//UpdateMyMessageByUserName 修改医生管理员信息
func UpdateMyMessageByUserName(username string, doctor *models.Doctor) (err error) {
	db.Table("doctors").Where("username = ?", username).Updates(map[string]interface{}{"hospital": doctor.Hospital, "id_number": doctor.IDNumber, "phone_number": doctor.PhoneNumber, "realname": doctor.Realname, "updated_at": time.Now()})
	return err
}
