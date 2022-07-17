package mysql

import (
	"Fever_backend/models"
	"go.uber.org/zap"
)

//查找所有的医生  并分页
func GetDoctorList(page, size int) (doctors []*models.Doctor, err error) {
	// 查看所有的文章  并分页
	db.Table("doctors").Limit(size).Offset((page - 1) * size).Find(&doctors)
	return
}

//AddDoctor 添加医生
func AddDoctor(doctor *models.Doctor) (err error) {
	// 把博客插入数据库
	db.Table("doctors").Create(map[string]interface{}{
		"hospital": doctor.Hospital, "id_number": doctor.IDNumber, "phone_number": doctor.PhoneNumber, "realname": doctor.Realname, "username": doctor.Username,
	})
	if err != nil {
		zap.L().Error("add  doctor failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}
	return
}
