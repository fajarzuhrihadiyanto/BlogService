package models

import "time"

type Article struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	AuthorId  uint      `json:"author_id" gorm:"type:int"`
	Title     string    `json:"title" validate:"required" gorm:"type:varchar(255)"`
	Content   string    `json:"content"  validate:"required" gorm:"type:text"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ArticleUpdate struct {
	ID      uint   `param:"id" validate:"required"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
