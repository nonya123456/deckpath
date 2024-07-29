package game

type Tile struct{}

func NewPath(length int) []Tile {
	path := make([]Tile, length)
	return path
}
