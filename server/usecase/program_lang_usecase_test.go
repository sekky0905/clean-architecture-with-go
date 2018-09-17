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
		ctx  context.Context
		repo repository.ProgrammingLangRepository
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "適切な引数を与えると、ProgrammingLangUseCaseが返されること",
			args: args{
				ctx:  context.TODO(),
				repo: mock,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProgrammingLangUseCase(tt.args.ctx, tt.args.repo); got == nil {
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
		Ctx  context.Context
		Repo repository.ProgrammingLangRepository
	}
	type args struct {
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
			name: "limitを6件にした時に、6件のProgrammingLangを含んだProgrammingLangを返すこと",
			fields: fields{
				Ctx:  context.TODO(),
				Repo: mock,
			},
			args:    args{limit: 6},
			want:    model.CreateProgrammingLangs(6),
			wantErr: false,
		},
		{
			name: "limitを99件にした時に、99件のProgrammingLangを含んだProgrammingLangを返すこと",
			fields: fields{
				Ctx:  context.TODO(),
				Repo: mock,
			},
			args:    args{limit: 99},
			want:    model.CreateProgrammingLangs(99),
			wantErr: false,
		},
		{
			name: "limitを20件にした時に、20件のProgrammingLangを含んだProgrammingLangを返すこと",
			fields: fields{
				Ctx:  context.TODO(),
				Repo: mock,
			},
			args:    args{limit: 20},
			want:    model.CreateProgrammingLangs(20),
			wantErr: false,
		},
		{
			name: "limitを101件にした時に、21件のProgrammingLangを含んだProgrammingLangを返すこと",
			fields: fields{
				Ctx:  context.TODO(),
				Repo: mock,
			},
			args:    args{limit: 21},
			want:    model.CreateProgrammingLangs(21),
			wantErr: false,
		},
		{
			name: "limitを4件にした時に、20件のProgrammingLangを含んだProgrammingLangを返すこと",
			fields: fields{
				Ctx:  context.TODO(),
				Repo: mock,
			},
			args:    args{limit: 21},
			want:    model.CreateProgrammingLangs(21),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &ProgrammingLangUseCase{
				Ctx:  tt.fields.Ctx,
				Repo: tt.fields.Repo,
			}

			mock.EXPECT().List(tt.args.limit).Return(tt.want, nil)

			got, err := u.List(tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProgrammingLangUseCase.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProgrammingLangUseCase.List() = %v, want %v", got, tt.want)
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
		Ctx  context.Context
		Repo repository.ProgrammingLangRepository
	}
	type args struct {
		id int
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
				Ctx:  context.TODO(),
				Repo: mock,
			},
			args: args{
				id: 1,
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
				Ctx:  context.TODO(),
				Repo: mock,
			},
			args: args{
				id: 1,
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
				Ctx:  tt.fields.Ctx,
				Repo: tt.fields.Repo,
			}

			mock.EXPECT().Read(tt.args.id).Return(tt.want, tt.wantErr.err)

			got, err := u.Get(tt.args.id)
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

	wantErrValue := &model.AlreadyExistError{
		ID:        1,
		ModelName: model.ModelNameProgrammingLang,
		Name:      model.CreateProgrammingLangs(1)[0].Name,
	}

	type fields struct {
		Ctx  context.Context
		Repo repository.ProgrammingLangRepository
	}
	type args struct {
		param *model.ProgrammingLang
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
		readWant *model.ProgrammingLang
		wantErr  wantErr
	}{
		{
			name: "同一のProgrammingLang存在しない場合、ProgrammingLangを登録すること",
			fields: fields{
				Ctx:  context.TODO(),
				Repo: mock,
			},
			args: args{
				param: model.CreateProgrammingLangs(1)[0],
			},
			want:     model.CreateProgrammingLangs(1)[0],
			readWant: nil,
			wantErr: wantErr{
				isErr: false,
				err:   nil,
			},
		},
		{
			name: "同一のProgrammingLang存在する場合、nilとエラーを返すこと",
			fields: fields{
				Ctx:  context.TODO(),
				Repo: mock,
			},
			args: args{
				param: model.CreateProgrammingLangs(1)[0],
			},
			want:     nil,
			readWant: model.CreateProgrammingLangs(1)[0],
			wantErr: wantErr{
				isErr: true,
				err:   wantErrValue,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &ProgrammingLangUseCase{
				Ctx:  tt.fields.Ctx,
				Repo: tt.fields.Repo,
			}

			mock.EXPECT().ReadByName(tt.args.param.Name).Return(tt.readWant, tt.wantErr.err)

			if !tt.wantErr.isErr {
				mock.EXPECT().Create(tt.args.param).Return(tt.want, tt.wantErr.err)
			}

			got, err := u.Create(tt.args.param)

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

	wantErrValue := &model.NoSuchDataError{
		ID:        1,
		ModelName: model.ModelNameProgrammingLang,
		Name:      model.CreateProgrammingLangs(1)[0].Name,
	}

	type fields struct {
		Ctx  context.Context
		Repo repository.ProgrammingLangRepository
	}
	type args struct {
		param *model.ProgrammingLang
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
		readWant *model.ProgrammingLang
		wantErr  wantErr
	}{
		{
			name: "同一のProgrammingLang存在する場合、ProgrammingLangを更新すること",
			fields: fields{
				Ctx:  context.TODO(),
				Repo: mock,
			},
			args: args{
				param: model.CreateProgrammingLangs(1)[0],
			},
			want:     model.CreateProgrammingLangs(1)[0],
			readWant: model.CreateProgrammingLangs(1)[0],
			wantErr: wantErr{
				isErr: false,
				err:   nil,
			},
		},
		{
			name: "指定したProgrammingLang存在しない場合、nilとエラーを返すこと",
			fields: fields{
				Ctx:  context.TODO(),
				Repo: mock,
			},
			args: args{
				param: model.CreateProgrammingLangs(1)[0],
			},
			want:     nil,
			readWant: nil,
			wantErr: wantErr{
				isErr: true,
				err:   wantErrValue,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &ProgrammingLangUseCase{
				Ctx:  tt.fields.Ctx,
				Repo: tt.fields.Repo,
			}

			mock.EXPECT().Read(tt.args.param.ID).Return(tt.readWant, tt.wantErr.err)

			if !tt.wantErr.isErr {
				mock.EXPECT().Update(tt.args.param).Return(tt.want, tt.wantErr.err)
			}

			got, err := u.Update(tt.args.param)
			if (err != nil) != tt.wantErr.isErr {
				t.Errorf("ProgrammingLangUseCase.Update() error = %v, wantErr %v", err, tt.wantErr)
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
		id int
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
				Ctx:  context.TODO(),
				Repo: mock,
			},
			args: args{
				id: 1,
			},
			wantErr: wantErr{
				isErr: false,
				err:   nil,
			},
		},
		{
			name: "IDで指定したProgrammingLang存在しない場合、nilとエラーを返すこと",
			fields: fields{
				Ctx:  context.TODO(),
				Repo: mock,
			},
			args: args{
				id: 2,
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
				Ctx:  tt.fields.Ctx,
				Repo: tt.fields.Repo,
			}

			mock.EXPECT().Delete(tt.args.id).Return(tt.wantErr.err)

			if err := u.Delete(tt.args.id); (err != nil) != tt.wantErr.isErr {
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
