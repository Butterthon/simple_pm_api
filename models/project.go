package models

import "time"

// Project 登録フォーム
type Project struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt" sql:"index"`
	UUID      string     `json:"uuid" gorm:"default: gen_random_uuid()"`
	Version   int        `json:"version" gorm:"default: 1"`
	Name      string     `json:"name" binding:"required,max=100"`
}
