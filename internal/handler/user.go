package handler

import (
	"context"
	"time"

	"github.com/Ndraaa15/diabetix-server/cmd/bootstrap"
	"github.com/Ndraaa15/diabetix-server/internal/dto"
	"github.com/Ndraaa15/diabetix-server/internal/middleware"
	"github.com/Ndraaa15/diabetix-server/internal/usecase"
	"github.com/Ndraaa15/diabetix-server/pkg/errx"
	"github.com/Ndraaa15/diabetix-server/pkg/util"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
)

type UserHandler struct {
	userUsecase usecase.IUserUsecase
	validator   *validator.Validate
}

func NewUserHandler(userUsecase usecase.IUserUsecase, validator *validator.Validate) bootstrap.Handler {
	return &UserHandler{
		userUsecase: userUsecase,
		validator:   validator,
	}
}

func (h *UserHandler) InitRoutes(app router.Party) {
	group := app.Party("/users")
	group.Use(middleware.Authentication())
	group.Get("/profile", h.GetProfile)
	group.Get("/homepage", h.GetHomepage)

	groupPublic := app.Party("/users")
	groupPublic.Post("/personalization", h.CreatePersonalization)
}

func (h *UserHandler) CreatePersonalization(ctx iris.Context) {
	c, cancel := context.WithTimeout(ctx.Clone(), 5*time.Second)
	defer cancel()

	var req dto.CreatePersonalizationRequest
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	if err := h.validator.Struct(req); err != nil {
		if valErr, ok := err.(validator.ValidationErrors); ok {
			ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
				"message": "Invalid request body",
				"error":   util.HandleValidationErrors(valErr),
			})
			return
		}
	}

	err := h.userUsecase.CreatePersonalization(c, req)
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
		"message": "User has been filled personalization",
		"id":      req.UserID,
	})
}

func (h *UserHandler) GetProfile(ctx iris.Context) {
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

	profiles, err := h.userUsecase.GetProfile(c, userID)
	if err != nil {
		if errx, ok := err.(*errx.Errx); ok {
			ctx.StopWithJSON(errx.Code, iris.Map{
				"message": errx.Message,
				"error":   errx.Err.Error(),
			})
			return
		}
	}

	ctx.StopWithJSON(iris.StatusOK, iris.Map{
		"message":  "Profiles has been fetched",
		"profiles": profiles,
	})
}

func (h *UserHandler) GetHomepage(ctx iris.Context) {
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

	homePageData, err := h.userUsecase.GetDataForHomePage(c, userID)
	if err != nil {
		if errx, ok := err.(*errx.Errx); ok {
			ctx.StopWithJSON(errx.Code, iris.Map{
				"message": errx.Message,
				"error":   errx.Err.Error(),
			})
			return
		}
	}

	ctx.StopWithJSON(iris.StatusOK, iris.Map{
		"message":  "Homepage data has been fetched",
		"response": homePageData,
	})
}
