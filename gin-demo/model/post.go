package model

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/uuid"
	"time"
)

type Post struct {
	ID			uuid.UUID	`json:"id,omitempty" gorm:"type:char(36);primary_key"`
	UserId		uint		`json:"user_id,omitempty" gorm:"not null"`
	CategoryId	uint		`json:"category_id,omitempty" gorm:"not null"`
	Category	*Category
	Title		string		`json:"title,omitempty" gorm:"type:varchar(50);not null"`
	HeadImg		string		`json:"head_img,omitempty"`
	Content		string		`json:"content,omitempty" gorm:"type:timestamp"`
	CreateAt	time.Time	`json:"create_at,omitempty" gorm:"type:timestamp"`
	UpdateAt	time.Time	`json:"update_at,omitempty" gorm:"type:timestamp"`
}

func (post *Post) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4())
}
