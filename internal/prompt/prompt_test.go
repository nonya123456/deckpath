package prompt

import "testing"

func Test_toCommand(t *testing.T) {
	type args struct {
		commandStr string
	}
	tests := []struct {
		name    string
		args    args
		want    Command
		wantErr bool
	}{
		{
			name:    "empty",
			args:    args{commandStr: ""},
			want:    CommandUnspecified,
			wantErr: true,
		},
		{
			name:    "map",
			args:    args{commandStr: "mAp"},
			want:    CommandMap,
			wantErr: false,
		},
		{
			name:    "unspecified",
			args:    args{commandStr: "some command"},
			want:    CommandUnspecified,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toCommand(tt.args.commandStr)
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
