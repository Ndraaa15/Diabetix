package bootstrap

import (
	"fmt"
	"time"

	"github.com/Ndraaa15/diabetix-server/pkg/env"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
	"go.uber.org/dig"
	"go.uber.org/zap"
	"gorm.io/gorm"
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
	Srv *iris.Application
	Zap *zap.Logger
	Env *env.Env
	// Cron     *cron.Cron
	Db       *gorm.DB
	Handlers []Handler `group:"handlers"`
}

func Run(params BootstrapParams) {
	// ctx := context.Background()

	params.Srv.Use(func(ctx iris.Context) {
		start := time.Now()

		params.Zap.Sugar().Infof("request: method=%s, path=%s, ip=%s, user-agent=%s",
			ctx.Method(), ctx.Path(), ctx.RemoteAddr(), ctx.GetHeader("User-Agent"))

		ctx.Next()

		duration := time.Since(start)
		params.Zap.Sugar().Infof("response: status=%d, duration=%s, path=%s",
			ctx.GetStatusCode(), duration, ctx.Path())
	})

	params.Srv.Get("/", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello, world🌍"})
	})

	group := params.Srv.Party("/api/v1")
	group.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("Path: %s", ctx.Path())
		ctx.Next()
	})

	for _, handler := range params.Handlers {
		handler.InitRoutes(group)
	}

	// 	params.Cron.AddFunc("@daily",
	// 	err := cronx.CreateTracker(ctx, params.Db)
	// 	err = cronx.GenerateMission(ctx, params.Db)
	// )

	params.Zap.Sugar().Info("Server is running on port 8080...")
	if err := params.Srv.Run(iris.Addr(fmt.Sprintf("%s:%s", params.Env.AppAddr, params.Env.AppPort))); err != nil {
		params.Zap.Sugar().Fatal("Server failed to start:", err)
	}
}
