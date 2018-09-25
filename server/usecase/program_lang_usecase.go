package usecase

import (
	"context"
	"time"

	"github.com/SekiguchiKai/clean-architecture-with-go/server/domain/model"
	"github.com/SekiguchiKai/clean-architecture-with-go/server/domain/repository"
	"github.com/SekiguchiKai/clean-architecture-with-go/server/usecase/input"
	"github.com/pkg/errors"
)

// ProgrammingLangUseCase は、ProgrammingLangのUseCase。
type ProgrammingLangUseCase struct {
	Repo repository.ProgrammingLangRepository
}

// NewProgrammingLangUseCase は、ProgrammingLangUseCaseを生成し、返す。
func NewProgrammingLangUseCase(repo repository.ProgrammingLangRepository) input.ProgrammingLangInputPort {
	return &ProgrammingLangUseCase{
		Repo: repo,
	}
}

// List は、ProgrammingLangの一覧を返す。
func (u *ProgrammingLangUseCase) List(ctx context.Context, limit int) ([]*model.ProgrammingLang, error) {
	return u.Repo.List(ctx, limit)
}

// Get は、ProgrammingLang1件返す。
func (u *ProgrammingLangUseCase) Get(ctx context.Context, id int) (*model.ProgrammingLang, error) {
	return u.Repo.Read(ctx, id)
}

// Create は、ProgrammingLangを生成する。
func (u *ProgrammingLangUseCase) Create(ctx context.Context, param *model.ProgrammingLang) (*model.ProgrammingLang, error) {
	lang, err := u.Repo.ReadByName(ctx, param.Name)
	if lang != nil {
		return nil, &model.AlreadyExistError{
			ID:        lang.ID,
			Name:      lang.Name,
			ModelName: model.ModelNameProgrammingLang,
		}
	}

	if _, ok := errors.Cause(err).(*model.AlreadyExistError); !ok {
		return nil, errors.WithStack(err)
	}

	param.CreatedAt = time.Now().UTC()
	param.UpdatedAt = time.Now().UTC()

	lang, err = u.Repo.Create(ctx, param)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return lang, nil
}

// Update は、ProgrammingLangを更新する。
func (u *ProgrammingLangUseCase) Update(ctx context.Context, id int, param *model.ProgrammingLang) (*model.ProgrammingLang, error) {
	lang, err := u.Repo.Read(ctx, id)
	if lang == nil {
		return nil, &model.NoSuchDataError{
			ID:        id,
			Name:      param.Name,
			ModelName: model.ModelNameProgrammingLang,
		}
	} else if err != nil {
		return nil, errors.WithStack(err)
	}

	lang.ID = id
	lang.Name = param.Name
	lang.Feature = param.Feature
	lang.UpdatedAt = time.Now().UTC()

	return u.Repo.Update(ctx, lang)
}

// Delete は、ProgrammingLangを削除する。
func (u *ProgrammingLangUseCase) Delete(ctx context.Context, id int) error {
	lang, err := u.Repo.Read(ctx, id)
	if lang == nil {
		return  &model.NoSuchDataError{
			ID:        id,
			ModelName: model.ModelNameProgrammingLang,
		}
	} else if err != nil {
		return  errors.WithStack(err)
	}

	return u.Repo.Delete(ctx, id)
}
