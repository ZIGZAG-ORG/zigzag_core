package service

import (
	"time"
	"zigzag-trade/model"

	"github.com/LabStars/selpo-common/status/error_code"
	"github.com/LabStars/selpo-common/status/success_code"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Token 검증 API 엔드포인트

type inputAPIKey struct {
	UserPK    int    `json:"user_pk" gorm:"not null"`
	AppKey    string `json:"app_key" gorm:"type:text;not null"`
	AppSecret string `json:"app_secret" gorm:"type:text;not null"`
	OpenAIKey string `json:"open_ai_key" gorm:"type:text;not null"`
	GeminiKey string `json:"gemini_key" gorm:"type:text;not null"`
	Account01 string `json:"account_01" gorm:"column:account_01;not null"`
	Account02 string `json:"account_02" gorm:"column:account_02;not null"`
}

func CreateAPIKey(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input inputAPIKey

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(error_code.ErrInvalidParameterSyntax.Code, gin.H{
				"error":  error_code.ErrInvalidParameterSyntax.Message,
				"detail": err.Error(),
			})
			return
		}

		tx := db.Begin()
		if tx.Error != nil {
			c.JSON(error_code.ErrTransactionBegin.Code, gin.H{
				"error":  error_code.ErrTransactionBegin.Message,
				"detail": tx.Error.Error(),
			})
			return
		}

		apiKey := model.APIKey{
			UserPK:    input.UserPK,
			AppKey:    input.AppKey,
			AppSecret: input.AppSecret,
			OpenAIKey: input.OpenAIKey,
			GeminiKey: input.GeminiKey,
			Account01: input.Account01,
			Account02: input.Account02,
			CreatedAt: time.Now(),
		}

		if err := tx.Table(apiKey.TableName()).Create(&apiKey).Error; err != nil {
			tx.Rollback() // 롤백 수행
			c.JSON(error_code.ErrCreateRecordFailed.Code, gin.H{
				"error":  error_code.ErrCreateRecordFailed.Message,
				"detail": err.Error(),
			})
			return
		}

		if err := tx.Commit().Error; err != nil {
			tx.Rollback()
			c.JSON(error_code.ErrTransactionCommitFailed.Code, gin.H{
				"error":  error_code.ErrTransactionCommitFailed.Message,
				"detail": err.Error(),
			})
			return
		}

		c.JSON(success_code.Success_codeCreated.Code, gin.H{
			"success_code": success_code.Success_codeCreated.Message,
		})
	}
}
