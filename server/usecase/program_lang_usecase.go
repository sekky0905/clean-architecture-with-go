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
	limit = ManageLimit(limit, MaxLimit, MinLimit, DefaultLimit)
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

	if err != nil {
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
func (u *ProgrammingLangUseCase) Update(ctx context.Context, param *model.ProgrammingLang) (*model.ProgrammingLang, error) {
	lang, err := u.Repo.Read(ctx, param.ID)
	if lang == nil {
		return nil, &model.NoSuchDataError{
			ID:        param.ID,
			Name:      param.Name,
			ModelName: model.ModelNameProgrammingLang,
		}
	}

	if err != nil {
		return nil, errors.WithStack(err)
	}

	param.UpdatedAt = time.Now().UTC()
	return u.Repo.Update(ctx, param)
}

// Delete は、ProgrammingLangを削除する。
func (u *ProgrammingLangUseCase) Delete(ctx context.Context, id int) error {
	return u.Repo.Delete(ctx, id)
}
