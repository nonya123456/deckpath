package game_test

import (
	"testing"

	"github.com/nonya123456/deckpath/internal/game"
)

func TestCard_ToString(t *testing.T) {
	type fields struct {
		Effect game.CardEffect
		Amount int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "move_left",
			fields: fields{Effect: game.CardEffectMove, Amount: -2},
			want:   "Move Left By 2 Tiles",
		},
		{
			name:   "move_right",
			fields: fields{Effect: game.CardEffectMove, Amount: 1},
			want:   "Move Right By 1 Tiles",
		},
		{
			name:   "plant",
			fields: fields{Effect: game.CardEffectPlant},
			want:   "Plant",
		},
		{
			name:   "harvest",
			fields: fields{Effect: game.CardEffectHarvest},
			want:   "Harvest",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := game.Card{
				Effect: tt.fields.Effect,
				Amount: tt.fields.Amount,
			}
			if got := c.ToString(); got != tt.want {
				t.Errorf("Card.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
