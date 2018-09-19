package rdb_test

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/SekiguchiKai/clean-architecture-with-go/server/domain/model"
	"github.com/SekiguchiKai/clean-architecture-with-go/server/domain/repository"
	"github.com/SekiguchiKai/clean-architecture-with-go/server/infra/dao/rdb"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestNewProgrammingLangDAO(t *testing.T) {
	type args struct {
		ctx     context.Context
		manager rdb.SQLManagerInterface
	}
	tests := []struct {
		name string
		args args
		want repository.ProgrammingLangRepository
	}{
		{
			name: "適切な引数を与えると、ProgrammingLangDAOが返されること",
			args: args{
				ctx: context.Background(),
				manager: &rdb.SQLManager{
					Conn: &sql.DB{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rdb.NewProgrammingLangDAO(tt.args.manager); got == nil {
				t.Errorf("NewProgrammingLangDAO() = nil")
			}
		})
	}
}

func TestProgrammingLangDAO_Create(t *testing.T) {
	// sqlmockの設定を行う
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	type fields struct {
		SQLManager rdb.SQLManagerInterface
	}
	type args struct {
		ctx  context.Context
		lang *model.ProgrammingLang
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		want        *model.ProgrammingLang
		rowAffected int64
		wantErr     bool
	}{
		{
			name: "NameとFeatureを保持するProgrammingLangを与えられた場合、IDを付与したProgrammingLangを返すこと",
			fields: fields{
				SQLManager: &rdb.SQLManager{Conn: db},
			},
			args: args{
				ctx: context.Background(),
				lang: &model.ProgrammingLang{
					Name:      model.TestName,
					Feature:   model.TestFeature,
					CreatedAt: model.GetTestTime(time.September, 1),
					UpdatedAt: model.GetTestTime(time.September, 2),
				},
			},
			want: &model.ProgrammingLang{
				ID:        1,
				Name:      model.TestName,
				Feature:   model.TestFeature,
				CreatedAt: model.GetTestTime(time.September, 1),
				UpdatedAt: model.GetTestTime(time.September, 2),
			},
			rowAffected: 1,
			wantErr:     false,
		},
		{
			name: "Nameのみを保持するProgrammingLangを与えられた場合、IDを付与したProgrammingLangを返すこと",
			fields: fields{
				SQLManager: &rdb.SQLManager{Conn: db},
			},
			args: args{
				ctx: context.Background(),
				lang: &model.ProgrammingLang{
					Name:      model.TestName,
					CreatedAt: model.GetTestTime(time.September, 1),
					UpdatedAt: model.GetTestTime(time.September, 2),
				},
			},
			want: &model.ProgrammingLang{
				ID:        1,
				Name:      model.TestName,
				CreatedAt: model.GetTestTime(time.September, 1),
				UpdatedAt: model.GetTestTime(time.September, 2),
			},
			rowAffected: 1,
			wantErr:     false,
		},
		{
			name: "RowAffectedが1以外の場合、エラーを返すこと",
			fields: fields{
				SQLManager: &rdb.SQLManager{Conn: db},
			},
			args: args{
				ctx: context.Background(),
				lang: &model.ProgrammingLang{
					Name:      model.TestName,
					CreatedAt: model.GetTestTime(time.September, 1),
					UpdatedAt: model.GetTestTime(time.September, 2),
				},
			},
			want:        nil,
			rowAffected: 2,
			wantErr:     true,
		},
		{
			name: "空のProgrammingLangを与えられた場合、エラーを返すこと",
			fields: fields{
				SQLManager: &rdb.SQLManager{Conn: db},
			},
			args: args{
				ctx:  context.Background(),
				lang: &model.ProgrammingLang{},
			},
			want:        nil,
			rowAffected: 0,
			wantErr:     true,
		},
		{
			name: "Featureを保持する空のProgrammingLangを与えられた場合、エラーを返すこと",
			fields: fields{
				SQLManager: &rdb.SQLManager{Conn: db},
			},
			args: args{
				ctx:  context.Background(),
				lang: &model.ProgrammingLang{},
			},
			want:        nil,
			rowAffected: 0,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := "INSERT INTO programming_languages"
			prep := mock.ExpectPrepare(query)

			if tt.rowAffected == 0 {
				prep.ExpectExec().WithArgs(tt.args.lang.Name, tt.args.lang.Feature, tt.args.lang.CreatedAt, tt.args.lang.UpdatedAt).WillReturnError(fmt.Errorf(model.TestDBSomeErr))
			} else {
				prep.ExpectExec().WithArgs(tt.args.lang.Name, tt.args.lang.Feature, tt.args.lang.CreatedAt, tt.args.lang.UpdatedAt).WillReturnResult(sqlmock.NewResult(1, tt.rowAffected))
			}

			dao := rdb.NewProgrammingLangDAO(tt.fields.SQLManager)

			got, err := dao.Create(tt.args.ctx, tt.args.lang)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProgrammingLangDAO.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProgrammingLangDAO.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProgrammingLangDAO_List(t *testing.T) {
	testName1 := fmt.Sprintf("%s1", model.TestName)
	testName2 := fmt.Sprintf("%s2", model.TestName)
	testName3 := fmt.Sprintf("%s3", model.TestName)

	testFeature1 := fmt.Sprintf("%s1", model.TestFeature)
	testFeature2 := fmt.Sprintf("%s2", model.TestFeature)
	testFeature3 := fmt.Sprintf("%s3", model.TestFeature)

	// sqlmockの設定を行う
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	type fields struct {
		SQLManager rdb.SQLManagerInterface
	}
	type args struct {
		ctx   context.Context
		limit int
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.ProgrammingLang
		wantErr bool
	}{
		{
			name: "NameとFeatureを保持するProgrammingLangを与えられた場合、IDを付与したProgrammingLangを返すこと",
			fields: fields{
				SQLManager: &rdb.SQLManager{Conn: db},
			},
			args: args{
				ctx:   context.Background(),
				limit: 100,
			},
			want: []*model.ProgrammingLang{
				{
					ID:        1,
					Name:      testName1,
					Feature:   testFeature1,
					CreatedAt: model.GetTestTime(time.September, 1),
					UpdatedAt: model.GetTestTime(time.September, 2),
				},
				{
					ID:        2,
					Name:      testName2,
					Feature:   testFeature2,
					CreatedAt: model.GetTestTime(time.September, 3),
					UpdatedAt: model.GetTestTime(time.September, 4),
				},
				{
					ID:        3,
					Name:      testName3,
					Feature:   testFeature3,
					CreatedAt: model.GetTestTime(time.September, 5),
					UpdatedAt: model.GetTestTime(time.September, 6),
				},
			},
			wantErr: false,
		},
		{
			name: "NameとFeatureを保持するProgrammingLangを与えられた場合、IDを付与したProgrammingLangを返すこと",
			fields: fields{
				SQLManager: &rdb.SQLManager{Conn: db},
			},
			args: args{
				ctx:   context.Background(),
				limit: 100,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := "SELECT id, name, feature, created_at, updated_at FROM programming_languages ORDER BY name LIMIT \\?"
			prep := mock.ExpectPrepare(query)

			if tt.wantErr {
				prep.ExpectQuery().WillReturnError(fmt.Errorf(model.TestDBSomeErr))
			} else {
				rows := sqlmock.NewRows([]string{"id", "name", "feature", "created_at", "updated_at"}).
					AddRow(tt.want[0].ID, tt.want[0].Name, tt.want[0].Feature, tt.want[0].CreatedAt, tt.want[0].UpdatedAt).
					AddRow(tt.want[1].ID, tt.want[1].Name, tt.want[1].Feature, tt.want[1].CreatedAt, tt.want[1].UpdatedAt).
					AddRow(tt.want[2].ID, tt.want[2].Name, tt.want[2].Feature, tt.want[2].CreatedAt, tt.want[2].UpdatedAt)
				prep.ExpectQuery().WillReturnRows(rows)
			}

			dao := rdb.NewProgrammingLangDAO(tt.fields.SQLManager)

			got, err := dao.List(tt.args.ctx, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProgrammingLangDAO.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			for i := range got {
				if !reflect.DeepEqual(got[i], tt.want[i]) {
					t.Errorf("ProgrammingLangDAO.List() = %v, want %v", got[i], tt.want[i])
				}
			}
		})
	}
}

func TestProgrammingLangDAO_Read(t *testing.T) {
	// sqlmockの設定を行う
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	type fields struct {
		SQLManager rdb.SQLManagerInterface
	}
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.ProgrammingLang
		wantErr bool
	}{
		{
			name: "IDで指定したProgrammingLangが存在する場合、ProgrammingLangを1件返すこと",
			fields: fields{
				SQLManager: &rdb.SQLManager{Conn: db},
			},
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			want: &model.ProgrammingLang{
				ID:        1,
				Name:      model.TestName,
				Feature:   model.TestFeature,
				CreatedAt: model.GetTestTime(time.September, 1),
				UpdatedAt: model.GetTestTime(time.September, 2),
			},
			wantErr: false,
		},
		{
			name: "IDで指定したProgrammingLangが存在しない場合、エラー返すこと",
			fields: fields{
				SQLManager: &rdb.SQLManager{Conn: db},
			},
			args: args{
				ctx: context.Background(),
				id:  2,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := "SELECT id, name, feature, created_at, updated_at FROM programming_languages WHERE ID=\\?"
			prep := mock.ExpectPrepare(query)

			if tt.wantErr {
				prep.ExpectQuery().WillReturnError(fmt.Errorf(model.TestDBSomeErr))
			} else {
				rows := sqlmock.NewRows([]string{"id", "name", "feature", "created_at", "updated_at"}).
					AddRow(tt.want.ID, tt.want.Name, tt.want.Feature, tt.want.CreatedAt, tt.want.UpdatedAt)
				prep.ExpectQuery().WillReturnRows(rows)
			}

			dao := rdb.NewProgrammingLangDAO(tt.fields.SQLManager)

			got, err := dao.Read(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProgrammingLangDAO.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProgrammingLangDAO.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProgrammingLangDAO_ReadByName(t *testing.T) {
	// sqlmockの設定を行う
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	type fields struct {
		SQLManager rdb.SQLManagerInterface
	}
	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.ProgrammingLang
		wantErr bool
	}{
		{
			name: "IDで指定したProgrammingLangが1件存在する場合、ProgrammingLangを1件返すこと",
			fields: fields{
				SQLManager: &rdb.SQLManager{Conn: db},
			},
			args: args{
				ctx:  context.Background(),
				name: model.TestName,
			},
			want: &model.ProgrammingLang{
				ID:        1,
				Name:      model.TestName,
				Feature:   model.TestFeature,
				CreatedAt: model.GetTestTime(time.September, 1),
				UpdatedAt: model.GetTestTime(time.September, 2),
			},
			wantErr: false,
		},
		{
			name: "IDで指定したProgrammingLangが存在しない場合、エラー返すこと",
			fields: fields{
				SQLManager: &rdb.SQLManager{Conn: db},
			},
			args: args{
				ctx:  context.Background(),
				name: "test",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := "SELECT id, name, feature, created_at, updated_at FROM programming_languages WHERE name=\\? ORDER BY name LIMIT \\?"
			prep := mock.ExpectPrepare(query)

			if tt.wantErr {
				prep.ExpectQuery().WillReturnError(fmt.Errorf(model.TestDBSomeErr))
			} else {
				rows := sqlmock.NewRows([]string{"id", "name", "feature", "created_at", "updated_at"}).
					AddRow(tt.want.ID, tt.want.Name, tt.want.Feature, tt.want.CreatedAt, tt.want.UpdatedAt)
				prep.ExpectQuery().WillReturnRows(rows)
			}

			dao := rdb.NewProgrammingLangDAO(tt.fields.SQLManager)

			got, err := dao.ReadByName(tt.args.ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProgrammingLangDAO.ReadByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("ProgrammingLangDAO.ReadByName() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestProgrammingLangDAO_Update(t *testing.T) {
	// sqlmockの設定を行う
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	type fields struct {
		SQLManager rdb.SQLManagerInterface
	}
	type args struct {
		ctx  context.Context
		lang *model.ProgrammingLang
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		want        *model.ProgrammingLang
		rowAffected int64
		wantErr     bool
	}{
		{
			name: "NameとFeatureを保持するProgrammingLangを与えられた場合、IDを付与したProgrammingLangを返すこと",
			fields: fields{
				SQLManager: &rdb.SQLManager{Conn: db},
			},
			args: args{
				ctx: context.Background(),
				lang: &model.ProgrammingLang{
					Name:      model.TestName,
					Feature:   model.TestFeature,
					CreatedAt: model.GetTestTime(time.September, 1),
					UpdatedAt: model.GetTestTime(time.September, 2),
				},
			},
			want: &model.ProgrammingLang{
				ID:        1,
				Name:      model.TestName,
				Feature:   model.TestFeature,
				CreatedAt: model.GetTestTime(time.September, 1),
				UpdatedAt: model.GetTestTime(time.September, 2),
			},
			rowAffected: 1,
			wantErr:     false,
		},
		{
			name: "Nameのみを保持するProgrammingLangを与えられた場合、IDを付与したProgrammingLangを返すこと",
			fields: fields{
				SQLManager: &rdb.SQLManager{Conn: db},
			},
			args: args{
				ctx: context.Background(),
				lang: &model.ProgrammingLang{
					Name:      model.TestName,
					CreatedAt: model.GetTestTime(time.September, 1),
					UpdatedAt: model.GetTestTime(time.September, 2),
				},
			},
			want: &model.ProgrammingLang{
				ID:        1,
				Name:      model.TestName,
				CreatedAt: model.GetTestTime(time.September, 1),
				UpdatedAt: model.GetTestTime(time.September, 2),
			},
			rowAffected: 1,
			wantErr:     false,
		},
		{
			name: "RowAffectedが1以外の場合、エラーを返すこと",
			fields: fields{
				SQLManager: &rdb.SQLManager{Conn: db},
			},
			args: args{
				ctx: context.Background(),
				lang: &model.ProgrammingLang{
					Name:      model.TestName,
					CreatedAt: model.GetTestTime(time.September, 1),
					UpdatedAt: model.GetTestTime(time.September, 2),
				},
			},
			want:        nil,
			rowAffected: 2,
			wantErr:     true,
		},
		{
			name: "空のProgrammingLangを与えられた場合、エラーを返すこと",
			fields: fields{
				SQLManager: &rdb.SQLManager{Conn: db},
			},
			args: args{
				ctx:  context.Background(),
				lang: &model.ProgrammingLang{},
			},
			want:        nil,
			rowAffected: 0,
			wantErr:     true,
		},
		{
			name: "Featureのみを保持するProgrammingLangを与えられた場合、エラーを返すこと",
			fields: fields{
				SQLManager: &rdb.SQLManager{Conn: db},
			},
			args: args{
				ctx:  context.Background(),
				lang: &model.ProgrammingLang{},
			},
			want:        nil,
			rowAffected: 0,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := "UPDATE programming_languages SET name=\\?, feature=\\?, created_at=\\?, updated_at=\\? WHERE id=\\?"
			prep := mock.ExpectPrepare(query)

			if tt.wantErr {
				prep.ExpectExec().WithArgs(tt.args.lang.Name, tt.args.lang.Feature, tt.args.lang.CreatedAt, tt.args.lang.UpdatedAt, tt.args.lang.ID).WillReturnError(fmt.Errorf(model.TestDBSomeErr))
			} else {
				prep.ExpectExec().WithArgs(tt.args.lang.Name, tt.args.lang.Feature, tt.args.lang.CreatedAt, tt.args.lang.UpdatedAt, tt.args.lang.ID).WillReturnResult(sqlmock.NewResult(1, tt.rowAffected))
			}

			dao := rdb.NewProgrammingLangDAO(tt.fields.SQLManager)

			got, err := dao.Update(tt.args.ctx, tt.args.lang)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProgrammingLangDAO.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProgrammingLangDAO.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProgrammingLangDAO_Delete(t *testing.T) {
	// sqlmockの設定を行う
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	type fields struct {
		SQLManager rdb.SQLManagerInterface
	}
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		rowAffected int64
		wantErr     bool
	}{
		{
			name: "Nameのみを保持するProgrammingLangを与えられた場合、IDを付与したProgrammingLangを返すこと",
			fields: fields{
				SQLManager: &rdb.SQLManager{Conn: db},
			},
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			rowAffected: 1,
			wantErr:     false,
		},
		{
			name: "Nameのみを保持するProgrammingLangを与えられた場合、IDを付与したProgrammingLangを返すこと",
			fields: fields{
				SQLManager: &rdb.SQLManager{Conn: db},
			},
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			rowAffected: 2,
			wantErr:     true,
		},
		{
			name: "IDが空の場合、ProgrammingLangを与えられた場合、IDを付与したProgrammingLangを返すこと",
			fields: fields{
				SQLManager: &rdb.SQLManager{Conn: db},
			},
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			rowAffected: 1,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := "DELETE FROM programming_languages WHERE id=\\?"
			prep := mock.ExpectPrepare(query)

			if tt.wantErr {
				prep.ExpectExec().WithArgs(tt.args.id).WillReturnError(fmt.Errorf(model.TestDBSomeErr))
			} else {
				prep.ExpectExec().WithArgs(tt.args.id).WillReturnResult(sqlmock.NewResult(1, tt.rowAffected))
			}

			dao := rdb.NewProgrammingLangDAO(tt.fields.SQLManager)

			if err := dao.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("ProgrammingLangDAO.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
