package handler

import (
	"context"
	"strconv"
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

type DoctorHandler struct {
	doctorUsecase usecase.IDoctorUsecase
	validator     *validator.Validate
}

func NewDoctorHandler(doctorUsecase usecase.IDoctorUsecase, validator *validator.Validate) bootstrap.Handler {
	return &DoctorHandler{
		doctorUsecase: doctorUsecase,
		validator:     validator,
	}
}

func (h *DoctorHandler) InitRoutes(app router.Party) {
	group := app.Party("/doctors")
	group.Use(middleware.Authentication())
	group.Get("", h.GetAllDoctor)
	group.Get("/:doctorID", h.GetDoctorByID)
	group.Post("/:doctorScheduleID", h.BookDoctor)

}

func (h *DoctorHandler) GetAllDoctor(ctx iris.Context) {
	c, cancel := context.WithTimeout(ctx.Clone(), 5*time.Second)
	defer cancel()

	var filter dto.GetDoctorsFilter
	if err := util.ParseGetDoctorsFilter(ctx, &filter); err != nil {
		if errx, ok := err.(*errx.Errx); ok {
			ctx.StopWithJSON(errx.Code, iris.Map{
				"message": errx.Message,
				"error":   errx.Err.Error(),
			})
			return
		}
	}

	doctors, err := h.doctorUsecase.GetAllDoctor(c, filter)
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
		"message": "Doctors has been fetched",
		"doctors": doctors,
	})

}

func (h *DoctorHandler) GetDoctorByID(ctx iris.Context) {
	c, cancel := context.WithTimeout(ctx.Clone(), 5*time.Second)
	defer cancel()

	doctorIDStr := ctx.Params().Get("doctorID")
	doctorID, err := strconv.ParseUint(doctorIDStr, 10, 64)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"message": "Invalid doctor ID",
			"error":   err.Error(),
		})
		return
	}

	article, err := h.doctorUsecase.GetDoctorByID(c, doctorID)
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
		"message": "Doctor has been fetched",
		"doctor":  article,
	})

}

func (h *DoctorHandler) BookDoctor(ctx iris.Context) {
	c, cancel := context.WithTimeout(ctx.Clone(), 5*time.Second)
	defer cancel()

	doctorScheduleIDStr := ctx.Params().Get("doctorScheduleID")
	doctorScheduleID, err := strconv.ParseUint(doctorScheduleIDStr, 10, 64)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"message": "Invalid doctor schedule ID",
			"error":   err.Error(),
		})
		return
	}

	userID, ok := ctx.Values().Get("id").(string)
	if !ok {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"message": "User ID context not found",
			"error":   "User ID context not found",
		})
		return
	}

	res, err := h.doctorUsecase.CreateConsultation(c, doctorScheduleID, userID)
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
		"message":  "Consultation has been created",
		"response": res,
	})
}
