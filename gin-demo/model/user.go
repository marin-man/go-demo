package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username	string	`gorm:"varchar(25);not null" json:"name,omitempty"`
	Password	string	`gorm:"varchar(25);not null" json:"password,omitempty"`
	Phone		string	`gorm:"size:255;not null" json:"phone,omitempty"`
}
