package route

import (
	"fmt"
	service "zigzag-core/service/crud"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const prefix = "/core_api"

func StartServer(db *gorm.DB, port int) {
	router := SetupRouter(db)
	router.Run(fmt.Sprintf(":%d", port))
}

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	router.POST(fmt.Sprintf("%s/api_key", prefix), service.CreateAPIKey(db))
	router.GET(fmt.Sprintf("%s/api_key", prefix), service.FindAPIKey(db))
	router.PUT(fmt.Sprintf("%s/api_key", prefix), service.UpdateAPIKey(db))
	router.DELETE(fmt.Sprintf("%s/api_key", prefix), service.DeleteAPIKey(db))

	return router
}
