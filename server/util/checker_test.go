package util

import (
	"testing"
	"github.com/SekiguchiKai/clean-architecture-go/server/domain/model"
)

func TestIsEmpty(t *testing.T) {
	type args struct {
		target string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name :"空の場合は、trueを返す",
			args :args{target:""},
			want:true,
		},
		{
			name :"空でない場合は、falseを返す",
			args :args{target:model.TestName},
			want:false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmpty(tt.args.target); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
