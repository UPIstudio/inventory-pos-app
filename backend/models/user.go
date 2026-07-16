package models

import (
	"time"
)

type User struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Username    string    `json:"username" gorm:"unique;not null; size:255"`
	Password    string    `json:"password" gorm:"not null; size:255"`
	Role        string    `json:"role" gorm:"default:'staff'; size:50"`
	ActiveToken string    `json:"active_token" gorm:"type:text"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
