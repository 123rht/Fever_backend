package models

import (
	"encoding/json"
	"errors"
	"time"
)

type County struct {
	Card      int64     `json:"card" db:"card" gorm:"primaryKey"` // 编号
	District  string    `json:"district" db:"district"`           // 区县名称
	UpdatedAt time.Time `db:"updated_at"`
}

func (con *County) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		District string `json:"district" db:"district"` // 区县名称
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.District) == 0 {
		err = errors.New("缺少必填字段District")
	} else {
		con.District = required.District
	}
	return
}

type HospitalAdmin struct {
	Card     int64  `json:"card" db:"card"`
	Credit   uint64 `json:"credit" db:"credit"`
	Phone    uint64 `json:"phone" db:"phone"`
	ID       uint64 `json:"ID" db:"ID"`
	Role     string `json:"role" db:"role"`
	District string `json:"district" db:"district"`
	Hospital string `json:"hospital" db:"hospital"`
	Head     string `json:"head" db:"head"`
	UserName string `json:"user_name" db:"user_name"`
	Password string `json:"password" db:"password"`
	Address  string `json:"address" db:"address"`
}
