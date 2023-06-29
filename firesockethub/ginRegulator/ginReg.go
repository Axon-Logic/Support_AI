package ginReg

import (
	"firesocketHub/config"
	"firesocketHub/grpcMessage"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GinInit(logger log.Logger, StreamServer *grpcMessage.Server) {
	router := gin.Default()
	router.GET("/getProvider", func(c *gin.Context) {
		provider := StreamServer.GetProvider()
		c.JSON(http.StatusOK, provider)
	})

	if err := router.Run(":" + config.Cfg.GIN_SERVER_PORT); err != nil {
		StreamServer.Logger.Fatalf("failed to serve GIN: %v", err)
	}
}
