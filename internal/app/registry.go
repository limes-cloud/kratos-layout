package app

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type registryFunc func(hs *http.Server)

var registries []registryFunc

func register(fn registryFunc) {
	registries = append(registries, fn)
}

func Register(hs *http.Server, _ *grpc.Server) {
	for _, registry := range registries {
		registry(hs)
	}
}
