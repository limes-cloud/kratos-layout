package user

type PageRequest struct {
	Page     uint32
	PageSize uint32
	Email    *string
}
