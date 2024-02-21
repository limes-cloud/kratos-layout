package service

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/go-kratos/kratos-layout/api/noticepb"
	"github.com/go-kratos/kratos-layout/api/userpb"
	"github.com/go-kratos/kratos-layout/internal/conf"
)

func New(c *conf.Config, hs *http.Server, gs *grpc.Server) {
	userSrv := NewUser(c)
	userpb.RegisterServiceHTTPServer(hs, userSrv)
	userpb.RegisterServiceServer(gs, userSrv)

	noticeSrv := NewNotice(c)
	noticepb.RegisterServiceHTTPServer(hs, noticeSrv)
	noticepb.RegisterServiceServer(gs, noticeSrv)
}
