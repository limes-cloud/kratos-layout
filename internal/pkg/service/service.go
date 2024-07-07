package service

import (
	"context"

	"github.com/limes-cloud/kratosx"
	export "github.com/limes-cloud/resource/api/resource/export/v1"
	file "github.com/limes-cloud/resource/api/resource/file/v1"
	user "github.com/limes-cloud/usercenter/api/usercenter/user/v1"

	"partyaffairs/api/errors"
)

const (
	Resource   = "Resource"
	UserCenter = "UserCenter"
)

func NewResourceExport(ctx context.Context) (export.ExportClient, error) {
	conn, err := kratosx.MustContext(ctx).GrpcConn(Resource)
	if err != nil {
		return nil, errors.ResourceError()
	}
	return export.NewExportClient(conn), nil
}

func NewResourceFile(ctx context.Context) (file.FileClient, error) {
	conn, err := kratosx.MustContext(ctx).GrpcConn(Resource)
	if err != nil {
		return nil, errors.ResourceError()
	}
	return file.NewFileClient(conn), nil
}

func NewUser(ctx context.Context) (user.UserClient, error) {
	conn, err := kratosx.MustContext(ctx).GrpcConn(UserCenter)
	if err != nil {
		return nil, errors.UserCenterError()
	}
	return user.NewUserClient(conn), nil
}
