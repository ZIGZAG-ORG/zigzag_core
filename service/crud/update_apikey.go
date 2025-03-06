package service

import (
	"net/http"
	"time"
	"zigzag-trade/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type updateAPIKeyinput struct {
	UserPK    int    `json:"user_pk" gorm:"primaryKey;not null"`
	AppKey    string `json:"app_key" gorm:"not null"`
	AppSecret string `json:"app_secret" gorm:"not null"`
	GeminiKey string `json:"gemini_key" gorm:"not null"`
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
			"app_key":    input.AppKey,
			"app_secret": input.AppSecret,
			"gemini_key": input.GeminiKey,
			"updatedAt":  time.Now().UTC(),
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
