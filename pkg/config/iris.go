package config

import (
	"github.com/Ndraaa15/diabetix-server/pkg/env"
	"github.com/kataras/iris/v12"
)

func NewIris(env *env.Env) *iris.Application {
	app := iris.New()
	app.Configure(iris.WithConfiguration(iris.Configuration{
		SocketSharding: true,
	}))

	app.Logger().SetFormat("json")

	return app
}
