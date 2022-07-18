package models

type county struct {
	Card     int64  `json:"card" db:"card" gorm:"primaryKey"` // 编号
	District string `json:"district" db:"district"`           // 区县名称
}
