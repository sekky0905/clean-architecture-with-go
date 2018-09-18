package api

import (
	"fmt"
	"net/http"

	"github.com/SekiguchiKai/clean-architecture-with-go/server/domain/model"
	"github.com/SekiguchiKai/clean-architecture-with-go/server/usecase/input"
	"github.com/gin-gonic/gin"
)

// ProgrammingLangAPI は、ProgrammingLangのAPI。
type ProgrammingLangAPI struct {
	UseCase input.ProgrammingLangInputPort
}

// NewProgrammingLangAPI は、ProgrammingLangAPIを生成し、返す。
func NewProgrammingLangAPI(useCase input.ProgrammingLangInputPort) *ProgrammingLangAPI {
	return &ProgrammingLangAPI{
		UseCase: useCase,
	}
}

// InitAPI は、APIを初期設定する。
func (api *ProgrammingLangAPI) InitAPI(g *gin.RouterGroup) {
	g.GET(ProgrammingLangAPIPath, api.List)
	g.GET(fmt.Sprintf("%s/:%s", ProgrammingLangAPIPath, ID), api.Get)
	g.POST(ProgrammingLangAPIPath, api.Create)
	g.PUT(fmt.Sprintf(fmt.Sprintf("%s/:%s", ProgrammingLangAPIPath, ID), api.Update))
	g.DELETE(fmt.Sprintf("%s/:%s", ProgrammingLangAPIPath, ID), api.Delete)
}

// List は、ProgrammingLangの一覧を返す。
func (api *ProgrammingLangAPI) List(c *gin.Context) {
	limit, err := getLimit(c)
	if err != nil {
		he := handleError(err)
		c.JSON(he.code, he.message)
		return
	}

	ctx := c.Request.Context()
	langSlice, err := api.UseCase.List(ctx, limit)
	if err != nil {
		he := handleError(err)
		c.JSON(he.code, he.message)
		return
	}

	c.JSON(http.StatusOK, langSlice)
}

// Get は、ProgrammingLangを取得する。
func (api *ProgrammingLangAPI) Get(c *gin.Context) {
	id, err := getID(c)
	if err != nil {
		he := handleError(err)
		c.JSON(he.code, he.message)
		return
	}

	ctx := c.Request.Context()
	lang, err := api.UseCase.Get(ctx, id)
	if err != nil {
		he := handleError(err)
		c.JSON(he.code, he.message)
		return
	}

	c.JSON(http.StatusOK, lang)
}

// Create は、ProgrammingLangを生成する。
func (api *ProgrammingLangAPI) Create(c *gin.Context) {
	var params *model.ProgrammingLang
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx := c.Request.Context()
	lang, err := api.UseCase.Create(ctx, params)
	if err != nil {
		he := handleError(err)
		c.JSON(he.code, he.message)
		return
	}

	c.JSON(http.StatusOK, lang)
}

// Update は、ProgrammingLangを更新する。
func (api *ProgrammingLangAPI) Update(c *gin.Context) {
	var params *model.ProgrammingLang
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx := c.Request.Context()
	lang, err := api.UseCase.Update(ctx, params)
	if err != nil {
		he := handleError(err)
		c.JSON(he.code, he.message)
		return
	}

	c.JSON(http.StatusOK, lang)
}

// Delete は、ProgrammingLangを削除する。
func (api *ProgrammingLangAPI) Delete(c *gin.Context) {
	id, err := getID(c)
	if err != nil {
		he := handleError(err)
		c.JSON(he.code, he.message)
		return
	}

	ctx := c.Request.Context()
	if err := api.UseCase.Delete(ctx, id); err != nil {
		he := handleError(err)
		c.JSON(he.code, he.message)
		return
	}

	c.JSON(http.StatusOK, nil)
}
