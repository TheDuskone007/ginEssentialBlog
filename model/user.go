package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null" json:"name"`
	Telephone string `gorm:"type:varchar(11);not null;unique" json:"telephone"`
	Password  string `gorm:"type:varchar(110);not null" json:"password"`
}
