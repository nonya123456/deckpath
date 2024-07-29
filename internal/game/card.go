package game

import (
	"fmt"
	"math/rand"
)

type CardEffect int

const (
	CardEffectMove CardEffect = iota
	CardEffectPlant
	CardEffectHarvest
)

const (
	MaxMoveAmount = 3
)

type Card struct {
	Effect CardEffect
	Amount int
}

func (c Card) ToString() string {
	switch c.Effect {
	case CardEffectMove:
		if c.Amount > 0 {
			return fmt.Sprintf("Move Right For %d Tiles", c.Amount)
		}
		return fmt.Sprintf("Move Left For %d Tiles", -c.Amount)
	case CardEffectPlant:
		return "Plant"
	case CardEffectHarvest:
		return "Harvest"
	}

	return "Unknown"
}

func NewCard() Card {
	effects := []CardEffect{CardEffectMove, CardEffectPlant, CardEffectHarvest}
	idx := rand.Intn(len(effects))
	card := Card{
		Effect: effects[idx],
	}

	if card.Effect == CardEffectMove {
		card.Amount = rand.Intn(MaxMoveAmount) + 1
		if rand.Intn(2) == 0 {
			card.Amount = -card.Amount
		}
	}

	return card
}

func NewDeck(num int) []Card {
	deck := make([]Card, 0, num)
	for i := 0; i < num; i += 1 {
		deck = append(deck, NewCard())
	}

	return deck
}
