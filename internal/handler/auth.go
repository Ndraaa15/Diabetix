package handler

import (
	"context"
	"time"

	"github.com/Ndraaa15/diabetix-server/internal/dto"
	"github.com/Ndraaa15/diabetix-server/internal/usecase"
	"github.com/Ndraaa15/diabetix-server/pkg/util"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
)

type AuthHandler struct {
	authUsecase usecase.IAuthUsecase
}

func NewAuthHandler(authUsecase usecase.IAuthUsecase) *AuthHandler {
	return &AuthHandler{
		authUsecase: authUsecase,
	}
}

func (h *AuthHandler) InitRoutes(app router.Party) {
	group := app.Party("/auth")
	group.Post("/register", h.Register)
	group.Post("/verify", h.Verify)
	group.Post("/login", h.Login)
}

func (h *AuthHandler) Register(ctx iris.Context) {
	c, cancel := context.WithTimeout(ctx.Clone(), 5*time.Second)
	defer cancel()

	var req dto.RegisterRequest
	if err := ctx.ReadJSON(&req); err != nil {
		if validationErr, ok := err.(validator.ValidationErrors); ok {
			ctx.StopWithJSON(iris.StatusBadRequest, util.WrapValidationErrors(validationErr))
			return
		}

		ctx.StopWithJSON(iris.StatusBadRequest, err)
		return
	}

	id, err := h.authUsecase.Register(c, req)
	if err != nil {
		ctx.StopWithJSON(iris.StatusInternalServerError, err)
		return
	}

	ctx.StopWithJSON(iris.StatusCreated, iris.Map{
		"message": "User has been registered",
		"id":      id,
	})
}

func (h *AuthHandler) Verify(ctx iris.Context) {
	c, cancel := context.WithTimeout(ctx.Clone(), 5*time.Second)
	defer cancel()

	var req dto.VerificationRequest
	if err := ctx.ReadJSON(&req); err != nil {
		if validationErr, ok := err.(validator.ValidationErrors); ok {
			ctx.StopWithJSON(iris.StatusBadRequest, util.WrapValidationErrors(validationErr))
			return
		}

		ctx.StopWithJSON(iris.StatusBadRequest, err)
		return
	}

	err := h.authUsecase.Verify(c, req)
	if err != nil {
		ctx.StopWithJSON(iris.StatusInternalServerError, err)
		return
	}

	ctx.StopWithJSON(iris.StatusOK, iris.Map{
		"message": "User has been verified",
		"id":      req.ID,
	})

}

func (h *AuthHandler) Login(ctx iris.Context) {
	c, cancel := context.WithTimeout(ctx.Clone(), 5*time.Second)
	defer cancel()

	var req dto.LoginRequest
	if err := ctx.ReadJSON(&req); err != nil {
		if validationErr, ok := err.(validator.ValidationErrors); ok {
			ctx.StopWithJSON(iris.StatusBadRequest, util.WrapValidationErrors(validationErr))
			return
		}

		ctx.StopWithJSON(iris.StatusBadRequest, err)
		return
	}

	res, err := h.authUsecase.Login(c, req)
	if err != nil {
		ctx.StopWithJSON(iris.StatusInternalServerError, err)
		return
	}

	ctx.StopWithJSON(iris.StatusOK, iris.Map{
		"message":  "User has been logged in",
		"response": res,
	})
}
