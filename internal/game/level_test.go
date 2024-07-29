package game_test

import (
	"testing"

	"github.com/nonya123456/deckpath/internal/game"
	"github.com/stretchr/testify/suite"
)

type LevelTestSuite struct {
	suite.Suite
	pathLength int
	underTest  game.Level
}

func TestLevelTestSuite(t *testing.T) {
	suite.Run(t, new(LevelTestSuite))
}

func (t *LevelTestSuite) SetupTest() {
	t.pathLength = 10
	t.underTest = game.NewLevel(t.pathLength)
}

func (t *LevelTestSuite) TestPath() {
	expected := make([]game.Tile, t.pathLength)
	actual := t.underTest.Path()
	t.Equal(expected, actual)
}
