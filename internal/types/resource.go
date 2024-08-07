package types

type ListResourceRequest struct {
	Page     uint32  `json:"page"`
	PageSize uint32  `json:"pageSize"`
	Order    *string `json:"order"`
	OrderBy  *string `json:"orderBy"`
	Title    *string `json:"title"`
	Status   *bool   `json:"status"`
}
