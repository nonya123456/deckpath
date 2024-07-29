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
	CommandPath        Command = "PATH"
)

func (c Command) ToString() string {
	return string(c)
}

type Reader interface {
	ReadNext() (Command, error)
}

type reader struct {
}

func (r *reader) ReadNext() (Command, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("> ")

	commandStr, err := reader.ReadString('\n')
	if err != nil {
		return CommandUnspecified, err
	}

	cmd, err := ToCommand(commandStr)
	if err != nil {
		return CommandUnspecified, err
	}

	return cmd, nil
}

var _ Reader = (*reader)(nil)

func ProvideReader() Reader {
	return &reader{}
}

func ToCommand(commandStr string) (Command, error) {
	commandStr = strings.TrimSpace(commandStr)
	commandStr = strings.ToUpper(commandStr)

	switch commandStr {
	case CommandPath.ToString():
		return CommandPath, nil
	default:
		return CommandUnspecified, errors.New("failed to convert command")
	}
}
