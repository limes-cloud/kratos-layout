package repository

import (
	"github.com/limes-cloud/kratosx"

	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/types"
)

type FileRepository interface {
	// GetFileURL 获取指定文件的url
	GetFileURL(ctx kratosx.Context, sha string) string

	// GetFile 获取指定文件信息
	GetFile(ctx kratosx.Context, sha string) (*entity.File, error)

	// ExportExcel 导出excel
	ExportExcel(ctx kratosx.Context, req *types.ExportExcelRequest) (uint32, error)
}
