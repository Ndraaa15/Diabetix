package handler

import (
	"context"
	"time"

	"github.com/Ndraaa15/diabetix-server/cmd/bootstrap"
	"github.com/Ndraaa15/diabetix-server/internal/dto"
	"github.com/Ndraaa15/diabetix-server/internal/middleware"
	"github.com/Ndraaa15/diabetix-server/internal/usecase"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
)

type TrackerHandler struct {
	trackerUsecase usecase.ITrackerUsecase
}

func NewTrackerHandler(trackerUsecase usecase.ITrackerUsecase) bootstrap.Handler {
	return &TrackerHandler{
		trackerUsecase: trackerUsecase,
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
	c, cancel := context.WithTimeout(ctx.Clone(), 5*time.Second)
	defer cancel()

	id, ok := ctx.Values().Get("id").(string)
	if !ok {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"message": "User ID context not found",
		})
	}

	file, fileHeader, err := ctx.FormFile("foodImage")
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, err)
		return
	}

	defer file.Close()

	result, err := h.trackerUsecase.PredictFood(c, fileHeader, file, id)
	if err != nil {
		ctx.StopWithJSON(iris.StatusInternalServerError, err)
		return
	}

	ctx.StopWithJSON(iris.StatusOK, iris.Map{
		"message": "Food has been predicted",
		"result":  result,
	})
}

func (h *TrackerHandler) AddFood(ctx iris.Context) {
	c, cancel := context.WithTimeout(ctx.Clone(), 5*time.Second)
	defer cancel()

	var req dto.CreateTrackerDetailRequest
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, err)
		return
	}

	id, ok := ctx.Values().Get("id").(string)
	if !ok {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"message": "User ID context not found",
		})
	}

	err := h.trackerUsecase.AddFood(c, req, id)
	if err != nil {
		ctx.StopWithJSON(iris.StatusInternalServerError, err)
		return
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
			"message": "User ID not found",
		})
		return
	}

	listTracker, err := h.trackerUsecase.GetAllTracker(c, userID)
	if err != nil {
		ctx.StopWithJSON(iris.StatusInternalServerError, err)
		return
	}

	ctx.StopWithJSON(iris.StatusOK, listTracker)
}
