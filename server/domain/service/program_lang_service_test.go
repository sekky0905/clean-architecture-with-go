package service

import (
	"testing"

	"github.com/SekiguchiKai/clean-architecture-with-go/server/domain/model"
)

func TestNewProgrammingLang(t *testing.T) {
	type args struct {
		name string
	}

	type wantErr struct {
		isErr bool
		err   error
	}

	invalidPropertyErr := &model.InvalidPropertyError{
		Property: model.PropertyName,
		Message:  model.NameShouldBeMoreThanOneUnderTheTwenty,
	}

	tests := []struct {
		name    string
		args    args
		want    *model.ProgrammingLang
		wantErr wantErr
	}{
		{
			name: "Nameが20文字の場合、ProgrammingLangを返す",
			args: args{
				name: "abcdefghijklmnopqrst",
			},
			wantErr: wantErr{
				isErr: false,
			},
		},
		{
			name: "Nameが21文字の場合、エラーを返す",
			args: args{
				name: "abcdefghijklmnopqrstu",
			},
			wantErr: wantErr{
				isErr: true,
				err:   invalidPropertyErr,
			},
		},
		{
			name: "Nameが空文字の場合、エラーを返す",
			args: args{
				name: "",
			},
			wantErr: wantErr{
				isErr: true,
				err:   invalidPropertyErr,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewProgrammingLang(tt.args.name)
			if (err != nil) != tt.wantErr.isErr {
				t.Errorf("NewProgrammingLang() error = %v, wantErr %v", err, tt.wantErr.isErr)
				return
			}

			if tt.wantErr.isErr {
				if err.Error() != tt.wantErr.err.Error() {
					t.Errorf("ValidateProgrammingLang() error = %v, wantErr %v", err.Error(), tt.wantErr.err.Error())
				}
			} else {
				if got == nil {
					t.Errorf("NewProgrammingLang() = %v, want not nil", got)
				}
			}
		})
	}
}

func TestValidateProgrammingLang(t *testing.T) {
	type args struct {
		name string
	}

	type wantErr struct {
		isErr bool
		err   error
	}

	invalidPropertyErr := &model.InvalidPropertyError{
		Property: model.PropertyName,
		Message:  model.NameShouldBeMoreThanOneUnderTheTwenty,
	}

	tests := []struct {
		name    string
		args    args
		wantErr wantErr
	}{
		{
			name: "Nameが20文字の場合、エラーを返さない",
			args: args{
				name: "abcdefghijklmnopqrst",
			},
			wantErr: wantErr{
				isErr: false,
			},
		},
		{
			name: "Nameが21文字の場合、エラーを返す",
			args: args{
				name: "abcdefghijklmnopqrstu",
			},
			wantErr: wantErr{
				isErr: true,
				err:   invalidPropertyErr,
			},
		},
		{
			name: "Nameが空文字の場合、エラーを返す",
			args: args{
				name: "",
			},
			wantErr: wantErr{
				isErr: true,
				err:   invalidPropertyErr,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateProgrammingLang(tt.args.name)
			if (err != nil) != tt.wantErr.isErr {
				t.Errorf("ValidateProgrammingLang() error = %v, wantErr %v", err, tt.wantErr.isErr)
			}

			if tt.wantErr.isErr {
				if err.Error() != tt.wantErr.err.Error() {
					t.Errorf("ValidateProgrammingLang() error = %v, wantErr %v", err.Error(), tt.wantErr.err.Error())
				}
			}
		})
	}
}
