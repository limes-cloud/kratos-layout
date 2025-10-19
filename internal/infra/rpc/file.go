package rpc

import (
	"github.com/limes-cloud/resource/api/export"
	"github.com/limes-cloud/resource/api/file"
	"partyaffairs/internal/core"

	"partyaffairs/api/errors"
	"partyaffairs/internal/domain/entity"
	"partyaffairs/internal/types"
)

const (
	Resource = "Resource"
)

type File struct {
}

func (i File) GetFileURL(ctx core.Context, sha string) string {
	//TODO implement me
	panic("implement me")
}

func NewFile() *File {
	return &File{}
}

func (i File) client(ctx core.Context) (file.FileClient, error) {
	conn, err := core.MustContext(ctx).GrpcConn(Resource)
	if err != nil {
		return nil, errors.ResourceServerError()
	}
	return file.NewFileClient(conn), nil
}

func (i File) exportClient(ctx core.Context) (export.ExportClient, error) {
	conn, err := core.MustContext(ctx).GrpcConn(Resource)
	if err != nil {
		return nil, errors.ResourceServerError()
	}
	return export.NewExportClient(conn), nil
}

func (i File) GetFile(ctx core.Context, sha string) (*entity.File, error) {
	return nil, nil
	//client, err := i.client(ctx)
	//if err != nil {
	//	return nil, err
	//}
	//return nil, err
	//reply, err := client.GetFile(ctx, &file.GetFileRequest{Sha: &sha})
	//if err != nil {
	//	return nil, err
	//}
	//return &entity.File{
	//	Name: reply.Name,
	//	Type: reply.Type,
	//	Size: reply.Size,
	//	URL:  reply.Url,
	//}, nil
}

func (i File) ExportExcel(ctx core.Context, req *types.ExportExcelRequest) (uint32, error) {
	return 0, nil
	//
	//ec, err := i.exportClient(ctx)
	//if err != nil {
	//	return 0, err
	//}
	//in := &export.ExportExcelRequest{
	//	Scene: req.Scene,
	//	Name:  req.Name,
	//}
	//for _, item := range req.Rows {
	//	var cols []*export.ExportExcelRequest_Col
	//	for _, col := range item {
	//		cols = append(cols, &export.ExportExcelRequest_Col{
	//			Type:  col.Type,
	//			Value: col.Value,
	//		})
	//	}
	//	in.Rows = append(in.Rows, &export.ExportExcelRequest_Row{
	//		Cols: cols,
	//	})
	//}
	//
	//for _, item := range req.Files {
	//	in.Files = append(in.Files, &export.ExportExcelRequest_ExportFile{
	//		Value:  item.Value,
	//		Rename: item.Rename,
	//	})
	//}
	//
	//reply, err := ec.ExportExcel(ctx, in)
	//if err != nil {
	//	return 0, err
	//}
	//return reply.Id, nil
}
