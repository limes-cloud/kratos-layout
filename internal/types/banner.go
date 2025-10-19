package types

import "github.com/limes-cloud/kratosx/model/page"

type ListBannerRequest struct {
	*page.Search
	Title  *string `json:"title"`
	Status *bool   `json:"status"`
}

type ListClientBannerRequest struct {
	Status *bool `json:"status"`
}
