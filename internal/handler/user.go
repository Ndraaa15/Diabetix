package handler

import (
	"context"
	"time"

	"github.com/Ndraaa15/diabetix-server/cmd/bootstrap"
	"github.com/Ndraaa15/diabetix-server/internal/dto"
	"github.com/Ndraaa15/diabetix-server/internal/middleware"
	"github.com/Ndraaa15/diabetix-server/internal/usecase"
	"github.com/Ndraaa15/diabetix-server/pkg/util"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
)

type UserHandler struct {
	userUsecase usecase.IUserUsecase
}

func NewUserHandler(userUsecase usecase.IUserUsecase) bootstrap.Handler {
	return &UserHandler{
		userUsecase: userUsecase,
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
		if validationErr, ok := err.(validator.ValidationErrors); ok {
			ctx.StopWithJSON(iris.StatusBadRequest, util.WrapValidationErrors(validationErr))
			return
		}

		ctx.StopWithJSON(iris.StatusBadRequest, err)
		return
	}

	err := h.userUsecase.CreatePersonalization(c, req)
	if err != nil {
		ctx.StopWithJSON(iris.StatusInternalServerError, err)
		return
	}

	ctx.StopWithJSON(iris.StatusCreated, iris.Map{
		"message": "User has been filled personalization",
		"id":      req.UserID,
	})
}

func (h *UserHandler) GetProfile(ctx iris.Context) {
	c, cancel := context.WithTimeout(ctx.Clone(), 5*time.Second)
	defer cancel()

	id, ok := ctx.Values().Get("id").(string)
	if !ok {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"message": "User ID context not found",
		})
	}

	profiles, err := h.userUsecase.GetProfile(c, id)
	if err != nil {
		ctx.StopWithJSON(iris.StatusInternalServerError, err)
		return
	}

	ctx.StopWithJSON(iris.StatusCreated, iris.Map{
		"message":  "Profiles has been fetched",
		"profiles": profiles,
	})
}

func (h *UserHandler) GetHomepage(ctx iris.Context) {
	c, cancel := context.WithTimeout(ctx.Clone(), 5*time.Second)
	defer cancel()

	id, ok := ctx.Values().Get("id").(string)
	if !ok {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"message": "User ID context not found",
		})
	}

	homePageData, err := h.userUsecase.GetDataForHomePage(c, id)
	if err != nil {
		ctx.StopWithJSON(iris.StatusInternalServerError, err)
		return
	}

	ctx.StopWithJSON(iris.StatusOK, iris.Map{
		"message": "Homepage has been fetched",
		"data":    homePageData,
	})
}
