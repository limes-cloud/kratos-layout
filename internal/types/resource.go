package types

type ListResourceRequest struct {
	Page       uint32  `json:"page"`
	PageSize   uint32  `json:"pageSize"`
	Title      *string `json:"title"`
	ClassifyId *uint32 `json:"classifyId"`
}
