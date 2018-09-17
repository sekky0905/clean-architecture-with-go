package repository

import (
	"github.com/SekiguchiKai/clean-architecture-with-go/server/domain/model"
)

// ProgrammingLangRepository は、ProgrammingLangのRepository。
type ProgrammingLangRepository interface {
	List(limit int) ([]*model.ProgrammingLang, error)
	Create(lang *model.ProgrammingLang) (*model.ProgrammingLang, error)
	Read(id int) (*model.ProgrammingLang, error)
	ReadByName(name string) (*model.ProgrammingLang, error)
	Update(lang *model.ProgrammingLang) (*model.ProgrammingLang, error)
	Delete(id int) error
}
