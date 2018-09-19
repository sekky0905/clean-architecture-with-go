package input

import (
	"context"

	"github.com/SekiguchiKai/clean-architecture-with-go/server/domain/model"
)

// ProgrammingLangInputPort は、ProgrammingLangのInputPort。
type ProgrammingLangInputPort interface {
	List(ctx context.Context, limit int) ([]*model.ProgrammingLang, error)
	Get(ctx context.Context, id int) (*model.ProgrammingLang, error)
	Create(ctx context.Context, param *model.ProgrammingLang) (*model.ProgrammingLang, error)
	Update(ctx context.Context, id int, param *model.ProgrammingLang) (*model.ProgrammingLang, error)
	Delete(ctx context.Context, id int) error
}
