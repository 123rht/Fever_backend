package models

import (
	"encoding/json"
	"errors"
	"time"
)

type Doctor struct {
	Card        int64     `json:"card" db:"card" gorm:"primaryKey"` // 编号
	Hospital    string    `json:"hospital" db:"hospital"`           // 所属药店
	IDNumber    string    `json:"id_number" db:"id_number"`         // 身份证号
	PhoneNumber string    `json:"phone_number" db:"phone_number"`   // 电话号码
	Realname    string    `json:"realname" db:"realname"`           // 真实姓名
	Username    string    `json:"username" db:"username"`           // 用户名
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func (d *Doctor) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		Hospital    string `json:"hospital" db:"hospital"`         // 所属药店
		IDNumber    string `json:"id_number" db:"id_number"`       // 身份证号
		PhoneNumber string `json:"phone_number" db:"phone_number"` // 电话号码
		Realname    string `json:"realname" db:"realname"`         // 真实姓名
		Username    string `json:"username" db:"username"`         // 用户名
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.Hospital) == 0 {
		err = errors.New("所属药店不能为空")
	} else if len(required.IDNumber) != 18 {
		err = errors.New("身份证号格式错误")
	} else if len(required.PhoneNumber) != 11 {
		err = errors.New("手机号格式错误")
	} else if len(required.Realname) == 0 {
		err = errors.New("真是姓名不能为空")
	} else if len(required.Realname) == 0 {
		err = errors.New("用户名不能为空")
	} else {
		d.Hospital = required.Hospital
		d.IDNumber = required.IDNumber
		d.PhoneNumber = required.PhoneNumber
		d.Realname = required.Realname
		d.Username = required.Username
	}
	return
}
