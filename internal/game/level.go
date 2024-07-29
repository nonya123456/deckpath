package game

import (
	"errors"
)

type LevelConfig struct {
	PathLength int
	DeckSize   int
	Player     int
}

type Level interface {
	Path() []Tile
	Deck() []Card
	Player() int
	Play(cards []int) error
}

type level struct {
	path   []Tile
	deck   []Card
	player int
}

func (l *level) Path() []Tile {
	return l.path
}

func (l *level) Deck() []Card {
	return l.deck
}

func (l *level) Player() int {
	return l.player
}

func (l *level) Play(cards []int) error {
	err := l.validateCards(cards)
	if err != nil {
		return err
	}

	return nil
}

func (l *level) validateCards(cards []int) error {
	for _, card := range cards {
		if card < 0 || card >= len(l.deck) {
			return errors.New("card index out of range")
		}
	}

	return nil
}

var _ Level = (*level)(nil)

func NewLevel(cfg LevelConfig) Level {
	path := NewPath(cfg.PathLength)
	deck := NewDeck(cfg.DeckSize)
	return &level{
		path:   path,
		deck:   deck,
		player: cfg.Player,
	}
}
