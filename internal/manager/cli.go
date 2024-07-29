package manager

import (
	"errors"
	"fmt"
	"strconv"
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
		cmd, args, err := m.promptReader.ReadNext()
		if err != nil {
			m.fallbackWithError(err)
			continue
		}

		switch cmd {
		case prompt.CommandHelp:
			m.showHelp()
		case prompt.CommandPath:
			m.showPath()
		case prompt.CommandDeck:
			m.showDeck()
		case prompt.CommandPlay:
			err = m.play(args)
			if err != nil {
				m.fallbackWithError(err)
				continue
			}
			m.showPath()
		case prompt.CommandQuit:
			return
		default:
			m.fallbackWithError(errors.New("command not found"))
		}
	}
}

func (m *cliManager) showHelp() {
	helpText := `Available commands:
- help                                      : Show this help message
- path                                      : Display the current path
- deck                                      : Display the current deck
- play <card_index> [card_index ...]        : Play cards from the current deck
- quit                                      : Quit game`
	fmt.Println(helpText)
}

func (m *cliManager) showPath() {
	turn := m.level.Turn()
	score := m.level.Score()
	fmt.Printf("Turn: %d, Score: %d\n", turn, score)

	path := m.level.Path()
	tiles := make([]string, 0, len(path))
	for _, tile := range path {
		tiles = append(tiles, tile.ToString())
	}
	fmt.Println(strings.Join(tiles, " "))

	player := m.level.Player()
	fmt.Println(strings.Repeat("  ", player) + "^ " + strings.Repeat("  ", len(path)-player-1))
}

func (m *cliManager) showDeck() {
	deck := m.level.Deck()
	cards := make([]string, 0, len(deck))
	for i, card := range deck {
		cards = append(cards, fmt.Sprintf("%d) %s", i, card.ToString()))
	}

	fmt.Println(strings.Join(cards, "\n"))
}

func (m *cliManager) play(args []string) error {
	cards := make([]int, 0, len(args))
	for _, arg := range args {
		card, err := strconv.Atoi(arg)
		if err != nil {
			return err
		}

		cards = append(cards, card)
	}

	return m.level.Play(cards)
}

func (m *cliManager) fallbackWithError(e error) {
	fmt.Println("SKIPPED: " + e.Error())
}

var _ CLIManager = (*cliManager)(nil)

func ProvideCLIManager(r prompt.Reader) CLIManager {
	return &cliManager{promptReader: r}
}
