package main

import (
	"fmt"

	"github.com/nonya123456/deckpath/internal/wire"
)

func main() {
	container := wire.InitializeContainer()
	cmd, err := container.PromptReader.ReadNext()
	if err != nil {
		panic(err)
	}

	fmt.Println(cmd)
}
