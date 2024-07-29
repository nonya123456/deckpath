package wire

import (
	"github.com/google/wire"
	"github.com/nonya123456/deckpath/internal/prompt"
)

var ProviderSet = wire.NewSet(
	prompt.ProvidePromptReader,
)
