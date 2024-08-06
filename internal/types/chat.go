package types

type ListChatRecordRequest struct {
	Page      uint32   `json:"page"`
	PageSize  uint32   `json:"pageSize"`
	Distinct  *bool    `json:"distinct"`
	Order     *string  `json:"order"`
	OrderBy   *string  `json:"orderBy"`
	UserId    *uint32  `json:"userId"`
	UserIds   []uint32 `json:"-"`
	SessionId *string  `json:"sessionId"`
	UserName  *string  `json:"userName"`
}

type SendChatMessageRequest struct {
	SessionId string
	Message   string
	Reply     func(string)
	Event     func(string, any)
	Done      func()
}

type SendChatMessageReply struct {
	SessionId string `json:"sessionId"`
	Message   string `json:"message"`
}
