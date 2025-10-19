package repository

import (
	"partyaffairs/internal/core"
	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/types"
)

type FileRepository interface {
	// GetFileURL 获取指定文件的url
	GetFileURL(ctx core.Context, sha string) string

	// GetFile 获取指定文件信息
	GetFile(ctx core.Context, sha string) (*entity.File, error)

	// ExportExcel 导出excel
	ExportExcel(ctx core.Context, req *types.ExportExcelRequest) (uint32, error)
}
