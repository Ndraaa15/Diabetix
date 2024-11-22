package util

import (
	"crypto/rand"
	"errors"
	"math/big"
	"time"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"github.com/Ndraaa15/diabetix-server/internal/dto"
	"github.com/Ndraaa15/diabetix-server/pkg/errx"
	"github.com/kataras/iris/v12"
)

const numberCharset = "0123456789"

func GenerateCode(length int) (string, error) {
	code := make([]byte, length)
	for i := range code {
		n, err := rand.Int(rand.Reader, big.NewInt(9))
		if err != nil {
			return "", errx.New().WithCode(iris.StatusInternalServerError).WithMessage("Failed to generate code").WithError(err)
		}

		code[i] = numberCharset[n.Int64()]
	}
	return string(code), nil
}

func ParseGetArticlesFilter(ctx iris.Context, filter *dto.GetArticlesFilter) error {
	if keywordStr := ctx.URLParam("keyword"); keywordStr != "" {
		filter.Keyword = keywordStr
	}

	userID, ok := ctx.Values().Get("id").(string)
	if !ok {
		return errx.New().
			WithCode(iris.StatusBadRequest).
			WithMessage("User ID context not found").
			WithError(errors.New("User ID context not found"))
	}

	filter.UserID = userID

	return nil
}

func ParseGetReportsFilter(ctx iris.Context, filter *dto.GetReportsFilter) error {
	if dateStr := ctx.URLParam("date"); dateStr != "" {
		date, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			return errx.New().WithCode(iris.StatusBadRequest).WithMessage("Invalid date format").WithError(err)
		}

		filter.Date = date
	}

	return nil
}

func ParseGetDoctorsFilter(ctx iris.Context, filter *dto.GetDoctorsFilter) error {
	if keywordStr := ctx.URLParam("keyword"); keywordStr != "" {
		filter.Keyword = keywordStr
	}

	return nil
}

func GetCurrentDate() time.Time {
	return time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)
}

func GenerateRandomMission(mission domain.Mission, total int) []domain.Mission {
	missions := make([]domain.Mission, total)
	for i := 0; i < total; i++ {
		missions[i] = mission
	}
	return missions
}

func GenerateRandomString(length int) string {
	code := make([]byte, length)
	for i := range code {
		n, err := rand.Int(rand.Reader, big.NewInt(9))
		if err != nil {
			return ""
		}

		code[i] = numberCharset[n.Int64()]
	}
	return string(code)
}
