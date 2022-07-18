package mysql

import (
	"Fever_backend/models"
	"time"
)

//ChangeCountyDetailByC  修改医生
func ChangeCountyDetailByC(oldcounty string, county *models.County) (err error) {
	db.Table("counties").Where("district = ?", oldcounty).Updates(map[string]interface{}{"district": county.District, "updated_at": time.Now()})
	return err
}

//DeleteCountyDetail 删除

func DeleteCountyDetail(county string) (err error) {
	db.Table("counties").Where("district = ?", county).Delete(county)
	return err
}

//FindCountyDetail 通过区县查医院
func FindCountyDetail(county string) (hospitals []*models.HospitalAdmin, err error) {
	db.Table("hospital_admins").Where("district = ?", county).Find(&hospitals)
	return
}

//GetDistrictList 展示区县
func GetDistrictList(page, size int) (con []*models.County, err error) {
	// 查看所有的文章  并分页
	db.Table("counties").Limit(size).Offset((page - 1) * size).Find(&con)
	return
}
