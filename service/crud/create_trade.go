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

type createtrade struct {
	UserPK     int       `gorm:"not null" json:"user_pk"`
	Timestamp  time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"timestamp"`
	Stock      string    `gorm:"type:varchar(50);not null" json:"stock"`
	Price      float64   `gorm:"type:decimal(10,2);not null" json:"price"`
	Quantity   int       `gorm:"not null" json:"quantity"`
	TradeType  int       `gorm:"not null" json:"trade_type"` // 1: BUY, 2: SELL
	TotalPrice float64   `gorm:"type:decimal(15,2);not null" json:"total_price"`
	Status     int       `gorm:"not null;default:1" json:"status"`  // 1: PENDING, 2: COMPLETED, 3: CANCELLED
	Reason     *string   `gorm:"type:text" json:"reason,omitempty"` // 거래 실패/취소 사유 (NULL 허용)
}

func Createtrade(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input createtrade

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(error_code.ErrInvalidParameterSyntax.Code, gin.H{"error": error_code.ErrInvalidParameterSyntax.Message, "detail": err.Error()})
			return
		}

		// start transaction
		tx := db.Begin()

		if tx.Error != nil {
			c.JSON(error_code.ErrTransactionBegin.Code, gin.H{"error": error_code.ErrTransactionBegin.Message, "detail": tx.Error.Error()})
			return
		}

		var trade model.Trade

		// create query
		if err := tx.Table(trade.TableName()).Create(&input).Error; err != nil {
			c.JSON(error_code.ErrCreateRecordFailed.Code, gin.H{"error": error_code.ErrCodeCreateRecordFailed, "detail": err.Error()})
			return
		}

		// commit table
		if err := tx.Commit().Error; err != nil {
			tx.Rollback()
			c.JSON(error_code.ErrTransactionCommitFailed.Code, gin.H{"error": error_code.ErrCodeTransactionCommitFailed, "detail": err.Error()})
			return
		}

		c.JSON(success_code.Success_codeCreated.Code, gin.H{"success_code": success_code.Success_codeCreated.Message})
	}
}
