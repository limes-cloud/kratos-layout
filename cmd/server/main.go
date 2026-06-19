package main

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/limes-cloud/kratosx"
	_ "go.uber.org/automaxprocs"

	"layout/api/errors"
	"layout/internal/app"
)

func main() {
	srv := kratosx.New(
		kratosx.WithRegistrarServer(app.Register),
		kratosx.WithValidateErrHook(func(ctx context.Context, err error) error {
			kratosx.MustContext(ctx).Logger().Warnw("msg", "params validate error", "err", err)
			return errors.ParamsError()
		}),
	)

	if err := srv.App().Run(); err != nil {
		log.Fatal(err.Error())
	}
}
