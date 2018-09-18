package api

import (
	"net/http"

	"github.com/SekiguchiKai/clean-architecture-with-go/server/domain/model"
	"github.com/pkg/errors"
)

// エラーの定数。
const (
	OtherErr            = "some error has occurred"
	IDShouldBeIntErr    = "ID Should be int"
	LimitShouldBeIntErr = "Limit Should be int"
)

// handledError はハンドリング後のエラー。
type handledError struct {
	code    int
	message string
}

// handleError は、エラーをハンドリングする。
func handleError(err error) *handledError {
	switch errors.Cause(err).(type) {
	case *model.NoSuchDataError:
		return &handledError{
			code:    http.StatusNotFound,
			message: errors.Cause(err).Error(),
		}
	case *model.RequiredError:
		return &handledError{
			code:    http.StatusBadRequest,
			message: errors.Cause(err).Error(),
		}
	case *model.InvalidPropertyError:
		return &handledError{
			code:    http.StatusBadRequest,
			message: errors.Cause(err).Error(),
		}
	case *model.InvalidParameterError:
		return &handledError{
			code:    http.StatusBadRequest,
			message: errors.Cause(err).Error(),
		}
	case *model.AlreadyExistError:
		return &handledError{
			code:    http.StatusConflict,
			message: errors.Cause(err).Error(),
		}
	case *model.DBError:
		return &handledError{
			code:    http.StatusInternalServerError,
			message: errors.Cause(err).Error(),
		}
	default:
		return &handledError{
			code:    http.StatusInternalServerError,
			message: OtherErr,
		}
	}
}
