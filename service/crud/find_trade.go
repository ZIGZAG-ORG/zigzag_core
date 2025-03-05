package service

import (
	"net/http"
	"zigzag-trade/model"

	"github.com/LabStars/selpo-common/status/error_code"
	"github.com/LabStars/selpo-common/status/success_code"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FindTradeLogs(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(error_code.ErrUnauthorizedAccess.Code, gin.H{"error": error_code.ErrUnauthorizedAccess.Message, "detail": "missing access token"})
			return
		}
		userPK := c.Query("user_pk")

		// 현재 유저가 userPK 인지 확인 (auth server 에서 받은 claim으로 확인)

		if userPK == "" {
			c.JSON(error_code.ErrInvalidParameterSyntax.Code, gin.H{"error": error_code.ErrInvalidParameterSyntax.Message})
			return
		}

		var trades []model.Trade

		result := db.Table("trade_log").Where("user_pk = ?", userPK).Find(&trades)
		if result.Error != nil {
			c.JSON(error_code.ErrFindRecordFailed.Code, gin.H{"error": error_code.ErrFindRecordFailed.Message})
			return
		}

		if len(trades) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "No records found"})
			return
		}

		c.JSON(success_code.Success_codeFound.Code, gin.H{"success_code": success_code.Success_codeFound.Message, "data": trades})
	}
}
