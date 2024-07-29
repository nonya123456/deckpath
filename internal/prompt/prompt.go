package prompt

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Command string

const (
	CommandUnspecified Command = "UNSPECIFIED"
	CommandHelp        Command = "HELP"
	CommandPath        Command = "PATH"
	CommandDeck        Command = "DECK"
	CommandPlay        Command = "PLAY"
	CommandQuit        Command = "QUIT"
)

func (c Command) ToString() string {
	return string(c)
}

type Reader interface {
	ReadNext() (Command, []string, error)
}

type reader struct {
}

func (r *reader) ReadNext() (Command, []string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("> ")

	commandStr, err := reader.ReadString('\n')
	if err != nil {
		return CommandUnspecified, nil, err
	}

	cmd, args, err := ToCommandAndArgs(commandStr)
	if err != nil {
		return CommandUnspecified, nil, err
	}

	return cmd, args, nil
}

var _ Reader = (*reader)(nil)

func ProvideReader() Reader {
	return &reader{}
}

func ToCommandAndArgs(commandStr string) (Command, []string, error) {
	commandStr = strings.ToUpper(strings.TrimSpace(commandStr))
	commandList := strings.Split(commandStr, " ")
	args := make([]string, 0, len(commandList)-1)
	for _, cmd := range commandList[1:] {
		trimmed := strings.TrimSpace(cmd)
		if len(trimmed) == 0 {
			continue
		}
		args = append(args, strings.TrimSpace(cmd))
	}

	switch commandList[0] {
	case CommandHelp.ToString():
		return CommandHelp, args, nil
	case CommandPath.ToString():
		return CommandPath, args, nil
	case CommandDeck.ToString():
		return CommandDeck, args, nil
	case CommandPlay.ToString():
		return CommandPlay, args, nil
	case CommandQuit.ToString():
		return CommandQuit, args, nil
	default:
		return CommandUnspecified, nil, errors.New("command not found")
	}
}
