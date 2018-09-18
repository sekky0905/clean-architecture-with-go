package input

import "github.com/SekiguchiKai/clean-architecture-with-go/server/domain/model"

// ProgrammingLangInputPort は、ProgrammingLangのInputPort。
type ProgrammingLangInputPort interface {
	List(limit int) ([]*model.ProgrammingLang, error)
	Get(id int) (*model.ProgrammingLang, error)
	Create(param *model.ProgrammingLang) (*model.ProgrammingLang, error)
	Update(param *model.ProgrammingLang) (*model.ProgrammingLang, error)
	Delete(id int) error
}
