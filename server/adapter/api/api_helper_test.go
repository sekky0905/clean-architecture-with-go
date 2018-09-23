package api_test

import (
	"testing"
	"github.com/SekiguchiKai/clean-architecture-with-go/server/adapter/api"
)

func TestManageLimit(t *testing.T) {
	type args struct {
		targetLimit  int
		maxLimit     int
		defaultLimit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ターゲットが10、Maxが30、Defaultが20のとき、10を返す",
			args: args{
				targetLimit:  10,
				maxLimit:     30,
				defaultLimit: 20,
			},
			want: 10,
		},
		{
			name: "ターゲットが40、Maxが30、Defaultが20のとき、20を返す",
			args: args{
				targetLimit:  20,
				maxLimit:     30,
				defaultLimit: 20,
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := api.ManageLimit(tt.args.targetLimit, tt.args.maxLimit,  tt.args.defaultLimit); got != tt.want {
				t.Errorf("ManageLimit() = %v, want %v", got, tt.want)
			}
		})
	}
}
