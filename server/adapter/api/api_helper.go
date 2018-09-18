package api

import (
	"strconv"

	"github.com/SekiguchiKai/clean-architecture-with-go/server/domain/model"
	"github.com/gin-gonic/gin"
)

// getID は、URLからIDの値を取得する。
func getID(c *gin.Context) (int, error) {
	id, err := strconv.Atoi(c.Param(ID))
	if err != nil {
		return -1, &model.InvalidParameterError{
			Parameter: ID,
			Message:   IDShouldBeIntErr,
		}
	}

	return id, nil
}

// getLimit は、Query StringからLimitの値を取得する。
func getLimit(c *gin.Context) (int, error) {
	limit, err := strconv.Atoi(c.Query(Limit))
	if err != nil {
		return -1, &model.InvalidParameterError{
			Parameter: Limit,
			Message:   LimitShouldBeIntErr,
		}
	}

	return limit, nil
}
