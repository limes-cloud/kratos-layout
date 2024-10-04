package types

type ListActivityRequest struct {
	Page     uint32  `json:"page"`
	PageSize uint32  `json:"pageSize"`
	Title    *string `json:"title"`
	IsTop    *bool   `json:"isTop"`
	Status   *bool   `json:"status"`
}
