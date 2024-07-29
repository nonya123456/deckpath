package prompt_test

import (
	"reflect"
	"testing"

	"github.com/nonya123456/deckpath/internal/prompt"
)

func TestToCommandAndArgs(t *testing.T) {
	type args struct {
		commandStr string
	}
	tests := []struct {
		name    string
		args    args
		want    prompt.Command
		want1   []string
		wantErr bool
	}{
		{
			name:    "empty",
			args:    args{commandStr: ""},
			want:    prompt.CommandUnspecified,
			want1:   nil,
			wantErr: true,
		},
		{
			name:    "help",
			args:    args{commandStr: "help"},
			want:    prompt.CommandHelp,
			want1:   make([]string, 0),
			wantErr: false,
		},
		{
			name:    "path",
			args:    args{commandStr: "pAth"},
			want:    prompt.CommandPath,
			want1:   make([]string, 0),
			wantErr: false,
		},
		{
			name:    "deck",
			args:    args{commandStr: "deCk"},
			want:    prompt.CommandDeck,
			want1:   make([]string, 0),
			wantErr: false,
		},
		{
			name:    "play",
			args:    args{commandStr: "PLaY 1  2    3 4"},
			want:    prompt.CommandPlay,
			want1:   []string{"1", "2", "3", "4"},
			wantErr: false,
		},
		{
			name:    "quit",
			args:    args{commandStr: "Quit"},
			want:    prompt.CommandQuit,
			want1:   make([]string, 0),
			wantErr: false,
		},
		{
			name:    "unspecified",
			args:    args{commandStr: "some command"},
			want:    prompt.CommandUnspecified,
			want1:   nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := prompt.ToCommandAndArgs(tt.args.commandStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToCommandAndArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToCommandAndArgs() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ToCommandAndArgs() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
