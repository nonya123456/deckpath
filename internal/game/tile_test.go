package game

import "testing"

func TestTile_ToString(t *testing.T) {
	type fields struct {
		Growth int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "normal",
			fields: fields{Growth: 1},
			want:   "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Tile{
				Growth: tt.fields.Growth,
			}
			if got := tr.ToString(); got != tt.want {
				t.Errorf("Tile.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
