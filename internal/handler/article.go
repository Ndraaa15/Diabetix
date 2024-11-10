package handler

import (
	"context"
	"time"

	"github.com/Ndraaa15/diabetix-server/cmd/bootstrap"
	"github.com/Ndraaa15/diabetix-server/internal/dto"
	"github.com/Ndraaa15/diabetix-server/internal/middleware"
	"github.com/Ndraaa15/diabetix-server/internal/usecase"
	"github.com/Ndraaa15/diabetix-server/pkg/util"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
)

type ArticleHandler struct {
	articleUsecase usecase.IArticleUsecase
}

func NewArticleHandler(articleUsecase usecase.IArticleUsecase) bootstrap.Handler {
	return &ArticleHandler{
		articleUsecase: articleUsecase,
	}
}

func (h *ArticleHandler) InitRoutes(app router.Party) {
	group := app.Party("/articles")
	group.Use(middleware.Authentication())
	group.Get("", h.GetArticles)
}

func (h *ArticleHandler) GetArticles(ctx iris.Context) {
	c, cancel := context.WithTimeout(ctx.Clone(), 5*time.Second)
	defer cancel()

	var filter dto.GetArticlesFilter
	if err := util.ParseGetArticlesFilter(ctx, &filter); err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"message": err.Error(),
		})
		return
	}

	articles, err := h.articleUsecase.GetArticles(c, filter)
	if err != nil {
		ctx.StopWithJSON(iris.StatusInternalServerError, iris.Map{
			"message": err.Error(),
		})
		return
	}

	ctx.StopWithJSON(iris.StatusCreated, iris.Map{
		"message":  "Articles has been fetched",
		"articles": articles,
	})
}
