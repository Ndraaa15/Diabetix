package handler

import (
	"context"
	"time"

	"github.com/Ndraaa15/diabetix-server/cmd/bootstrap"
	"github.com/Ndraaa15/diabetix-server/internal/dto"
	"github.com/Ndraaa15/diabetix-server/internal/middleware"
	"github.com/Ndraaa15/diabetix-server/internal/usecase"
	"github.com/Ndraaa15/diabetix-server/pkg/util"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
)

type ReportHandler struct {
	reportUsecase usecase.IReportUsecase
}

func NewReportHandler(reportUsecase usecase.IReportUsecase) bootstrap.Handler {
	return &ReportHandler{
		reportUsecase: reportUsecase,
	}
}

func (h *ReportHandler) InitRoutes(app router.Party) {
	group := app.Party("/reports")
	group.Use(middleware.Authentication())
	group.Get("", h.GetAllReport)
}

func (h *ReportHandler) GetAllReport(ctx iris.Context) {
	c, cancel := context.WithTimeout(ctx.Clone(), 5*time.Second)
	defer cancel()

	id, ok := ctx.Values().Get("id").(string)
	if !ok {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"message": "User ID context not found",
		})
	}

	var filter dto.GetReportsFilter
	err := util.ParseGetReportsFilter(ctx, &filter)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"message": err.Error(),
		})
		return
	}

	result, err := h.reportUsecase.GetAllReport(c, id, filter)
	if err != nil {
		ctx.StopWithJSON(iris.StatusInternalServerError, err)
		return
	}

	ctx.StopWithJSON(iris.StatusOK, iris.Map{
		"message": "Reports has been fetched",
		"reports": result,
	})
}
