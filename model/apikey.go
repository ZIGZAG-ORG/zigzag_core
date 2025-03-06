package model

import (
	"time"
)

func (a APIKey) TableName() string {
	return "api_key"
}

type APIKey struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	UserPK    int       `json:"user_pk" gorm:"not null"`
	AppKey    string    `json:"app_key" gorm:"type:text;not null"`
	AppSecret string    `json:"app_secret" gorm:"type:text;not null"`
	GeminiKey string    `json:"gemini_key" gorm:"type:text;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null;autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null;autoUpdateTime"`
}
