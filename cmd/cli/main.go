package main

import (
	"github.com/nonya123456/deckpath/internal/game"
	"github.com/nonya123456/deckpath/internal/wire"
)

func main() {
	container := wire.InitializeContainer()
	cfg := game.LevelConfig{
		PathLength: 10,
		DeckSize:   10,
		Player:     9,
	}
	container.Manager.Start(cfg)
}
