package notice

type PageRequest struct {
	Page     uint32
	PageSize uint32
	Email    *string
}

type PageNoticeUserRequest struct {
	Page     uint32
	PageSize uint32
	NoticeID uint32
}
