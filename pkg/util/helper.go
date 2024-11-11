package util

import (
	"crypto/rand"
	"math/big"
	"time"

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
