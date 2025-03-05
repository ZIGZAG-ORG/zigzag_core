package route

import (
	"fmt"
	service "zigzag-trade/service/crud"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const prefix = "/trade_api"

func StartServer(db *gorm.DB, port int) {
	router := SetupRouter(db)
	router.Run(fmt.Sprintf(":%d", port))
}

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	router.POST(fmt.Sprintf("%s/trade_log", prefix), service.Createtrade(db))
	router.GET(fmt.Sprintf("%s/trade_log", prefix), service.FindTradeLogs(db))
	// router.PUT(fmt.Sprintf("%s/trade", prefix), service.UpdateVendor(db))
	// router.DELETE(fmt.Sprintf("%s/trade", prefix), service.DeleteVendors(db))

	return router
}
