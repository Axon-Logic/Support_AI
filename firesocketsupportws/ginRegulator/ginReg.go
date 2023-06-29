package ginReg

import (
	"firesocketSupportWs/config"
	pb "firesocketSupportWs/grpcApi"
	"firesocketSupportWs/wsManage"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GinInit(logger log.Logger, GrpcModel *wsManage.Grpc) {
	router := gin.Default()
	router.POST("firesocketsupportapi/temp/update/:userid", func(c *gin.Context) {
		userid := c.Param("userid")
		if response, err := updateHistory(userid); err == nil {
			for conn := range *GrpcModel.ChatsConnections {
				conn.Send <- response
			}
		}
	})
	router.POST("firesocketsupportapi/temp/newMessage/:userid", func(c *gin.Context) {
		var message pb.ClientMessage
		message.Sender = c.Param("userid")
		if err := c.ShouldBindJSON(&message); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		message.IsSupport = true
		if err := GrpcModel.SendMessageToProvider(message); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})
	if err := router.Run(":" + config.Cfg.GIN_SERVER_PORT); err != nil {
		logger.Fatalf("failed to serve GIN: %v", err)
	}
}

func updateHistory(userid string) ([]byte, error) {
	body := []byte{}
	err := error(nil)
	url := config.Cfg.HISTORY_SERVICE_URL + "/history/update/" + userid
	if resp, err := http.Post(url, "application/json", nil); err == nil {
		body, err = ioutil.ReadAll(resp.Body)
	}
	return body, err
}
