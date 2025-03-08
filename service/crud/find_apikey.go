package service

import (
	"net/http"
	"zigzag-trade/model"

	"github.com/LabStars/selpo-common/status/error_code"
	"github.com/LabStars/selpo-common/status/success_code"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FindAPIKey(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userPK := c.Query("user_pk")

		if userPK == "" {
			c.JSON(error_code.ErrInvalidParameterSyntax.Code, gin.H{"error": error_code.ErrInvalidParameterSyntax.Message})
			return
		}

		var apikey []model.APIKey

		result := db.Table("api_key").Where("user_pk = ?", userPK).Find(&apikey)
		if result.Error != nil {
			c.JSON(error_code.ErrFindRecordFailed.Code, gin.H{"error": error_code.ErrFindRecordFailed.Message})
			return
		}

		if len(apikey) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "No records found"})
			return
		}

		c.JSON(success_code.Success_codeFound.Code, gin.H{"success_code": success_code.Success_codeFound.Message, "data": apikey})
	}
}
