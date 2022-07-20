package logic

import (
	"Fever_backend/dao/mysql"
	"Fever_backend/models"
	"go.uber.org/zap"
)

//ChangeCountyDetail 修改区县信息
func ChangeCountyDetail(oldcounty string, con *models.County) error {
	return mysql.ChangeCountyDetailByC(oldcounty, con)
}

//DeleteCountyDetail 删除区县

func DeleteCountyDetail(county string) error {
	return mysql.DeleteCountyDetail(county)
}

//FindCountyList  查找区县的所有医院
func FindCountyList(county string) (data []*models.HospitalAdmin, err error) {
	hospitals, err := mysql.FindCountyDetail(county)
	if err != nil {
		zap.L().Error("mysql.GetDoctorList(page,size) failed", zap.Error(err))
		return
	}
	data = hospitals
	return
}

//GetDistrictList  展示区县
func GetDistrictList(page int, size int) (data []*models.County, err error) {
	districts, err := mysql.GetDistrictList(page, size)
	if err != nil {
		zap.L().Error("mysql.GetDistrictList(page,size) failed", zap.Error(err))
		return
	}
	data = districts
	return
}

//MyDistrictHosList 展示本区所有医院
func MyDistrictHosList(username string, page int, size int) (data []*models.HospitalAdmin, err error) {
	doctors, err := mysql.MyDistrictHosList(username, page, size)
	if err != nil {
		zap.L().Error("mysql.MyDistrictHosList(username,page,size) failed", zap.Error(err))
		return
	}
	data = doctors
	return
}

//MyDistrictDocList 展示本区所有医生
func MyDistrictDocList(username string, page int, size int) (data []*models.Doctor, err error) {
	doctors, err := mysql.MyDistrictDocList(username, page, size)
	if err != nil {
		zap.L().Error("mysql.MyDistrictDocList(username,page,size) failed", zap.Error(err))
		return
	}
	data = doctors
	return
}
