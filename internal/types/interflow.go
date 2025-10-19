package types

import "github.com/limes-cloud/kratosx/model/page"

type ListInterflowRequest struct {
	*page.Search
	FromUserId uint32 `json:"fromUserId"`
	ToUserId   uint32 `json:"toUserId"`
	FromCur    bool   `json:"fromCur"`
}
