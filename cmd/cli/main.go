package main

import (
	"github.com/nonya123456/deckpath/internal/wire"
)

func main() {
	container := wire.InitializeContainer()
	container.Manager.Start(10)
}
