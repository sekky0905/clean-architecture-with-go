package usecase

import "testing"

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
			name: "ターゲットが20、Maxが30、Minが10、Defaultが20のとき、20を返す",
			args: args{
				targetLimit:  20,
				maxLimit:     30,
				minLimit:     10,
				defaultLimit: 20,
			},
			want: 20,
		},
		{
			name: "ターゲットが40、Maxが30、Minが10、Defaultが20のとき、20を返す",
			args: args{
				targetLimit:  20,
				maxLimit:     30,
				minLimit:     10,
				defaultLimit: 20,
			},
			want: 20,
		},
		{
			name: "ターゲットが5、Maxが30、Minが10、Defaultが20のとき、20を返す",
			args: args{
				targetLimit:  20,
				maxLimit:     30,
				minLimit:     10,
				defaultLimit: 20,
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ManageLimit(tt.args.targetLimit, tt.args.maxLimit, tt.args.minLimit, tt.args.defaultLimit); got != tt.want {
				t.Errorf("ManageLimit() = %v, want %v", got, tt.want)
			}
		})
	}
}
