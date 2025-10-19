package main

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/library"
	"github.com/limes-cloud/kratosx/library/db"
	"github.com/limes-cloud/manager/api/scope"
	"partyaffairs/api/errors"
	"partyaffairs/internal/app"
	"partyaffairs/internal/core"
	"partyaffairs/internal/middleware"

	_ "go.uber.org/automaxprocs"
)

func main() {
	srv := core.InitApp(
		kratosx.WithRegistrarServer(app.Register),
		kratosx.WithValidateErrHook(func(ctx context.Context, err error) error {
			c := kratosx.MustContext(ctx)
			c.Logger().Warnw("msg", "params validate error", "err", err)
			return errors.ParamsError()
		}),
		kratosx.WithLibraryOptions(
			library.WithDBOptions(db.WithHookScope(scope.Hook)),
		),
		kratosx.WithMiddleware(middleware.Middleware()...),
	)

	if err := srv.App().Run(); err != nil {
		log.Fatal(err)
	}
}
