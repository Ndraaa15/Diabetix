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

type BMIHandler struct {
	bmiUsecase usecase.IBMIUsecase
	validator  *validator.Validate
}

func NewBMIHandler(bmiUsecase usecase.IBMIUsecase, validator *validator.Validate) bootstrap.Handler {
	return &BMIHandler{
		bmiUsecase: bmiUsecase,
		validator:  validator,
	}
}

func (h *BMIHandler) InitRoutes(app router.Party) {
	group := app.Party("/bmis")
	group.Use(middleware.Authentication())
	group.Post("", h.CreateBMI)
	group.Get("", h.GetBMI)
}

func (h *BMIHandler) CreateBMI(ctx iris.Context) {
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
	var req dto.CreateBMIRequest
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

	bmi, err := h.bmiUsecase.CreateBMI(c, req, userID)
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
		"message": "BMI has been created",
		"bmi":     bmi,
	})
}

func (h *BMIHandler) GetBMI(ctx iris.Context) {
	c, cancel := context.WithTimeout(ctx.Clone(), 5*time.Second)
	defer cancel()

	id, ok := ctx.Values().Get("id").(string)
	if !ok {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"message": "User ID context not found",
		})
	}

	bmi, err := h.bmiUsecase.GetCurrentBMI(c, id)
	if err != nil {
		ctx.StopWithJSON(iris.StatusInternalServerError, err)
		return
	}

	ctx.StopWithJSON(iris.StatusOK, iris.Map{
		"message": "BMI has been fetched",
		"data":    bmi,
	})
}
