package models

import "time"

type Doctor struct {
	Card        int64     `json:"card" db:"card" gorm:"primaryKey"` // 编号
	Hospital    string    `json:"hospital" db:"hospital"`           // 所属药店
	IDNumber    string    `json:"id_number" db:"id_number"`         // 身份证号
	PhoneNumber string    `json:"phone_number" db:"phone_number"`   // 电话号码
	Realname    string    `json:"realname" db:"realname"`           // 真实姓名
	Username    string    `json:"username" db:"username"`           // 用户名
	Password    string    `json:"password" db:"password"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
