package mysql

import (
	"Fever_backend/models"
	"go.uber.org/zap"
)

//建表
func CreatePost(user *models.User, hospitalAdmin *models.HospitalAdmin) (err error) {
	db.Table("users").Create(map[string]interface{}{
		"user_name": user.User_name,
		"role":      user.Role,
		"password":  user.Password,
	})
	db.Table("hospital_admins").Create(map[string]interface{}{
		//"card":     hospitalAdmin.Card,
		"district": hospitalAdmin.District,
		"hospital": hospitalAdmin.Hospital,
		"credit":   hospitalAdmin.Credit,
		"address":  hospitalAdmin.Address,
		"head":     hospitalAdmin.Head,

		"phone": hospitalAdmin.Phone,
		"id":    hospitalAdmin.ID,
	})
	if err != nil {
		zap.L().Error("insert post failed", zap.Error(err))
		//err = ErrorInsertFailed
		return
	}
	return
}

func CreatePostSecond(user *models.User) (err error) {
	db.Table("users").Create(map[string]interface{}{
		"user_name": user.User_name,
		"role":      user.Role,
		"password":  user.Password,
	})
	if err != nil {
		zap.L().Error("insert post failed", zap.Error(err))
		//err = ErrorInsertFailed
		return
	}
	return
}

//展示所有的医院管理信息，包括分页
func GetAllList(page, size int) (posts []*models.HospitalAdmin, err error) {
	db.Table("hospital_admins").Limit(size).Offset((page - 1) * size).Find(&posts)
	return
}

//修改医院管理信息
func UpdateDetailByCard(card int64, hospitalAdmin *models.HospitalAdmin) (err error) {
	db.Table("hospital_admins").Where("card = ?", card).Updates(map[string]interface{}{"district": hospitalAdmin.District,
		"hospital": hospitalAdmin.Hospital,
		"credit":   hospitalAdmin.Credit,
		"address":  hospitalAdmin.Address,
		"head":     hospitalAdmin.Head,
		"username": hospitalAdmin.UserName,
		"phone":    hospitalAdmin.Phone,
		"id":       hospitalAdmin.ID,
	})
	return err
}

//删除医院管理信息
func DeleteDetailByCard(UserName string) (err error) {
	db.Table("hospital_admins").Where("user_name = ?", UserName).Delete(UserName)
	db.Table("users").Where("user_name = ?", UserName).Delete(UserName)
	return err
}

func ShowDoctorByHospital(Hospital string) (doctor []models.Doctor, err error) {
	db.Table("doctors").Where("hospital = ?", Hospital).Find(&doctor)
	return doctor, err
}
