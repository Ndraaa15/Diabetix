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

type TrackerHandler struct {
	trackerUsecase usecase.ITrackerUsecase
	validator      *validator.Validate
}

func NewTrackerHandler(trackerUsecase usecase.ITrackerUsecase, validator *validator.Validate) bootstrap.Handler {
	return &TrackerHandler{
		trackerUsecase: trackerUsecase,
		validator:      validator,
	}
}

func (h *TrackerHandler) InitRoutes(app router.Party) {
	group := app.Party("/trackers")
	group.Use(middleware.Authentication())
	group.Post("/predict", h.PredictFood)
	group.Post("/add", h.AddFood)
	group.Get("", h.GetAllTracker)
}

func (h *TrackerHandler) PredictFood(ctx iris.Context) {
	c, cancel := context.WithTimeout(ctx.Clone(), 60*time.Second)
	defer cancel()

	userID, ok := ctx.Values().Get("id").(string)
	if !ok {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"message": "Failed to get user ID from context",
			"error":   "Failed to get user ID from context",
		})
		return
	}

	file, fileHeader, err := ctx.FormFile("foodImage")
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	defer file.Close()

	result, err := h.trackerUsecase.PredictFood(c, fileHeader, file, userID)
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
		"message": "Food has been predicted",
		"result":  result,
	})
}

func (h *TrackerHandler) AddFood(ctx iris.Context) {
	c, cancel := context.WithTimeout(ctx.Clone(), 60*time.Second)
	defer cancel()

	userID, ok := ctx.Values().Get("id").(string)
	if !ok {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"message": "Failed to get user ID from context",
			"error":   "Failed to get user ID from context",
		})
		return
	}

	var req dto.CreateTrackerDetailRequest
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

	err := h.trackerUsecase.AddFood(c, req, userID)
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
		"message": "Food has been added",
	})
}

func (h *TrackerHandler) GetAllTracker(ctx iris.Context) {
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

	listTracker, err := h.trackerUsecase.GetAllTracker(c, userID)
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
		"message": "Success retrieve all tracker",
		"result":  listTracker,
	})
}
