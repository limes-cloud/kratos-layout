package user

import (
	"github.com/limes-cloud/kratosx"
	ktypes "github.com/limes-cloud/kratosx/types"
)

type Repo interface {
	Get(ctx kratosx.Context, id uint32) (*User, error)
	Page(ctx kratosx.Context, options *ktypes.PageOptions) ([]*User, uint32, error)
	Create(ctx kratosx.Context, c *User) (uint32, error)
	Update(ctx kratosx.Context, c *User) error
	Delete(ctx kratosx.Context, uint322 uint32) error
}
