package game

type LevelConfig struct {
	PathLength int
	DeckSize   int
	Player     int
}

type Level interface {
	Path() []Tile
	Deck() []Card
	Player() int
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
