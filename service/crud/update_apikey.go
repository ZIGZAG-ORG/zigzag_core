package service

import (
	"net/http"
	"time"
	"zigzag-trade/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type updateAPIKeyinput struct {
	UserPK    int    `json:"user_pk" gorm:"not null"`
	AppKey    string `json:"app_key" gorm:"type:text;not null"`
	AppSecret string `json:"app_secret" gorm:"type:text;not null"`
	OpenAIKey string `json:"open_ai_key" gorm:"type:text;not null"`
	GeminiKey string `json:"gemini_key" gorm:"type:text;not null"`
	Account01 string `json:"account_01" gorm:"column:account_01;not null"`
	Account02 string `json:"account_02" gorm:"column:account_02;not null"`
}

func UpdateAPIKey(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input updateAPIKeyinput
		var apikey model.APIKey

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Where("user_pk = ?", input.UserPK).First(&apikey).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "API Key not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		updateApikey := map[string]interface{}{
			"app_key":     input.AppKey,
			"app_secret":  input.AppSecret,
			"open_ai_key": input.OpenAIKey,
			"gemini_key":  input.GeminiKey,
			"account_01":  input.Account01,
			"account_02":  input.Account02,
			"updated_at":  time.Now().UTC(),
		}

		tx := db.Begin()

		if tx.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": tx.Error.Error()})
			return
		}

		if err := tx.Model(&apikey).Where("user_pk = ?", input.UserPK).Updates(updateApikey).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := tx.Commit().Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"result": "API Key updated successfully"})
	}
}
