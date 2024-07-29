package game

type LevelConfig struct {
	PathLength int
	DeckSize   int
}

type Level interface {
	Path() []Tile
	Deck() []Card
}

type level struct {
	path []Tile
	deck []Card
}

func (l *level) Path() []Tile {
	return l.path
}

func (l *level) Deck() []Card {
	return l.deck
}

var _ Level = (*level)(nil)

func NewLevel(cfg LevelConfig) Level {
	path := NewPath(cfg.PathLength)
	deck := NewDeck(cfg.DeckSize)
	return &level{path: path, deck: deck}
}
