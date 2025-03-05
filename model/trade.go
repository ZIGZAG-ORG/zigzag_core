package model

import (
	"time"
)

func (t Trade) TableName() string {
	return "trade_log"
}

type Trade struct {
	ID         int       `gorm:"primaryKey;autoIncrement" json:"id"`
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
