package types

type PageVideoContentRequest struct {
	Page       uint32  `json:"page"`
	PageSize   uint32  `json:"page_size"`
	Title      *string `json:"title"`
	ClassifyID *uint32 `json:"classify_id"`
	UserID     *uint32 `json:"user_id"`
}

type PageVideoClassifyRequest struct {
	Page     uint32  `json:"page"`
	PageSize uint32  `json:"page_size"`
	Name     *string `json:"name"`
}
