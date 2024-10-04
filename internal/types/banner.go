package types

type ListBannerRequest struct {
	Page     uint32  `json:"page"`
	PageSize uint32  `json:"pageSize"`
	Title    *string `json:"title"`
	Status   *bool   `json:"status"`
}
