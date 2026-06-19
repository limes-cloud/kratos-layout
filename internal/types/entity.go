package types

import "github.com/limes-cloud/kratosx/model/page"

type ListEntityRequest struct {
	page.Search
	AppId    *uint32 `json:"appId"`
	Name     *string `json:"name"`
	Database *string `json:"database"`
}

type ListEntityFieldRequest struct {
	EntityId uint32  `json:"entityId"`
	Name     *string `json:"name"`
}

type ListEntityRuleRequest struct {
	*page.Search
	EntityId *uint32 `json:"entityId"`
	Name     *string `json:"name"`
	Status   *bool   `json:"status"`
}

type RegisterEntityRequest struct {
	AppKeyword string                `json:"appKeyword"`
	Mode       string                `json:"mode"`
	Entities   []*RegisterEntityNode `json:"entities"`
}

type RegisterEntityNode struct {
	Database string                     `json:"database"`
	Name     string                     `json:"name"`
	Comment  string                     `json:"comment"`
	Fields   []*RegisterEntityFieldNode `json:"fields"`
	Rules    []*RegisterEntityRuleNode  `json:"rules"`
}

type RegisterEntityFieldNode struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`
	Index   int    `json:"index"`
}

type RegisterEntityRuleNode struct {
	Name        string `json:"name"`
	Status      *bool  `json:"status"`
	Expression  string `json:"expression"`
	Description string `json:"description"`
}

type RegisterEntityReply struct {
	Created uint32 `json:"created"`
	Updated uint32 `json:"updated"`
	Deleted uint32 `json:"deleted"`
}
