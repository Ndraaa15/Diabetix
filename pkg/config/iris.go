package config

import (
	"github.com/Ndraaa15/diabetix-server/pkg/env"
	"github.com/kataras/iris/v12"
)

func NewIris(env *env.Env) *iris.Application {
	app := iris.New().SetName(env.AppName)
	app.Configure(iris.WithConfiguration(iris.Configuration{
		SocketSharding: true,
	}))

	return app
}
