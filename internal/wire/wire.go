//go:build wireinject
// +build wireinject

package wire

import "github.com/google/wire"

func InitializeContainer() *Container {
	wire.Build(wire.Struct(new(Container), "*"))
	return &Container{}
}
