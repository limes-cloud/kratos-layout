package types

type PageNoticeRequest struct {
	Page     uint32  `json:"page"`
	PageSize uint32  `json:"page_size"`
	Title    *string `json:"title"`
	NotRead  *bool   `json:"not_read"`
	UserID   *uint32 `json:"user_id"`
}
