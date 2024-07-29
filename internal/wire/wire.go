//go:build wireinject
// +build wireinject

package wire

import "github.com/google/wire"

func InitializeContainer() *Container {
	wire.Build(ProviderSet, wire.Struct(new(Container), "*"))
	return &Container{}
}
