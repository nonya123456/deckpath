package game_test

import (
	"testing"

	"github.com/nonya123456/deckpath/internal/game"
	"github.com/stretchr/testify/suite"
)

type LevelTestSuite struct {
	suite.Suite
	cfg       game.LevelConfig
	underTest game.Level
}

func TestLevelTestSuite(t *testing.T) {
	suite.Run(t, new(LevelTestSuite))
}

func (t *LevelTestSuite) SetupTest() {
	t.cfg = game.LevelConfig{
		PathLength: 10,
		DeckSize:   5,
		Player:     1,
	}
	t.underTest = game.NewLevel(t.cfg)
}

func (t *LevelTestSuite) TestPath() {
	expected := make([]game.Tile, t.cfg.PathLength)
	actual := t.underTest.Path()
	t.Equal(expected, actual)
}

func (t *LevelTestSuite) TestDeck() {
	expected := make([]game.Card, t.cfg.DeckSize)
	actual := t.underTest.Deck()
	t.Equal(expected, actual)
}

func (t *LevelTestSuite) TestPlayer() {
	player := t.underTest.Player()
	t.Equal(t.cfg.Player, player)
}

func (t *LevelTestSuite) TestPlay() {
	t.Run("error", func() {
		cards := []int{11, 23, 3, 4}
		err := t.underTest.Play(cards)
		t.Error(err)
	})

	t.Run("no_error", func() {
		cards := []int{1, 2, 3, 4}
		err := t.underTest.Play(cards)
		t.NoError(err)
	})
}
