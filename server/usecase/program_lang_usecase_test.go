package usecase

import (
	"context"
	"reflect"
	"testing"

	"github.com/SekiguchiKai/clean-architecture-with-go/server/domain/model"
	"github.com/SekiguchiKai/clean-architecture-with-go/server/domain/repository"
	"github.com/SekiguchiKai/clean-architecture-with-go/server/infra/dao/mock"
	"github.com/golang/mock/gomock"
)

func TestNewProgrammingLangUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_repository.NewMockProgrammingLangRepository(ctrl)

	type args struct {
		repo repository.ProgrammingLangRepository
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "適切な引数を与えると、ProgrammingLangUseCaseが返されること",
			args: args{
				repo: mock,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProgrammingLangUseCase(tt.args.repo); got == nil {
				t.Errorf("NewProgrammingLangUseCase() = %v, want not nil", got)
			}
		})
	}
}

func TestProgrammingLangUseCase_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_repository.NewMockProgrammingLangRepository(ctrl)

	type fields struct {
		Repo repository.ProgrammingLangRepository
	}
	type args struct {
		ctx   context.Context
		limit int
	}

	type wantErr struct {
		isErr bool
		err   error
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.ProgrammingLang
		wantErr wantErr
	}{
		{
			name: "limitを20件にした時に、20件のProgrammingLangを含んだProgrammingLangを返すこと",
			fields: fields{
				Repo: mock,
			},
			args: args{
				ctx:   context.Background(),
				limit: 20,
			},
			want:    model.CreateProgrammingLangs(20),
			wantErr: wantErr{
				isErr:false,
			},
		},
		{
			name: "サーバー側のエラーが発生した場合、ステータスコード500とエラーメッセージを返すこと",
			fields: fields{
				Repo: mock,
			},
			args: args{
				ctx:   context.Background(),
				limit: 20,
			},
			want:    nil,
			wantErr: wantErr{
				isErr:true,
				err:&model.DBError{
					ModelName: model.ModelNameProgrammingLang,
					DBMethod:  model.DBMethodList,
					Detail:    model.TestDBSomeErr,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &ProgrammingLangUseCase{
				Repo: tt.fields.Repo,
			}

			mock.EXPECT().List(tt.args.ctx, tt.args.limit).Return(tt.want, tt.wantErr.err)

			got, err := u.List(tt.args.ctx, tt.args.limit)
			if (err != nil) != tt.wantErr.isErr {
				t.Errorf("ProgrammingLangUseCase.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProgrammingLangUseCase.List() = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(err, tt.wantErr.err) {
				t.Errorf("ProgrammingLangUseCase.List() = %v, want %v", err,  tt.wantErr.err)
			}
		})
	}
}

func TestProgrammingLangUseCase_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_repository.NewMockProgrammingLangRepository(ctrl)

	wantErrValue := &model.NoSuchDataError{
		ID:        1,
		ModelName: model.ModelNameProgrammingLang,
	}

	type fields struct {
		Repo repository.ProgrammingLangRepository
	}
	type args struct {
		ctx context.Context
		id  int
	}

	type wantErr struct {
		isErr bool
		err   error
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.ProgrammingLang
		wantErr wantErr
	}{
		{
			name: "IDで指定したProgrammingLang存在する場合、ProgrammingLangを1件返すこと",
			fields: fields{
				Repo: mock,
			},
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			want: model.CreateProgrammingLangs(1)[0],
			wantErr: wantErr{
				isErr: false,
				err:   nil,
			},
		},
		{
			name: "IDで指定したProgrammingLang存在しない場合、nilとエラーを返すこと",
			fields: fields{
				Repo: mock,
			},
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			want: model.CreateProgrammingLangs(1)[0],
			wantErr: wantErr{
				isErr: true,
				err:   wantErrValue,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &ProgrammingLangUseCase{
				Repo: tt.fields.Repo,
			}

			mock.EXPECT().Read(tt.args.ctx, tt.args.id).Return(tt.want, tt.wantErr.err)

			got, err := u.Get(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr.isErr {
				t.Errorf("ProgrammingLangUseCase.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr.isErr {
				if err.Error() != tt.wantErr.err.Error() {
					t.Errorf("ProgrammingLangUseCase.Get() error = %v, wantErr %v", err.Error(), tt.wantErr.err.Error())
				}
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProgrammingLangUseCase.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProgrammingLangUseCase_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_repository.NewMockProgrammingLangRepository(ctrl)

	lang := model.CreateProgrammingLangs(1)[0]

	wantErrValue := &model.AlreadyExistError{
		ID:        1,
		ModelName: model.ModelNameProgrammingLang,
		Name:      model.CreateProgrammingLangs(1)[0].Name,
	}

	type fields struct {
		Repo repository.ProgrammingLangRepository
	}
	type args struct {
		ctx   context.Context
		param *model.ProgrammingLang
	}

	type readWant struct {
		result *model.ProgrammingLang
		err    error
	}

	type wantErr struct {
		isErr bool
		err   error
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		want     *model.ProgrammingLang
		readWant readWant
		wantErr  wantErr
	}{
		{
			name: "同一のProgrammingLang存在しない場合、ProgrammingLangを登録すること",
			fields: fields{
				Repo: mock,
			},
			args: args{
				ctx:   context.Background(),
				param: lang,
			},
			want: lang,
			readWant: readWant{
				result: nil,
				err: &model.NoSuchDataError{
					Name:      lang.Name,
					ModelName: model.ModelNameProgrammingLang,
				},
			},
			wantErr: wantErr{
				isErr: false,
				err:   nil,
			},
		},
		{
			name: "同一のProgrammingLang存在する場合、nilとエラーを返すこと",
			fields: fields{
				Repo: mock,
			},
			args: args{
				ctx:   context.Background(),
				param: lang,
			},
			want: nil,
			readWant: readWant{
				result: lang,
				err:    wantErrValue,
			},
			wantErr: wantErr{
				isErr: true,
				err:   wantErrValue,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &ProgrammingLangUseCase{
				Repo: tt.fields.Repo,
			}

			mock.EXPECT().ReadByName(tt.args.ctx, tt.args.param.Name).Return(tt.readWant.result, tt.readWant.err)

			if !tt.wantErr.isErr {
				mock.EXPECT().Create(tt.args.ctx, tt.args.param).Return(tt.want, tt.wantErr.err)
			}

			got, err := u.Create(tt.args.ctx, tt.args.param)

			if (err != nil) != tt.wantErr.isErr {
				t.Errorf("ProgrammingLangUseCase.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr.isErr {
				if err.Error() != tt.wantErr.err.Error() {
					t.Errorf("ProgrammingLangUseCase.Create() error = %v, wantErr %v", err.Error(), tt.wantErr.err.Error())
				}
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProgrammingLangUseCase.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProgrammingLangUseCase_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_repository.NewMockProgrammingLangRepository(ctrl)

	lang := model.CreateProgrammingLangs(1)[0]

	type fields struct {
		Repo repository.ProgrammingLangRepository
	}
	type args struct {
		ctx   context.Context
		id    int
		param *model.ProgrammingLang
	}

	type readWant struct {
		result *model.ProgrammingLang
		err    error
	}

	type wantErr struct {
		isErr bool
		err   error
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		want     *model.ProgrammingLang
		readWant readWant
		wantErr  wantErr
	}{
		{
			name: "同一のProgrammingLang存在する場合、ProgrammingLangを更新すること",
			fields: fields{
				Repo: mock,
			},
			args: args{
				ctx:   context.Background(),
				id:    1,
				param: lang,
			},
			want: lang,
			readWant: readWant{
				result: lang,
				err:    nil,
			},
			wantErr: wantErr{
				isErr: false,
				err:   nil,
			},
		},
		{
			name: "指定したProgrammingLang存在しない場合、nilとエラーを返すこと",
			fields: fields{
				Repo: mock,
			},
			args: args{
				ctx:   context.Background(),
				id:    100,
				param: lang,
			},
			want: nil,
			readWant: readWant{
				result: nil,
				err: &model.DBError{
					ModelName: model.ModelNameProgrammingLang,
					DBMethod:  model.DBMethodRead,
					Detail:    model.TestDBSomeErr,
				},
			},
			wantErr: wantErr{
				isErr: true,
				err: &model.NoSuchDataError{
					ID:        100,
					ModelName: model.ModelNameProgrammingLang,
					Name:      lang.Name,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &ProgrammingLangUseCase{
				Repo: tt.fields.Repo,
			}

			mock.EXPECT().Read(tt.args.ctx, tt.args.id).Return(tt.readWant.result, tt.readWant.err)

			if !tt.wantErr.isErr {
				mock.EXPECT().Update(tt.args.ctx, tt.args.param).Return(tt.want, tt.wantErr.err)
			}

			got, err := u.Update(tt.args.ctx, tt.args.id, tt.args.param)
			if (err != nil) != tt.wantErr.isErr {
				t.Errorf("ProgrammingLangUseCase.Update() error = %v, wantErr %v", err, tt.wantErr.isErr)
				return
			}

			if tt.wantErr.isErr {
				if err.Error() != tt.wantErr.err.Error() {
					t.Errorf("ProgrammingLangUseCase.Update() error = %v, wantErr %v", err.Error(), tt.wantErr.err.Error())
				}
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProgrammingLangUseCase.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProgrammingLangUseCase_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_repository.NewMockProgrammingLangRepository(ctrl)

	err := &model.NoSuchDataError{
		ID:        1,
		ModelName: model.ModelNameProgrammingLang,
	}

	type fields struct {
		Ctx  context.Context
		Repo repository.ProgrammingLangRepository
	}
	type args struct {
		ctx context.Context
		id  int
	}

	type wantErr struct {
		isErr bool
		err   error
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr wantErr
	}{
		{
			name: "同一のProgrammingLang存在する場合、ProgrammingLangを更新すること",
			fields: fields{
				Repo: mock,
			},
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			wantErr: wantErr{
				isErr: false,
				err:   nil,
			},
		},
		{
			name: "IDで指定したProgrammingLang存在しない場合、nilとエラーを返すこと",
			fields: fields{
				Repo: mock,
			},
			args: args{
				ctx: context.Background(),
				id:  2,
			},
			wantErr: wantErr{
				isErr: true,
				err:   err,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &ProgrammingLangUseCase{
				Repo: tt.fields.Repo,
			}

			mock.EXPECT().Delete(tt.args.ctx, tt.args.id).Return(tt.wantErr.err)

			if err := u.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr.isErr {
				t.Errorf("ProgrammingLangUseCase.Delete() error = %v, wantErr %v", err, tt.wantErr.isErr)
			}

			if tt.wantErr.isErr {
				if err.Error() != tt.wantErr.err.Error() {
					t.Errorf("ProgrammingLangUseCase.Update() error = %v, wantErr %v", err.Error(), tt.wantErr.err.Error())
				}
			}
		})
	}
}
