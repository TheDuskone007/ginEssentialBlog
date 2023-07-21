package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	ID         uuid.UUID `gorm:"type:char(36);primarykey;"`
	UserId     uint      `json:"user_id" gorm:"not null"`
	CategoryId uint      `json:"categoryId" gorm:"not null"`
	Category   *Category
	Title      string `json:"title" gorm:"type:varchar(50);not null"`
	//存储图片地址
	HeadImg   string `json:"head_img"`
	Content   string `json:"content" gorm:"type:text;not null"`
	CreatedAt Time   `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt Time   `json:"updated_at" gorm:"type:timestamp"`
}

func (p *Post) BeforeCreate(tx *gorm.DB) error {
	//id, err := uuid.NewRandom()
	//p.ID = MYTUUID(id)
	//return err

	id, err := uuid.NewRandom()
	p.ID = id
	return err
}
