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
	OpenAIKey string    `json:"open_ai_key" gorm:"type:text;not null"`
	GeminiKey string    `json:"gemini_key" gorm:"type:text;not null"`
	Account01 string    `json:"account_01" gorm:"column:account_01;not null"`
	Account02 string    `json:"account_02" gorm:"column:account_02;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null;autoCreateTime;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null;autoUpdateTime;default:CURRENT_TIMESTAMP"`
}
