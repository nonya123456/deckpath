package game

type Card struct{}

func NewDeck(num int) []Card {
	return make([]Card, num)
}
