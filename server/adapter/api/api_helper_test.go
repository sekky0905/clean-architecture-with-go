package api_test

import (
	"testing"
	"github.com/SekiguchiKai/clean-architecture-with-go/server/adapter/api"
)

func TestManageLimit(t *testing.T) {
	type args struct {
		targetLimit  int
		maxLimit     int
		minLimit     int
		defaultLimit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ターゲットが10、Maxが30、Minが5、Defaultが20のとき、10を返す",
			args: args{
				targetLimit:  10,
				maxLimit:     30,
				minLimit:5,
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
		{
			name: "ターゲットが31、Maxが30、Defaultが20のとき、20を返す",
			args: args{
				targetLimit:  31,
				maxLimit:     30,
				defaultLimit: 20,
			},
			want: 20,
		},
		{
			name: "ターゲットが29、Maxが30、Defaultが20のとき、29を返す",
			args: args{
				targetLimit:  29,
				maxLimit:     30,
				defaultLimit: 29,
			},
			want: 29,
		},
		{
			name: "ターゲットが30、Maxが30、Defaultが20のとき、30を返す",
			args: args{
				targetLimit:  30,
				maxLimit:     30,
				defaultLimit: 30,
			},
			want: 30,
		},
		{
			name: "ターゲットが4、Maxが30、Minが5、Defaultが20のとき、20を返す",
			args: args{
				targetLimit:  4,
				maxLimit:     30,
				minLimit:5,
				defaultLimit: 20,
			},
			want: 20,
		},
		{
			name: "ターゲットが6、Maxが30、Minが5、Defaultが20のとき、6を返す",
			args: args{
				targetLimit:  6,
				maxLimit:     30,
				minLimit:5,
				defaultLimit: 20,
			},
			want: 6,
		},
		{
			name: "ターゲットが5、Maxが30、Minが5、Defaultが20のとき、5を返す",
			args: args{
				targetLimit:  5,
				maxLimit:     30,
				minLimit:5,
				defaultLimit: 20,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := api.ManageLimit(tt.args.targetLimit, tt.args.maxLimit,  tt.args.minLimit, tt.args.defaultLimit); got != tt.want {
				t.Errorf("ManageLimit() = %v, want %v", got, tt.want)
			}
		})
	}
}
