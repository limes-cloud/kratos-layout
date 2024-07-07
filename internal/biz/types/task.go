package types

type PageTaskRequest struct {
	Page      uint32  `json:"page"`
	PageSize  uint32  `json:"page_size"`
	Title     *string `json:"title"`
	NotFinish *bool   `json:"not_finish"`
	UserID    *uint32 `json:"user_id"`
}

type PageTaskValueRequest struct {
	Page     uint32 `json:"page"`
	PageSize uint32 `json:"page_size"`
	TaskID   uint32 `json:"task_id"`
}
