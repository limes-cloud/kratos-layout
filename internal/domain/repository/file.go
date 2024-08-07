package repository

import (
	"github.com/limes-cloud/kratosx"
	"poverty/internal/domain/entity"
)

type FileRepository interface {
	// GetFileURL 获取指定文件的url
	GetFileURL(ctx kratosx.Context, sha string) string
	GetFile(ctx kratosx.Context, sha string) (*entity.File, error)
}
