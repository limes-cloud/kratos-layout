package types

type ListNoticeRequest struct {
	Page     uint32  `json:"page"`
	PageSize uint32  `json:"pageSize"`
	Order    *string `json:"order"`
	OrderBy  *string `json:"orderBy"`
	Title    *string `json:"title"`
	IsTop    *bool   `json:"isTop"`
	Status   *bool   `json:"status"`
}
