package router

import (
	"fmt"

	"github.com/SekiguchiKai/clean-architecture-with-go/server/adapter/api"
	"github.com/SekiguchiKai/clean-architecture-with-go/server/infra/dao/rdb"
	"github.com/SekiguchiKai/clean-architecture-with-go/server/usecase"
	"github.com/gin-gonic/gin"
)

// G は、ginのフレームワークのインスタンス。
var G *gin.Engine

// init は、アプリケーションの初期設定を行う。
func init() {
	g := gin.New()
	apiV1 := g.Group("/v1")

	sqlM := rdb.NewSQLManager()
	langAPI := initProgrammingLang(sqlM)
	langAPI.InitAPI(apiV1)

	G = g
}

// initProgrammingLang は、ProgrammingLangに関する初期設定を行う。
func initProgrammingLang(sqlM rdb.SQLManagerInterface) *api.ProgrammingLangAPI {
	rep := rdb.NewProgrammingLangDAO(sqlM)
	u := usecase.NewProgrammingLangUseCase(rep)
	api := api.NewProgrammingLangAPI(u)
	fmt.Printf("~~api~:%+v\n", api)
	return api
}
