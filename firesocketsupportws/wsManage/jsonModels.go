package wsManage

type ChatJson struct {
	Sender         string  `json:"client_id"`
	Message        string  `json:"message"`
	MessageDate    float32 `json:"date"`
	MessageCount   int32   `json:"message_count"`
	Count          bool    `json:"count"`
	UserName       string  `json:"user_name"`
	FirstName      string  `json:"first_name"`
	LastName       string  `json:"last_name"`
	Owner          bool    `json:"owner"`
	MessageType    string  `json:"type"`
	MessageId      string  `json:"message_id"`
	Caption        string  `json:"caption"`
	ReplyMessageId string  `json:"reply_message_id"`
}

type GetMessageByIdJson struct {
	Sender string `json:"client_id"`
}
