package service

import (
	"zigzag-trade/model"

	"github.com/LabStars/selpo-common/status/error_code"
	"github.com/LabStars/selpo-common/status/success_code"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteTrade(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(error_code.ErrUnauthorizedAccess.Code, gin.H{"error": error_code.ErrUnauthorizedAccess.Message, "detail": "missing access token"})
			return
		}

		userPK := c.Query("user_pk")

		tx := db.Begin()
		if tx.Error != nil {
			c.JSON(error_code.ErrTransactionBegin.Code, gin.H{"error": error_code.ErrTransactionBegin.Message, "detail": tx.Error.Error()})
			return
		}

		if err := db.Delete(&model.Trade{}, userPK).Error; err != nil {
			tx.Rollback()
			c.JSON(error_code.ErrDeleteRecordFailed.Code, gin.H{"error": error_code.ErrDeleteRecordFailed.Message, "detail": err.Error()})
			return
		}

		// 트랜잭션 커밋
		if err := tx.Commit().Error; err != nil {
			c.JSON(error_code.ErrTransactionCommitFailed.Code, gin.H{"error": error_code.ErrTransactionCommitFailed.Message, "detail": err.Error()})
			return
		}

		c.JSON(success_code.Success_codeDeleted.Code, gin.H{"success_code": success_code.Success_codeDeleted.Message, "data": nil})
	}
}
