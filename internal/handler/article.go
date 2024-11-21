package handler

import (
	"context"
	"strconv"
	"time"

	"github.com/Ndraaa15/diabetix-server/cmd/bootstrap"
	"github.com/Ndraaa15/diabetix-server/internal/dto"
	"github.com/Ndraaa15/diabetix-server/internal/middleware"
	"github.com/Ndraaa15/diabetix-server/internal/usecase"
	"github.com/Ndraaa15/diabetix-server/pkg/errx"
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
	group.Post("/:articleID/likes", h.CreateLikes)
	group.Delete("/:articleID/likes", h.DeleteLikes)
	group.Get("/:articleID", h.GetArticleByID)
}

func (h *ArticleHandler) GetArticles(ctx iris.Context) {
	c, cancel := context.WithTimeout(ctx.Clone(), 5*time.Second)
	defer cancel()

	var filter dto.GetArticlesFilter
	if err := util.ParseGetArticlesFilter(ctx, &filter); err != nil {
		if errx, ok := err.(*errx.Errx); ok {
			ctx.StopWithJSON(errx.Code, iris.Map{
				"message": errx.Message,
				"error":   errx.Err.Error(),
			})
			return
		}
	}

	articles, err := h.articleUsecase.GetArticles(c, filter)
	if err != nil {
		if errx, ok := err.(*errx.Errx); ok {
			ctx.StopWithJSON(errx.Code, iris.Map{
				"message": errx.Message,
				"error":   errx.Err.Error(),
			})
			return
		}
	}

	ctx.StopWithJSON(iris.StatusCreated, iris.Map{
		"message":  "Articles has been fetched",
		"articles": articles,
	})
}

func (h *ArticleHandler) CreateLikes(ctx iris.Context) {
	c, cancel := context.WithTimeout(ctx.Clone(), 5*time.Second)
	defer cancel()

	userID, ok := ctx.Values().Get("id").(string)
	if !ok {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"message": "Failed to get user ID from context",
			"error":   "Failed to get user ID from context",
		})
		return
	}

	articleIDStr := ctx.Params().Get("articleID")
	articleID, err := strconv.ParseUint(articleIDStr, 10, 64)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"message": "Invalid article ID",
			"error":   err.Error(),
		})
		return
	}

	err = h.articleUsecase.CreateLikes(c, userID, articleID)
	if err != nil {
		if errx, ok := err.(*errx.Errx); ok {
			ctx.StopWithJSON(errx.Code, iris.Map{
				"message": errx.Message,
				"error":   errx.Err.Error(),
			})
			return
		}
	}

	ctx.StopWithJSON(iris.StatusCreated, iris.Map{
		"message": "Like has been created",
	})
}

func (h *ArticleHandler) DeleteLikes(ctx iris.Context) {
	c, cancel := context.WithTimeout(ctx.Clone(), 5*time.Second)
	defer cancel()

	userID, ok := ctx.Values().Get("id").(string)
	if !ok {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"message": "Failed to get user ID from context",
			"error":   "Failed to get user ID from context",
		})
		return
	}

	articleIDStr := ctx.Params().Get("articleID")
	articleID, err := strconv.ParseUint(articleIDStr, 10, 64)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"message": "Invalid article ID",
			"error":   err.Error(),
		})
		return
	}

	err = h.articleUsecase.DeleteLikes(c, userID, articleID)
	if err != nil {
		if errx, ok := err.(*errx.Errx); ok {
			ctx.StopWithJSON(errx.Code, iris.Map{
				"message": errx.Message,
				"error":   errx.Err.Error(),
			})
			return
		}
	}

	ctx.StopWithJSON(iris.StatusCreated, iris.Map{
		"message": "Like has been deleted",
	})
}

func (h *ArticleHandler) GetArticleByID(ctx iris.Context) {
	c, cancel := context.WithTimeout(ctx.Clone(), 5*time.Second)
	defer cancel()

	articleIDStr := ctx.Params().Get("articleID")
	articleID, err := strconv.ParseUint(articleIDStr, 10, 64)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"message": "Invalid article ID",
			"error":   err.Error(),
		})
		return
	}

	article, err := h.articleUsecase.GetArticleByID(c, articleID)
	if err != nil {
		if errx, ok := err.(*errx.Errx); ok {
			ctx.StopWithJSON(errx.Code, iris.Map{
				"message": errx.Message,
				"error":   errx.Err.Error(),
			})
			return
		}
	}

	ctx.StopWithJSON(iris.StatusCreated, iris.Map{
		"message": "Article has been fetched",
		"article": article,
	})
}
