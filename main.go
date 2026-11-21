package main

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/limes-cloud/kratosx"
	_ "go.uber.org/automaxprocs"

	"tianji/api/errors"
	"tianji/internal/app"
	"tianji/internal/core"
)

func main() {
	srv := core.InitApp(
		kratosx.WithRegistrarServer(app.Register),
		kratosx.WithValidateErrHook(func(ctx context.Context, err error) error {
			core.MustContext(ctx).Logger().Warnw("msg", "params validate error", "err", err)
			return errors.ParamsError()
		}),
	)

	if err := srv.App().Run(); err != nil {
		log.Fatal(err.Error())
	}
}
