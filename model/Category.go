package model

type Category struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name" gorm:"type:varchar(50);not null;unique"`
	CreateAt Time   `json:"create_at" gorm:"type:timestamp"`
	UpdateAt Time   `json:"update_at" gorm:"type:timestamp"`
}
