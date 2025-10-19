package types

import "github.com/limes-cloud/kratosx/model/page"

type ListNoticeRequest struct {
	*page.Search
	Title   *string `json:"title"`
	IsTop   *bool   `json:"isTop"`
	Status  *bool   `json:"status"`
	NotRead *bool   `json:"notRead"`
}
