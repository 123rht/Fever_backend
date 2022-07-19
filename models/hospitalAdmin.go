package models

import (
	"encoding/json"
	"errors"
)

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
	Password uint64 `json:"password" db:"password"`
	Address  string `json:"address" db:"address"`
}

func (c *HospitalAdmin) UnmarshallJSON(data []byte) (err error) {
	required := struct {
		Hospital string `json:"business" db:"business"`
		Credit   string `json:"credit" db:"credit"`
		Address  string `json:"address" db:"address"`
		Head     string `json:"head" db:"hesd"`
		UserName string `json:"user_name" db:"user_name"`
		Phone    string `json:"phone" db:"phone"`
		ID       string `json:"ID" db:"ID"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.Hospital) == 0 {
		err = errors.New("企业名称不能为空")
	} else if len(required.Credit) == 0 {
		err = errors.New("社会信用代码不能为空")
	} else if len(required.Address) == 0 {
		err = errors.New("注册地址不能为空")
	} else if len(required.Head) == 0 {
		err = errors.New("负责人不能为空")
	} else if len(required.Phone) == 0 {
		err = errors.New("负责人电话不能为空")
	} else if len(required.ID) == 0 {
		err = errors.New("负责人身份证不能为空")
	} else if len(required.UserName) == 0 {
		err = errors.New("负责人用户名不能为空")
	}
	return
}
