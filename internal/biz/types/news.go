package types

type PageNewsContentRequest struct {
	Page       uint32  `json:"page"`
	PageSize   uint32  `json:"page_size"`
	Title      *string `json:"title"`
	ClassifyID *uint32 `json:"classify_id"`
}

type PageNewsCommentRequest struct {
	Page      uint32  `json:"page"`
	PageSize  uint32  `json:"page_size"`
	ContentID uint32  `json:"content_id"`
	Text      *string `json:"text"`
}
