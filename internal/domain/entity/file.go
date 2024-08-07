package entity

type File struct {
	Name string `gorm:"-" json:"name"`
	Type string `gorm:"-" json:"type"`
	Size uint32 `gorm:"-" json:"size"`
	URL  string `gorm:"-" json:"url"`
}
