package handler

import (
	"github.com/Ndraaa15/diabetix-server/cmd/bootstrap"
	"github.com/Ndraaa15/diabetix-server/internal/usecase"
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
	group.Get("", h.GetAllDoctor)
	group.Get("/:doctorID", h.GetDoctorByID)
	group.Post("/:doctorID", h.BookDoctor)

}

func (h *DoctorHandler) GetAllDoctor(ctx iris.Context) {

}

func (h *DoctorHandler) GetDoctorByID(ctx iris.Context) {

}

func (h *DoctorHandler) BookDoctor(ctx iris.Context) {

}
