package config

import (
	"github.com/Ndraaa15/diabetix-server/pkg/env"
	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
)

func NewIris(env *env.Env) *iris.Application {
	app := iris.New().SetName(env.AppName)
	app.Configure(iris.WithConfiguration(iris.Configuration{
		SocketSharding: true,
	}))

	app.Logger().SetFormat("json")

	app.Use(func(ctx iris.Context) {
		zap.L().Sugar().Infof("Request: ip=%s method=%s uri=%s", ctx.RemoteAddr(), ctx.Request().Method, ctx.Request().RequestURI)
		ctx.Next()
	})

	return app
}
