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
	Turn() int
	Score() int
}

type level struct {
	path   []Tile
	deck   []Card
	player int
	score  int
	turn   int
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

func (l *level) Turn() int {
	return l.turn
}

func (l *level) Score() int {
	return l.score
}

func (l *level) Play(cards []int) error {
	l.nextTurn()

	err := l.validateCards(cards)
	if err != nil {
		return err
	}

	for _, idx := range cards {
		l.playCard(l.deck[idx])
		l.deck[idx] = NewCard()
	}

	return nil
}

func (l *level) playCard(card Card) {
	switch card.Effect {
	case CardEffectMove:
		l.player += card.Amount
		l.player = max(0, min(len(l.path)-1, l.player))
	case CardEffectPlant:
		l.path[l.player].Growth = max(l.path[l.player].Growth, 1)
	case CardEffectHarvest:
		l.score += l.path[l.player].Growth
		l.path[l.player].Growth = 0
	}
}

func (l *level) nextTurn() {
	l.turn += 1

	for i := range l.path {
		if l.path[i].Growth == 0 {
			continue
		}

		if l.path[i].Growth >= TileMaxGrowth {
			l.path[i].Growth = 0
			continue
		}

		l.path[i].Growth += 1
	}
}

func (l *level) validateCards(cards []int) error {
	if len(cards) == 0 {
		return errors.New("no card provided")
	}

	cardSet := make(map[int]struct{})
	for _, card := range cards {
		if card < 0 || card >= len(l.deck) {
			return errors.New("card index out of range")
		}

		_, ok := cardSet[card]
		if ok {
			return errors.New("duplicate card found")
		}

		cardSet[card] = struct{}{}
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
