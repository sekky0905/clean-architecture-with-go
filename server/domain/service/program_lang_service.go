package service

import (
	"github.com/SekiguchiKai/clean-architecture-with-go/server/domain/model"
	"github.com/pkg/errors"
	"github.com/SekiguchiKai/clean-architecture-with-go/server/util"
)

// NewProgrammingLang は、ProgrammingLangを生成し、返す。
func NewProgrammingLang(name string) (*model.ProgrammingLang, error) {
	if err := ValidateProgrammingLang(name); err != nil {
		return nil, errors.WithStack(err)
	}

	return &model.ProgrammingLang{
		Name: name,
	}, nil
}

// ValidateProgrammingLang は、ProgrammingLangの生成に必要な属性に与えられる引数をチェックする。
func ValidateProgrammingLang(name string) error {
	if util.IsEmpty(name) || len(name) > 20 {
		return &model.InvalidPropertyError{
			Property: model.PropertyName,
			Message:  model.NameShouldBeMoreThanOneUnderTheTwenty,
		}
	}
	return nil
}
