package api

import (
	"strconv"

	"github.com/SekiguchiKai/clean-architecture-with-go/server/domain/model"
	"github.com/gin-gonic/gin"
	"github.com/SekiguchiKai/clean-architecture-with-go/server/util"
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
	var err error
	limit := 20

	limitStr := c.Query(Limit)
	if !util.IsEmpty(limitStr){
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			return -1, &model.InvalidParameterError{
				Parameter: Limit,
				Message:   LimitShouldBeIntErr,
			}
		}
	}

	return limit, nil
}

// ManageLimit は、Limitを制御する。
func ManageLimit(targetLimit, maxLimit, defaultLimit int) int {
	if  maxLimit < targetLimit {
		return defaultLimit
	}
	return targetLimit
}
