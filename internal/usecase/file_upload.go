package usecase

import (
	"context"
	"mime/multipart"

	"github.com/Ndraaa15/diabetix-server/pkg/cloudinary"
)

type IFileUploadUsecase interface {
	UploadFile(ctx context.Context, file multipart.File) (string, error)
}

type FileUploadUsecase struct {
	cloudinary *cloudinary.Cloudinary
}

func NewFileUploadUsecase(cloudinary *cloudinary.Cloudinary) IFileUploadUsecase {
	return &FileUploadUsecase{
		cloudinary: cloudinary,
	}
}

func (uc *FileUploadUsecase) UploadFile(ctx context.Context, file multipart.File) (string, error) {
	url, err := uc.cloudinary.UploadFile(ctx, file)
	if err != nil {
		return "", err
	}

	return url, nil
}
