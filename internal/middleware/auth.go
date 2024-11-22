package middleware

import (
	"strings"

	"github.com/Ndraaa15/diabetix-server/pkg/jwt"
	"github.com/kataras/iris/v12"
)

func Authentication() iris.Handler {
	return func(ctx iris.Context) {
		header := ctx.GetHeader("Authorization")

		if header == "" {
			ctx.StatusCode(iris.StatusUnauthorized)
			ctx.JSON(iris.Map{
				"message": "Unauthorized",
			})
			return
		}

		token := strings.SplitN(header, " ", 2)
		if token[0] != "Bearer" {
			ctx.StopWithJSON(iris.StatusUnauthorized, iris.Map{
				"message": "Unauthorized",
			})
			return
		}

		claims, err := jwt.DecodeToken(token[1])
		if err != nil {
			ctx.StopWithJSON(iris.StatusUnauthorized, iris.Map{
				"message": err.Error(),
			})
			return
		}

		ctx.Values().Set("id", claims.ID)
		ctx.Next()
	}
}
