package types

type ListInformationRequest struct {
	Page       uint32  `json:"page"`
	PageSize   uint32  `json:"pageSize"`
	Title      *string `json:"title"`
	IsTop      *bool   `json:"isTop"`
	Status     *bool   `json:"status"`
	ClassifyId *uint32 `json:"classifyId"`
}
