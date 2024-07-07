package types

type PageResourceContentRequest struct {
	Page       uint32  `json:"page"`
	PageSize   uint32  `json:"page_size"`
	Title      *string `json:"title"`
	ClassifyID *uint32 `json:"classify_id"`
}
