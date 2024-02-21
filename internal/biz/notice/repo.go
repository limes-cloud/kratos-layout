package notice

import (
	"github.com/limes-cloud/kratosx"
)

type Repo interface {
	Get(ctx kratosx.Context, id uint32) (*Notice, error)
	Page(ctx kratosx.Context, req *PageRequest) ([]*Notice, uint32, error)
	Create(ctx kratosx.Context, c *Notice) (uint32, error)
	Update(ctx kratosx.Context, c *Notice) error
	Delete(ctx kratosx.Context, uint322 uint32) error
}
