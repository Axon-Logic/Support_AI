package wsManage

type WSRecievedData struct {
	Sender    string `json:"Sender" binding:"required"`
	Subscribe int    `json:"Subscribe" binding:"required"`
}
