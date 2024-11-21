package handler

import (
	"context"
	"strconv"
	"time"

	"github.com/Ndraaa15/diabetix-server/cmd/bootstrap"
	"github.com/Ndraaa15/diabetix-server/internal/middleware"
	"github.com/Ndraaa15/diabetix-server/internal/usecase"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
)

type MissionHandler struct {
	missionUsecase usecase.IMissionUsecase
}

func NewMissionHandler(missionUsecase usecase.IMissionUsecase) bootstrap.Handler {
	return &MissionHandler{
		missionUsecase: missionUsecase,
	}
}

func (h *MissionHandler) InitRoutes(app router.Party) {
	group := app.Party("/missions")
	group.Use(middleware.Authentication())
	group.Get("", h.GetAllUserMission)
	group.Patch("/:missionID/accepted", h.AcceptMissionHandler)

}

func (h *MissionHandler) GetAllUserMission(ctx iris.Context) {
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

	mission, err := h.missionUsecase.GetAllUserMission(c, userID)
	if err != nil {
		ctx.StopWithJSON(iris.StatusInternalServerError, err)
		return
	}

	ctx.StopWithJSON(iris.StatusOK, iris.Map{
		"message":  "Mission has been fetched",
		"missions": mission,
	})
}

func (h *MissionHandler) AcceptMissionHandler(ctx iris.Context) {
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

	missionIDStr := ctx.Params().Get("missionID")
	missionID, err := strconv.ParseUint(missionIDStr, 10, 64)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"message": "Mission ID not valid",
			"error":   err.Error(),
		})
		return
	}

	err = h.missionUsecase.UpdateUserMission(c, userID, missionID)
	if err != nil {
		ctx.StopWithJSON(iris.StatusInternalServerError, err)
		return
	}

	ctx.StopWithJSON(iris.StatusOK, iris.Map{
		"message": "Challenge has been accepted",
	})
}
