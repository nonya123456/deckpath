package prompt_test

import (
	"testing"

	"github.com/nonya123456/deckpath/internal/prompt"
)

func Test_ToCommand(t *testing.T) {
	type args struct {
		commandStr string
	}
	tests := []struct {
		name    string
		args    args
		want    prompt.Command
		wantErr bool
	}{
		{
			name:    "empty",
			args:    args{commandStr: ""},
			want:    prompt.CommandUnspecified,
			wantErr: true,
		},
		{
			name:    "map",
			args:    args{commandStr: "pAth"},
			want:    prompt.CommandPath,
			wantErr: false,
		},
		{
			name:    "unspecified",
			args:    args{commandStr: "some command"},
			want:    prompt.CommandUnspecified,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := prompt.ToCommand(tt.args.commandStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("toCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("toCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
