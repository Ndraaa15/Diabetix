package util

import (
	"crypto/rand"
	"math/big"

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
