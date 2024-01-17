package types

type PageUserRequest struct {
	Page     uint32
	PageSize uint32
	Email    *string
}
