package types

type ListUserRequest struct {
	Page     uint32   `json:"page"`
	PageSize uint32   `json:"pageSize"`
	In       []uint32 `json:"in"`
	NotIn    []uint32 `json:"notIn"`
}
