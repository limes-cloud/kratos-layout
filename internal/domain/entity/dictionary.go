package entity

type Dictionary struct {
	Id          uint32  `json:"id" gorm:"id"`
	Keyword     string  `json:"keyword" gorm:"keyword"`
	Name        string  `json:"name" gorm:"name"`
	Description *string `json:"description" gorm:"description"`
	CreatedAt   int64   `json:"createdAt" gorm:"created_at"`
	UpdatedAt   int64   `json:"updatedAt" gorm:"updated_at"`
}
