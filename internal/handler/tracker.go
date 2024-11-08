package handler

import (
	"github.com/Ndraaa15/diabetix-server/cmd/bootstrap"
	"github.com/Ndraaa15/diabetix-server/internal/middleware"
	"github.com/Ndraaa15/diabetix-server/internal/usecase"
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
	group := app.Party("/tracker")
	group.Use(middleware.Authentication())
	// group.Post("/predict", h.PredictFood)
	// group.Post("/add", h.AddFood)
	// group.Get("/list", h.GetListTracker)

}

// func (h *TrackerHandler) PredictFood(ctx iris.Context) {
// 	c, cancel := context.WithTimeout(ctx.Clone(), 5*time.Second)
// 	defer cancel()

// 	file, fileHeader, err := ctx.FormFile("food_image")
// 	if err != nil {
// 		ctx.StopWithJSON(iris.StatusBadRequest, err)
// 		return
// 	}

// 	defer file.Close()

// 	result, err := h.trackerUsecase.PredictFood(c, file, fileHeader)
// 	if err != nil {
// 		ctx.StopWithJSON(iris.StatusInternalServerError, err)
// 		return
// 	}

// 	ctx.StopWithJSON(iris.StatusOK, iris.Map{
// 		"message": "Food has been predicted",
// 		"result":  result,
// 	})
// }

// func (h *TrackerHandler) AddFood(ctx iris.Context) {
// 	c, cancel := context.WithTimeout(ctx.Clone(), 5*time.Second)
// 	defer cancel()

// 	var req dto.AddFoodRequest
// 	if err := ctx.ReadJSON(&req); err != nil {
// 		ctx.StopWithJSON(iris.StatusBadRequest, err)
// 		return
// 	}

// 	err := h.trackerUsecase.AddFood(c, req)
// 	if err != nil {
// 		ctx.StopWithJSON(iris.StatusInternalServerError, err)
// 		return
// 	}

// 	ctx.StopWithJSON(iris.StatusCreated, iris.Map{
// 		"message": "Food has been added",
// 	})
// }

// func (h *TrackerHandler) GetListTracker(ctx iris.Context) {
// 	c, cancel := context.WithTimeout(ctx.Clone(), 5*time.Second)
// 	defer cancel()

// 	userID, ok := ctx.Values().Get("id").(string)
// 	if !ok {
// 		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
// 			"message": "User ID not found",
// 		})
// 		return
// 	}

// 	listTracker, err := h.trackerUsecase.GetListTracker(c, userID)
// 	if err != nil {
// 		ctx.StopWithJSON(iris.StatusInternalServerError, err)
// 		return
// 	}

// 	ctx.StopWithJSON(iris.StatusOK, listTracker)
// }
