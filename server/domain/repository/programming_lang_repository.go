package repository

import (
	"context"

	"github.com/SekiguchiKai/clean-architecture-with-go/server/domain/model"
)

// ProgrammingLangRepository は、ProgrammingLangのRepository。
type ProgrammingLangRepository interface {
	List(ctx context.Context, limit int) ([]*model.ProgrammingLang, error)
	Create(ctx context.Context, lang *model.ProgrammingLang) (*model.ProgrammingLang, error)
	Read(ctx context.Context, id int) (*model.ProgrammingLang, error)
	ReadByName(ctx context.Context, name string) (*model.ProgrammingLang, error)
	Update(ctx context.Context, lang *model.ProgrammingLang) (*model.ProgrammingLang, error)
	Delete(ctx context.Context, id int) error
}
