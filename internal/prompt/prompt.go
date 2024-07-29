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
	CommandMap         Command = "MAP"
)

type Reader interface {
	ReadCommand() (Command, error)
}

type reader struct {
}

func (r *reader) ReadCommand() (Command, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("> ")

	commandStr, err := reader.ReadString('\n')
	if err != nil {
		return CommandUnspecified, err
	}

	cmd, err := toCommand(commandStr)
	if err != nil {
		return CommandUnspecified, err
	}

	return cmd, nil
}

var _ Reader = (*reader)(nil)

func ProvidePromptReader() Reader {
	return &reader{}
}

func toCommand(commandStr string) (Command, error) {
	commandStr = strings.TrimSpace(commandStr)
	commandStr = strings.ToUpper(commandStr)

	switch commandStr {
	case "MAP":
		return CommandMap, nil
	default:
		return CommandUnspecified, errors.New("failed to convert command")
	}
}
