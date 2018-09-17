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
	Ctx  context.Context
	Repo repository.ProgrammingLangRepository
}

// NewProgrammingLangUseCase は、ProgrammingLangUseCaseを生成し、返す。
func NewProgrammingLangUseCase(ctx context.Context, repo repository.ProgrammingLangRepository) input.ProgrammingLangInputPort {
	return &ProgrammingLangUseCase{
		Ctx:  ctx,
		Repo: repo,
	}
}

// List は、ProgrammingLangの一覧を返す。
func (u *ProgrammingLangUseCase) List(limit int) ([]*model.ProgrammingLang, error) {
	limit = ManageLimit(limit, MaxLimit, MinLimit, DefaultLimit)
	return u.Repo.List(limit)
}

// Get は、ProgrammingLang1件返す。
func (u *ProgrammingLangUseCase) Get(id int) (*model.ProgrammingLang, error) {
	return u.Repo.Read(id)
}

// Create は、ProgrammingLangを生成する。
func (u *ProgrammingLangUseCase) Create(param *model.ProgrammingLang) (*model.ProgrammingLang, error) {
	lang, err := u.Repo.ReadByName(param.Name)
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

	lang, err = u.Repo.Create(param)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return lang, nil
}

// Update は、ProgrammingLangを更新する。
func (u *ProgrammingLangUseCase) Update(param *model.ProgrammingLang) (*model.ProgrammingLang, error) {
	lang, err := u.Repo.Read(param.ID)
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
	return u.Repo.Update(param)
}

// Delete は、ProgrammingLangを削除する。
func (u *ProgrammingLangUseCase) Delete(id int) error {
	return u.Repo.Delete(id)
}
