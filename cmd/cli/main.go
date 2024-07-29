package main

import (
	"fmt"

	"github.com/nonya123456/deckpath/internal/wire"
)

func main() {
	_ = wire.InitializeContainer()
	fmt.Println("Hello from cli")
}
