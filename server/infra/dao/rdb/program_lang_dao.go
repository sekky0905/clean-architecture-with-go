package rdb

import (
	"context"
	"fmt"

	"github.com/SekiguchiKai/clean-architecture-with-go/server/domain/model"
	"github.com/SekiguchiKai/clean-architecture-with-go/server/domain/repository"
	"github.com/pkg/errors"
)

// ProgrammingLangDAO は、ProgrammingLangのDAO。
type ProgrammingLangDAO struct {
	SQLManager SQLManagerInterface
}

// NewProgrammingLangDAO は、ProgrammingLangDAO生成して返す。
func NewProgrammingLangDAO(manager SQLManagerInterface) repository.ProgrammingLangRepository {
	fmt.Printf("NewProgrammingLangDAO")

	return &ProgrammingLangDAO{
		SQLManager: manager,
	}
}

// ErrorMsg は、エラー文を生成し、返す。
func (dao *ProgrammingLangDAO) ErrorMsg(method string, err error) error {
	return &model.DBError{
		ModelName: model.ModelNameProgrammingLang,
		DBMethod:  method,
		Detail:    err.Error(),
	}
}

// Create は、レコードを1件生成する。
func (dao *ProgrammingLangDAO) Create(ctx context.Context, lang *model.ProgrammingLang) (*model.ProgrammingLang, error) {
	query := "INSERT INTO programming_languages (name, feature, created_at, updated_at) VALUES (?, ?, ?, ?)"
	stmt, err := dao.SQLManager.PrepareContext(ctx, query)
	if err != nil {
		return nil, dao.ErrorMsg(model.DBMethodCreate, err)
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, lang.Name, lang.Feature, lang.CreatedAt, lang.UpdatedAt)
	if err != nil {
		return nil, dao.ErrorMsg(model.DBMethodCreate, err)
	}

	affect, err := result.RowsAffected()
	if affect != 1 {
		err = fmt.Errorf("%s: %d ", TotalAffected, affect)
		return nil, dao.ErrorMsg(model.DBMethodUpdate, err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, dao.ErrorMsg(model.DBMethodCreate, err)
	}

	lang.ID = int(id)

	return lang, nil
}

// List は、レコードの一覧を取得して返す。
func (dao *ProgrammingLangDAO) List(ctx context.Context, limit int) ([]*model.ProgrammingLang, error) {
	query := "SELECT id, name, feature, created_at, updated_at FROM programming_languages ORDER BY name LIMIT ?"
	return dao.list(ctx, query, limit)
}

// list は、レコードの一覧を取得して返す。
func (dao *ProgrammingLangDAO) list(ctx context.Context, query string, args ...interface{}) ([]*model.ProgrammingLang, error) {
	stmt, err := dao.SQLManager.PrepareContext(ctx, query)
	if err != nil {
		return nil, dao.ErrorMsg(model.DBMethodList, err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, dao.ErrorMsg(model.DBMethodList, err)
	}
	defer rows.Close()

	langSlice := make([]*model.ProgrammingLang, 0)
	for rows.Next() {
		lang := &model.ProgrammingLang{}

		err = rows.Scan(
			&lang.ID,
			&lang.Name,
			&lang.Feature,
			&lang.CreatedAt,
			&lang.UpdatedAt,
		)

		if err != nil {
			return nil, dao.ErrorMsg(model.DBMethodList, err)
		}
		langSlice = append(langSlice, lang)
	}

	return langSlice, nil
}

// Read は、レコードを1件取得して返す。
func (dao *ProgrammingLangDAO) Read(ctx context.Context, id int) (*model.ProgrammingLang, error) {
	query := "SELECT id, name, feature, created_at, updated_at FROM programming_languages WHERE ID=?"

	stmt, err := dao.SQLManager.PrepareContext(ctx, query)
	if err != nil {
		return nil, dao.ErrorMsg(model.DBMethodRead, err)
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, id)
	lang := &model.ProgrammingLang{}

	err = row.Scan(
		&lang.ID,
		&lang.Name,
		&lang.Feature,
		&lang.CreatedAt,
		&lang.UpdatedAt,
	)

	if err != nil {
		return nil, dao.ErrorMsg(model.DBMethodRead, err)
	}

	return lang, nil
}

// ReadByName は、指定したNameを保持するレコードを1返す。
func (dao *ProgrammingLangDAO) ReadByName(ctx context.Context, name string) (*model.ProgrammingLang, error) {
	query := "SELECT id, name, feature, created_at, updated_at FROM programming_languages WHERE name=? ORDER BY name LIMIT ?"
	langSlice, err := dao.list(ctx, query, name, 1)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if len(langSlice) == 0 {
		return nil, &model.NoSuchDataError{
			Name:      name,
			ModelName: model.ModelNameProgrammingLang,
		}
	}

	return langSlice[0], nil
}

// Update は、レコードを1件更新する。
func (dao *ProgrammingLangDAO) Update(ctx context.Context, lang *model.ProgrammingLang) (*model.ProgrammingLang, error) {
	query := "UPDATE programming_languages SET name=?, feature=?, created_at=?, updated_at=? WHERE id=?"

	stmt, err := dao.SQLManager.PrepareContext(ctx, query)
	defer stmt.Close()

	if err != nil {
		return nil, dao.ErrorMsg(model.DBMethodUpdate, err)
	}

	result, err := stmt.ExecContext(ctx, lang.Name, lang.Feature, lang.CreatedAt, lang.UpdatedAt, lang.ID)
	if err != nil {
		return nil, dao.ErrorMsg(model.DBMethodUpdate, err)
	}

	affect, err := result.RowsAffected()
	if affect != 1 {
		err = fmt.Errorf("%s: %d ", TotalAffected, affect)
		return nil, dao.ErrorMsg(model.DBMethodUpdate, err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, dao.ErrorMsg(model.DBMethodUpdate, err)
	}

	lang.ID = int(id)

	return lang, nil
}

// Delete は、レコードを1件削除する。
func (dao *ProgrammingLangDAO) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM programming_languages WHERE id=?"

	stmt, err := dao.SQLManager.PrepareContext(ctx, query)
	if err != nil {
		return dao.ErrorMsg(model.DBMethodDelete, err)
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return dao.ErrorMsg(model.DBMethodDelete, err)
	}

	affect, err := result.RowsAffected()
	if err != nil {
		return dao.ErrorMsg(model.DBMethodDelete, err)
	}
	if affect != 1 {
		err = fmt.Errorf("%s: %d ", TotalAffected, affect)
		return dao.ErrorMsg(model.DBMethodDelete, err)
	}

	return nil
}
