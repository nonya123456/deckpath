package game

type Tile struct{}

type Level interface {
	Path() []Tile
}

type level struct {
	path []Tile
}

func (l *level) Path() []Tile {
	return l.path
}

var _ Level = (*level)(nil)

func NewLevel(pathLength int) Level {
	path := make([]Tile, pathLength)
	return &level{path: path}
}
