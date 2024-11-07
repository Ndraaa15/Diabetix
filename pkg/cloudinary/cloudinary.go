package cloudinary

import (
	"context"
	"mime/multipart"

	"github.com/Ndraaa15/diabetix-server/pkg/env"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"go.uber.org/zap"
)

type Cloudinary struct {
	cloudinary *cloudinary.Cloudinary
	folder     string
}

func NewCloudinary(env env.Env) *Cloudinary {
	cld, err := cloudinary.NewFromParams(env.CloudinaryName, env.CloudinaryApiKey, env.CloudinaryApiSecret)
	if err != nil {
		zap.S().Fatal(err)
	}
	return &Cloudinary{
		cloudinary: cld,
		folder:     env.CloudinaryFolder,
	}
}

func (c *Cloudinary) UploadFile(ctx context.Context, file multipart.File) (string, error) {
	unique := true
	res, err := c.cloudinary.Upload.Upload(ctx, file, uploader.UploadParams{
		Folder:         c.folder,
		UniqueFilename: unique,
	})

	if err != nil {
		return "", err
	}

	return res.SecureURL, nil
}
