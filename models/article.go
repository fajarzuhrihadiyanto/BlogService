package models

import "time"

type Article struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	AuthorId  uint      `json:"-" gorm:"type:int"`
	Author    User      `json:"author" validate:"-"`
	Title     string    `json:"title" validate:"required" gorm:"type:varchar(255)"`
	Content   string    `json:"content"  validate:"required" gorm:"type:text"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time `json:"updated_at"`
}
