package handler

import (
	"context"
	"time"

	"github.com/Ndraaa15/diabetix-server/cmd/bootstrap"
	"github.com/Ndraaa15/diabetix-server/internal/middleware"
	"github.com/Ndraaa15/diabetix-server/internal/usecase"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
)

type FileUploadHandler struct {
	fileUploadUsecase usecase.IFileUploadUsecase
}

func NewFileUploadHandler(fileUploadUsecase usecase.IFileUploadUsecase) bootstrap.Handler {
	return &FileUploadHandler{
		fileUploadUsecase: fileUploadUsecase,
	}
}

func (h *FileUploadHandler) InitRoutes(app router.Party) {
	group := app.Party("/uploads")
	group.Use(middleware.Authentication())
	group.Post("/files", h.UploadFile)
}

func (h *FileUploadHandler) UploadFile(ctx iris.Context) {
	c, cancel := context.WithTimeout(ctx.Clone(), 60*time.Second)
	defer cancel()

	file, _, err := ctx.FormFile("file")
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, err)
		return
	}

	defer file.Close()

	result, err := h.fileUploadUsecase.UploadFile(c, file)
	if err != nil {
		ctx.StopWithJSON(iris.StatusInternalServerError, err)
		return
	}

	ctx.StopWithJSON(iris.StatusOK, iris.Map{
		"message": "File has been uploaded",
		"url":     result,
	})
}
