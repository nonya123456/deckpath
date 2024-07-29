package manager

import (
	"fmt"
	"strings"

	"github.com/nonya123456/deckpath/internal/game"
	"github.com/nonya123456/deckpath/internal/prompt"
)

type CLIManager interface {
	Start(cfg game.LevelConfig)
}

type cliManager struct {
	promptReader prompt.Reader
	level        game.Level
}

func (m *cliManager) Start(cfg game.LevelConfig) {
	m.level = game.NewLevel(cfg)

	for {
		cmd, err := m.promptReader.ReadNext()
		if err != nil {
			m.fallback()
			continue
		}

		switch cmd {
		case prompt.CommandHelp:
			m.handleHelp()
		case prompt.CommandPath:
			m.handlePath()
		case prompt.CommandDeck:
			m.handleDeck()
		case prompt.CommandQuit:
			return
		default:
			m.fallback()
		}
	}
}

func (m *cliManager) handleHelp() {
	helpText := `	Available commands:
	- help        : Show this help message
	- path        : Display the current path
	- deck        : Display the current deck
	- quit        : Quit game`
	fmt.Println(helpText)
}

func (m *cliManager) handlePath() {
	path := m.level.Path()
	tiles := make([]string, 0, len(path))
	for range path {
		tiles = append(tiles, "O")
	}

	fmt.Println(strings.Join(tiles, " "))
}

func (m *cliManager) handleDeck() {
	path := m.level.Deck()
	cards := make([]string, 0, len(path))
	for range path {
		cards = append(cards, "C")
	}

	fmt.Println(strings.Join(cards, " "))
}

func (m *cliManager) fallback() {
	fmt.Println("Error")
}

var _ CLIManager = (*cliManager)(nil)

func ProvideCLIManager(r prompt.Reader) CLIManager {
	return &cliManager{promptReader: r}
}
