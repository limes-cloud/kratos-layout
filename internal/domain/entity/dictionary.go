package entity

type Dictionary struct {
	Id          uint32  `json:"id"`
	Keyword     string  `json:"keyword"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	CreatedAt   int64   `json:"createdAt"`
	UpdatedAt   int64   `json:"updatedAt"`
}
