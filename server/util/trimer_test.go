package util

import "testing"

func TestTrimDoubleQuotes(t *testing.T) {
	type args struct {
		target string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ダブルクォーテーションが1つ存在する場合、ダブルクォーテーションを削除した文字列を返す",
			args: args{
				target: "\"test",
			},
			want: "test",
		},
		{
			name: "ダブルクォーテーションが3つ存在する場合、ダブルクォーテーションを削除した文字列を返す",
			args: args{
				target: "\"\"test\"",
			},
			want: "test",
		},
		{
			name: "ダブルクォーテーションが存在しない場合、そのままの文字列を返す",
			args: args{
				target: "test",
			},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimDoubleQuotes(tt.args.target); got != tt.want {
				t.Errorf("TrimDoubleQuotes() = %v, want %v", got, tt.want)
			}
		})
	}
}
