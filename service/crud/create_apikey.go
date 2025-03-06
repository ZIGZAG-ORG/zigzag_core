package service

import (
	"zigzag-trade/model"

	"github.com/LabStars/selpo-common/status/error_code"
	"github.com/LabStars/selpo-common/status/success_code"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Token 검증 API 엔드포인트

type inputAPIKey struct {
	UserPK    int    `json:"user_pk" gorm:"primaryKey;not null"`
	AppKey    string `json:"app_key" gorm:"not null"`
	AppSecret string `json:"app_secret" gorm:"not null"`
	GeminiKey string `json:"gemini_key" gorm:"not null"`
}

func CreateAPIKey(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input inputAPIKey

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(error_code.ErrInvalidParameterSyntax.Code, gin.H{"error": error_code.ErrInvalidParameterSyntax.Message, "detail": err.Error()})
			return
		}

		tx := db.Begin()

		if tx.Error != nil {
			c.JSON(error_code.ErrTransactionBegin.Code, gin.H{"error": error_code.ErrTransactionBegin.Message, "detail": tx.Error.Error()})
			return
		}

		var apiKey model.APIKey

		if err := tx.Table(apiKey.TableName()).Create(&input).Error; err != nil {
			c.JSON(error_code.ErrCreateRecordFailed.Code, gin.H{"error": error_code.ErrCodeCreateRecordFailed, "detail": err.Error()})
			return
		}

		if err := tx.Commit().Error; err != nil {
			tx.Rollback()
			c.JSON(error_code.ErrTransactionCommitFailed.Code, gin.H{"error": error_code.ErrCodeTransactionCommitFailed, "detail": err.Error()})
			return
		}

		c.JSON(success_code.Success_codeCreated.Code, gin.H{"success_code": success_code.Success_codeCreated.Message})
	}
}
