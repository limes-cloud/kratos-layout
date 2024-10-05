package types

type ListTaskRequest struct {
	Page      uint32  `json:"page"`
	PageSize  uint32  `json:"pageSize"`
	Title     *string `json:"title"`
	UserId    *uint32 `json:"userId"`
	NotFinish *bool   `json:"notFinish"`
}

type ListTaskValueRequest struct {
	Page     uint32 `json:"page"`
	PageSize uint32 `json:"pageSize"`
	TaskId   uint32 `json:"taskId"`
	Finish   *bool  `json:"finish"`
}
