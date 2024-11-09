package handler

import (
	"context"
	"time"

	"github.com/Ndraaa15/diabetix-server/cmd/bootstrap"
	"github.com/Ndraaa15/diabetix-server/internal/dto"
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
	group := app.Party("/user")
	group.Post("/personalization", h.CreatePersonalizaation)
}

func (h *UserHandler) CreatePersonalizaation(ctx iris.Context) {
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
