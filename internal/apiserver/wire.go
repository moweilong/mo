//go:build wireinject
// +build wireinject

package apiserver

import (
	"github.com/google/wire"

	"github.com/moweilong/mo/pkg/server"
)

func InitializeWebServer(*Config) (server.Server, error) {
	wire.Build(
		NewAggregatorServer,
		wire.Struct(new(ServerConfig), "*"), // * 表示注入全部字段
	)
	return nil, nil
}
