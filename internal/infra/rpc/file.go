package rpc

import (
	"github.com/limes-cloud/kratosx"
	export "github.com/limes-cloud/resource/api/resource/export/v1"
	file "github.com/limes-cloud/resource/api/resource/file/v1"

	"partyaffairs/api/partyaffairs/errors"
	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/types"
)

const (
	Resource = "Resource"
)

type File struct {
}

func NewFile() *File {
	return &File{}
}

func (i File) client(ctx kratosx.Context) (file.FileClient, error) {
	conn, err := kratosx.MustContext(ctx).GrpcConn(Resource)
	if err != nil {
		return nil, errors.ResourceServerError()
	}
	return file.NewFileClient(conn), nil
}

func (i File) exportClient(ctx kratosx.Context) (export.ExportClient, error) {
	conn, err := kratosx.MustContext(ctx).GrpcConn(Resource)
	if err != nil {
		return nil, errors.ResourceServerError()
	}
	return export.NewExportClient(conn), nil
}

func (i File) GetFileURL(ctx kratosx.Context, sha string) string {
	client, err := i.client(ctx)
	if err != nil {
		ctx.Logger().Warnw("msg", "connect resource server error", "err", err.Error())
		return ""
	}
	reply, err := client.GetFile(ctx, &file.GetFileRequest{Sha: &sha})
	if err != nil {
		ctx.Logger().Warnw("msg", "get resource file error", "err", err.Error())
		return ""
	}
	return reply.Url
}

func (i File) GetFile(ctx kratosx.Context, sha string) (*entity.File, error) {
	client, err := i.client(ctx)
	if err != nil {
		return nil, err
	}
	reply, err := client.GetFile(ctx, &file.GetFileRequest{Sha: &sha})
	if err != nil {
		return nil, err
	}
	return &entity.File{
		Name: reply.Name,
		Type: reply.Type,
		Size: reply.Size,
		URL:  reply.Url,
	}, nil
}

func (i File) ExportExcel(ctx kratosx.Context, req *types.ExportExcelRequest) (uint32, error) {
	ec, err := i.exportClient(ctx)
	if err != nil {
		return 0, err
	}
	in := &export.ExportExcelRequest{
		UserId:       req.UserId,
		DepartmentId: req.DepartmentId,
		Scene:        req.Scene,
		Name:         req.Name,
		Headers:      req.Headers,
	}
	for _, item := range req.Rows {
		var cols []*export.ExportExcelRequest_Col
		for _, col := range item {
			cols = append(cols, &export.ExportExcelRequest_Col{
				Type:  col.Type,
				Value: col.Value,
			})
		}
		in.Rows = append(in.Rows, &export.ExportExcelRequest_Row{
			Cols: cols,
		})
	}

	for _, item := range req.Files {
		in.Files = append(in.Files, &export.ExportExcelRequest_ExportFile{
			Value:  item.Value,
			Rename: item.Rename,
		})
	}

	reply, err := ec.ExportExcel(ctx, in)
	if err != nil {
		return 0, err
	}
	return reply.Id, nil
}
