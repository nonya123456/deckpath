package game

import "strconv"

const (
	TileMaxGrowth = 4
)

type Tile struct {
	Growth int
}

func (t Tile) ToString() string {
	return strconv.Itoa(t.Growth)
}

func NewPath(length int) []Tile {
	path := make([]Tile, length)
	return path
}
