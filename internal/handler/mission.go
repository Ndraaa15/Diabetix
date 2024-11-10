package handler

import (
	"context"
	"strconv"
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
	group.Get("", h.GetAllMissionUser)
	group.Patch("/:missionID", h.AcceptMissionHandler)

}

func (h *MissionHandler) GetAllMissionUser(ctx iris.Context) {
	c, cancel := context.WithTimeout(ctx.Clone(), 5*time.Second)
	defer cancel()

	id, ok := ctx.Values().Get("id").(string)
	if !ok {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"message": "User ID context not found",
		})
	}

	mission, err := h.missionUsecase.GetAllMissionUser(c, id)
	if err != nil {
		ctx.StopWithJSON(iris.StatusInternalServerError, err)
		return
	}

	ctx.StopWithJSON(iris.StatusCreated, iris.Map{
		"message":  "Mission has been fetched",
		"missions": mission,
	})
}

func (h *MissionHandler) AcceptMissionHandler(ctx iris.Context) {
	c, cancel := context.WithTimeout(ctx.Clone(), 5*time.Second)
	defer cancel()

	id, ok := ctx.Values().Get("id").(string)
	if !ok {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"message": "User ID context not found",
		})
	}

	missionIDStr := ctx.Params().Get("missionID")
	missionID, err := strconv.ParseUint(missionIDStr, 10, 64)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"message": "Mission ID not valid",
		})
		return
	}

	var req dto.UpdateUserMissionRequest
	if err := ctx.ReadJSON(&req); err != nil {
		if validationErr, ok := err.(validator.ValidationErrors); ok {
			ctx.StopWithJSON(iris.StatusBadRequest, util.WrapValidationErrors(validationErr))
			return
		}

		ctx.StopWithJSON(iris.StatusBadRequest, err)
		return
	}

	err = h.missionUsecase.UpdateUserMission(c, id, missionID, req)
	if err != nil {
		ctx.StopWithJSON(iris.StatusInternalServerError, err)
		return
	}

	ctx.StopWithJSON(iris.StatusCreated, iris.Map{
		"message": "Challenge has been accepted",
	})
}
