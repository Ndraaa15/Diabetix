package bootstrap

import (
	"fmt"

	"github.com/Ndraaa15/diabetix-server/pkg/env"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

type Handler interface {
	InitRoutes(app router.Party)
}

type HandlerOut struct {
	dig.Out
	Handlers []Handler `group:"handlers"`
}

type BootstrapParams struct {
	dig.In
	Srv      *iris.Application
	Zap      *zap.Logger
	Env      *env.Env
	Handlers []Handler `group:"handlers"`
}

func Run(params BootstrapParams) {
	params.Srv.Get("/", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello, world🌍"})
	})

	group := params.Srv.Party("/api/v1")
	for _, handler := range params.Handlers {
		handler.InitRoutes(group)
	}

	params.Zap.Sugar().Info("Server is running on port 8080...")
	if err := params.Srv.Run(iris.Addr(fmt.Sprintf("%s:%s", params.Env.AppAddr, params.Env.AppPort))); err != nil {
		params.Zap.Sugar().Fatal("Server failed to start:", err)
	}
}
