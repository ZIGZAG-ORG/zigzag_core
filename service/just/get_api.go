package service

import (
	"errors"
	"zigzag-core/model"

	"gorm.io/gorm"
)

func GetAPIKeyByUserPK(db *gorm.DB, userPK string) (*model.APIKey, error) {
	var apiKey model.APIKey

	// Gorm을 사용하여 user_pk로 API Key 조회
	result := db.Table("api_key").Where("user_pk = ?", userPK).First(&apiKey)
	if result.Error != nil {
		return nil, errors.New("API 키 조회 실패: " + result.Error.Error())
	}

	return &apiKey, nil
}
