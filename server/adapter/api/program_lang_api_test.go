package api_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"context"

	"github.com/SekiguchiKai/clean-architecture-with-go/server/adapter/api"
	"github.com/SekiguchiKai/clean-architecture-with-go/server/domain/model"
	"github.com/SekiguchiKai/clean-architecture-with-go/server/usecase/input"
	"github.com/SekiguchiKai/clean-architecture-with-go/server/usecase/mock"
	"github.com/SekiguchiKai/clean-architecture-with-go/server/util"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

func TestNewProgrammingLangAPI(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_input.NewMockProgrammingLangInputPort(ctrl)

	type args struct {
		useCase input.ProgrammingLangInputPort
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "適切な引数を与えるとProgrammingLangAPIを返すこと",
			args: args{
				useCase: mock,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := api.NewProgrammingLangAPI(tt.args.useCase); got == nil {
				t.Errorf("NewProgrammingLangAPI() = %v, want nil", got)
			}
		})
	}
}

func TestProgrammingLangAPI_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	u := mock_input.NewMockProgrammingLangInputPort(ctrl)

	langAPI := &api.ProgrammingLangAPI{
		UseCase: u,
	}
	handler := langAPI.List

	paramErr := &model.InvalidParameterError{
		Parameter: api.Limit,
		Message:   api.LimitShouldBeIntErr,
	}

	dbErr := &model.DBError{
		ModelName: model.ModelNameProgrammingLang,
		DBMethod:  model.DBMethodList,
		Detail:    "Test",
	}

	type mock struct {
		ctx      context.Context
		errLimit string
		limit    int
		result   []*model.ProgrammingLang
		err      error
	}

	type want struct {
		code       int
		result     []*model.ProgrammingLang
		errMessage string
	}

	tests := []struct {
		name string
		mock mock
		want want
	}{
		{
			name: "リクエストのクエリパラメータが20で、データが20件以上存在する場合、ステータスコード200と20件のデータを返すこと",
			mock: mock{
				ctx:    context.Background(),
				limit:  20,
				result: model.CreateProgrammingLangs(20),
				err:    nil,
			},
			want: want{
				code:   http.StatusOK,
				result: model.CreateProgrammingLangs(20),
			},
		},
		{
			name: "リクエストのクエリパラメータが文字列の場合、ステータスコード400とエラーメッセージを返すこと",
			mock: mock{
				ctx:      context.Background(),
				errLimit: "test",
				result:   nil,
				err:      paramErr,
			},
			want: want{
				code:       http.StatusBadRequest,
				result:     nil,
				errMessage: paramErr.Error(),
			},
		},
		{
			name: "サーバー側のエラーが発生した場合、ステータスコード500とエラーメッセージを返すこと",
			mock: mock{
				ctx:    context.Background(),
				limit:  20,
				result: nil,
				err:    dbErr,
			},
			want: want{
				code:       http.StatusInternalServerError,
				result:     nil,
				errMessage: dbErr.Error(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gin.New()
			r.GET(api.ProgrammingLangAPIPath, handler)

			url := fmt.Sprintf("%s?%s=%s", api.ProgrammingLangAPIPath, api.Limit, tt.mock.errLimit)
			if util.IsEmpty(tt.mock.errLimit) {
				url = fmt.Sprintf("%s?%s=%d", api.ProgrammingLangAPIPath, api.Limit, tt.mock.limit)
				u.EXPECT().List(tt.mock.ctx, tt.mock.limit).Return(tt.mock.result, tt.mock.err)
			}

			rec := httptest.NewRecorder()
			req, err := http.NewRequest(api.Get, url, nil)
			if err != nil {
				t.Fatal(err)
			}
			r.ServeHTTP(rec, req)

			if tt.want.code == http.StatusOK {
				var got []*model.ProgrammingLang
				err = json.Unmarshal(rec.Body.Bytes(), &got)
				if err != nil {
					t.Fatal(err)
				}

				for i, v := range tt.want.result {
					if !reflect.DeepEqual(got[i], v) {
						t.Errorf("Response Body = %v, want %v", got[i], v)
					}
				}

			} else {
				if util.TrimDoubleQuotes(rec.Body.String()) != tt.want.errMessage {
					t.Errorf("Error Message = %v, want %v", util.TrimDoubleQuotes(rec.Body.String()), tt.want.errMessage)
				}
			}
			if !reflect.DeepEqual(rec.Code, tt.want.code) {
				t.Errorf("Status Code = %v, want %v", rec.Code, tt.want.code)
			}
		})
	}
}

func TestProgrammingLangAPI_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	u := mock_input.NewMockProgrammingLangInputPort(ctrl)

	langAPI := &api.ProgrammingLangAPI{
		UseCase: u,
	}
	handler := langAPI.Get

	noDataErr := &model.NoSuchDataError{
		ID:        1,
		Name:      model.TestName,
		ModelName: model.ModelNameProgrammingLang,
	}

	dbErr := &model.DBError{
		ModelName: model.ModelNameProgrammingLang,
		DBMethod:  model.DBMethodRead,
		Detail:    "Test",
	}

	type mock struct {
		ctx    context.Context
		id     int
		result *model.ProgrammingLang
		err    error
	}

	type want struct {
		code       int
		result     *model.ProgrammingLang
		errMessage string
	}

	tests := []struct {
		name string
		mock mock
		want want
	}{
		{
			name: "リクエストのURLのIDのパラメータが適切な場合、ステータスコード200と1件のデータを返すこと",
			mock: mock{
				ctx:    context.Background(),
				id:     1,
				result: model.CreateProgrammingLangs(1)[0],
				err:    nil,
			},
			want: want{
				code:   http.StatusOK,
				result: model.CreateProgrammingLangs(1)[0],
			},
		},
		{
			name: "リクエストのURLのIDのパラメータと同一のIDを持つデータが存在しない場合、ステータスコード404とエラーメッセージを返すこと",
			mock: mock{
				ctx:    context.Background(),
				id:     100,
				result: nil,
				err:    noDataErr,
			},
			want: want{
				code:       http.StatusNotFound,
				result:     nil,
				errMessage: noDataErr.Error(),
			},
		},
		{
			name: "サーバー側のエラーが発生した場合、ステータスコード500とエラーメッセージを返すこと",
			mock: mock{
				ctx:    context.Background(),
				id:     20,
				result: nil,
				err:    dbErr,
			},
			want: want{
				code:       http.StatusInternalServerError,
				result:     nil,
				errMessage: dbErr.Error(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gin.New()
			r.GET(fmt.Sprintf("%s/:%s", api.ProgrammingLangAPIPath, api.ID), handler)
			u.EXPECT().Get(tt.mock.ctx, tt.mock.id).Return(tt.mock.result, tt.mock.err)

			rec := httptest.NewRecorder()
			url := fmt.Sprintf("%s/%d", api.ProgrammingLangAPIPath, tt.mock.id)
			req, err := http.NewRequest(api.Get, url, nil)
			if err != nil {
				t.Fatal(err)
			}
			r.ServeHTTP(rec, req)

			if tt.want.code == http.StatusOK {
				var got *model.ProgrammingLang
				err = json.Unmarshal(rec.Body.Bytes(), &got)
				if err != nil {
					t.Fatal(err)
				}
				if !reflect.DeepEqual(got, tt.want.result) {
					t.Errorf("Response Body = %v, want %v", got, tt.want.result)
				}

			} else {
				if util.TrimDoubleQuotes(rec.Body.String()) != tt.want.errMessage {
					t.Errorf("Error Message = %v, want %v", util.TrimDoubleQuotes(rec.Body.String()), tt.want.errMessage)
				}
			}
			if !reflect.DeepEqual(rec.Code, tt.want.code) {
				t.Errorf("Status Code = %v, want %v", rec.Code, tt.want.code)
			}
		})
	}
}

func TestProgrammingLangAPI_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	u := mock_input.NewMockProgrammingLangInputPort(ctrl)

	langAPI := &api.ProgrammingLangAPI{
		UseCase: u,
	}
	handler := langAPI.Create

	conflictErr := &model.AlreadyExistError{
		ID:        1,
		Name:      model.TestName,
		ModelName: model.ModelNameProgrammingLang,
	}

	dbErr := &model.DBError{
		ModelName: model.ModelNameProgrammingLang,
		DBMethod:  model.DBMethodRead,
		Detail:    "Test",
	}

	type mock struct {
		ctx    context.Context
		param  *model.ProgrammingLang
		result *model.ProgrammingLang
		err    error
	}

	type want struct {
		code       int
		result     *model.ProgrammingLang
		errMessage string
	}

	tests := []struct {
		name string
		mock mock
		want want
	}{
		{
			name: "リクエストボディのProgrammingLangが適切な場合、ステータスコード200と1件のデータを返すこと",
			mock: mock{
				ctx:    context.Background(),
				param:  model.CreateProgrammingLangs(1)[0],
				result: model.CreateProgrammingLangs(1)[0],
				err:    nil,
			},
			want: want{
				code:   http.StatusOK,
				result: model.CreateProgrammingLangs(1)[0],
			},
		},
		{
			name: "リクエストボディのProgrammingLangが既に存在している場合、ステータスコード409とエラーメッセージを返すこと",
			mock: mock{
				ctx:    context.Background(),
				param:  model.CreateProgrammingLangs(1)[0],
				result: nil,
				err:    conflictErr,
			},
			want: want{
				code:       http.StatusConflict,
				result:     nil,
				errMessage: conflictErr.Error(),
			},
		},
		{
			name: "サーバー側のエラーが発生した場合、ステータスコード500とエラーメッセージを返すこと",
			mock: mock{
				ctx:    context.Background(),
				param:  model.CreateProgrammingLangs(1)[0],
				result: nil,
				err:    dbErr,
			},
			want: want{
				code:       http.StatusInternalServerError,
				result:     nil,
				errMessage: dbErr.Error(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gin.New()
			r.POST(api.ProgrammingLangAPIPath, handler)

			u.EXPECT().Create(tt.mock.ctx, tt.mock.param).Return(tt.mock.result, tt.mock.err)

			rec := httptest.NewRecorder()
			b, err := json.Marshal(tt.mock.param)
			if err != nil {
				t.Fatal(err)
			}
			body := bytes.NewReader(b)

			req, err := http.NewRequest(api.Post, api.ProgrammingLangAPIPath, body)
			if err != nil {
				t.Fatal(err)
			}
			r.ServeHTTP(rec, req)

			if tt.want.code == http.StatusOK {
				var got *model.ProgrammingLang
				err = json.Unmarshal(rec.Body.Bytes(), &got)
				if err != nil {
					t.Fatal(err)
				}
				if !reflect.DeepEqual(got, tt.want.result) {
					t.Errorf("Response Body = %v, want %v", got, tt.want.result)
				}

			} else {
				if util.TrimDoubleQuotes(rec.Body.String()) != tt.want.errMessage {
					t.Errorf("Error Message = %v, want %v", util.TrimDoubleQuotes(rec.Body.String()), tt.want.errMessage)
				}
			}
			if !reflect.DeepEqual(rec.Code, tt.want.code) {
				t.Errorf("Status Code = %v, want %v", rec.Code, tt.want.code)
			}
		})
	}
}

func TestProgrammingLangAPI_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	u := mock_input.NewMockProgrammingLangInputPort(ctrl)

	langAPI := &api.ProgrammingLangAPI{
		UseCase: u,
	}
	handler := langAPI.Update

	noDataErr := &model.NoSuchDataError{
		ID:        1,
		Name:      model.TestName,
		ModelName: model.ModelNameProgrammingLang,
	}

	dbErr := &model.DBError{
		ModelName: model.ModelNameProgrammingLang,
		DBMethod:  model.DBMethodRead,
		Detail:    "Test",
	}

	type mock struct {
		ctx    context.Context
		id     int
		param  *model.ProgrammingLang
		result *model.ProgrammingLang
		err    error
	}

	type want struct {
		code       int
		result     *model.ProgrammingLang
		errMessage string
	}

	type fields struct {
		UseCase input.ProgrammingLangInputPort
	}

	type param struct {
		id int
	}

	tests := []struct {
		name   string
		fields fields
		mock   mock
		want   want
		param  param
	}{
		{
			name: "リクエストボディのProgrammingLangが適切な場合、ステータスコード200と1件のデータを返すこと",
			mock: mock{
				ctx:    context.Background(),
				param:  model.CreateProgrammingLangs(1)[0],
				result: model.CreateProgrammingLangs(1)[0],
				err:    nil,
			},
			want: want{
				code:   http.StatusOK,
				result: model.CreateProgrammingLangs(1)[0],
			},
			param: param{
				id: 1,
			},
		},
		{
			name: "リクエストのURLのIDのパラメータと同一のIDを持つデータが存在しない場合、ステータスコード404とエラーメッセージを返すこと",
			mock: mock{
				ctx:    context.Background(),
				param:  model.CreateProgrammingLangs(1)[0],
				result: nil,
				err:    noDataErr,
			},
			want: want{
				code:       http.StatusNotFound,
				result:     nil,
				errMessage: noDataErr.Error(),
			},
			param: param{
				id: 100,
			},
		},
		{
			name: "サーバー側のエラーが発生した場合、ステータスコード500とエラーメッセージを返すこと",
			mock: mock{
				ctx:    context.Background(),
				param:  model.CreateProgrammingLangs(1)[0],
				result: nil,
				err:    dbErr,
			},
			want: want{
				code:       http.StatusInternalServerError,
				result:     nil,
				errMessage: dbErr.Error(),
			},
			param: param{
				id: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gin.New()
			r.PUT(fmt.Sprintf("%s/:%s", api.ProgrammingLangAPIPath, api.ID), handler)

			u.EXPECT().Update(tt.mock.ctx, tt.mock.param).Return(tt.mock.result, tt.mock.err)

			rec := httptest.NewRecorder()
			b, err := json.Marshal(tt.mock.param)
			if err != nil {
				t.Fatal(err)
			}
			body := bytes.NewReader(b)

			url := fmt.Sprintf("%s/%d", api.ProgrammingLangAPIPath, tt.param.id)
			req, err := http.NewRequest(api.Put, url, body)
			if err != nil {
				t.Fatal(err)
			}
			r.ServeHTTP(rec, req)

			if tt.want.code == http.StatusOK {
				var got *model.ProgrammingLang
				err = json.Unmarshal(rec.Body.Bytes(), &got)
				if err != nil {
					t.Fatal(err)
				}
				if !reflect.DeepEqual(got, tt.want.result) {
					t.Errorf("Response Body = %v, want %v", got, tt.want.result)
				}

			} else {
				if util.TrimDoubleQuotes(rec.Body.String()) != tt.want.errMessage {
					t.Errorf("Error Message = %v, want %v", util.TrimDoubleQuotes(rec.Body.String()), tt.want.errMessage)
				}
			}
			if !reflect.DeepEqual(rec.Code, tt.want.code) {
				t.Errorf("Status Code = %v, want %v", rec.Code, tt.want.code)
			}
		})
	}
}

func TestProgrammingLangAPI_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	u := mock_input.NewMockProgrammingLangInputPort(ctrl)

	langAPI := &api.ProgrammingLangAPI{
		UseCase: u,
	}
	handler := langAPI.Delete

	noDataErr := &model.NoSuchDataError{
		ID:        1,
		Name:      model.TestName,
		ModelName: model.ModelNameProgrammingLang,
	}

	dbErr := &model.DBError{
		ModelName: model.ModelNameProgrammingLang,
		DBMethod:  model.DBMethodRead,
		Detail:    "Test",
	}

	type mock struct {
		ctx context.Context
		id  int
		err error
	}

	type want struct {
		code       int
		result     *model.ProgrammingLang
		errMessage string
	}

	type param struct {
		id int
	}

	tests := []struct {
		name  string
		mock  mock
		want  want
		param param
	}{
		{
			name: "リクエストボディのProgrammingLangが適切な場合、ステータスコード200を返すこと",
			mock: mock{
				ctx: context.Background(),
				id:  1,
				err: nil,
			},
			want: want{
				code:   http.StatusOK,
				result: nil,
			},
			param: param{
				id: 1,
			},
		},
		{
			name: "リクエストのURLのIDのパラメータと同一のIDを持つデータが存在しない場合、ステータスコード404とエラーメッセージを返すこと",
			mock: mock{
				ctx: context.Background(),
				id:  100,
				err: noDataErr,
			},
			want: want{
				code:       http.StatusNotFound,
				result:     nil,
				errMessage: noDataErr.Error(),
			},
			param: param{
				id: 100,
			},
		},
		{
			name: "サーバー側のエラーが発生した場合、ステータスコード500とエラーメッセージを返すこと",
			mock: mock{
				ctx: context.Background(),
				id:  1,
				err: dbErr,
			},
			want: want{
				code:       http.StatusInternalServerError,
				result:     nil,
				errMessage: dbErr.Error(),
			},
			param: param{
				id: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gin.New()
			r.DELETE(fmt.Sprintf("%s/:%s", api.ProgrammingLangAPIPath, api.ID), handler)

			u.EXPECT().Delete(tt.mock.ctx, tt.mock.id).Return(tt.mock.err)

			url := fmt.Sprintf("%s/%d", api.ProgrammingLangAPIPath, tt.param.id)

			req, err := http.NewRequest(api.Delete, url, nil)
			if err != nil {
				t.Fatal(err)
			}

			rec := httptest.NewRecorder()
			r.ServeHTTP(rec, req)

			if tt.want.code == http.StatusOK {
				var got *model.ProgrammingLang
				err = json.Unmarshal(rec.Body.Bytes(), &got)
				if err != nil {
					t.Fatal(err)
				}
				if !reflect.DeepEqual(got, tt.want.result) {
					t.Errorf("Response Body = %v, want %v", got, tt.want.result)
				}

			} else {
				if util.TrimDoubleQuotes(rec.Body.String()) != tt.want.errMessage {
					t.Errorf("Error Message = %v, want %v", util.TrimDoubleQuotes(rec.Body.String()), tt.want.errMessage)
				}
			}
			if !reflect.DeepEqual(rec.Code, tt.want.code) {
				t.Errorf("Status Code = %v, want %v", rec.Code, tt.want.code)
			}
		})
	}
}
