package types

type ListCommentRequest struct {
	Page     uint32   `json:"page"`
	PageSize uint32   `json:"pageSize"`
	Order    *string  `json:"order"`
	OrderBy  *string  `json:"orderBy"`
	UserName *string  `json:"userName"`
	UserIds  []uint32 `json:"-"`
}
