package bootstrap

import (
	"context"
	"fmt"
	"time"

	"github.com/Ndraaa15/diabetix-server/pkg/cronx"
	"github.com/Ndraaa15/diabetix-server/pkg/env"
	"github.com/Ndraaa15/diabetix-server/pkg/gemini"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
	"github.com/robfig/cron/v3"
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
	Srv      *iris.Application
	Zap      *zap.Logger
	Env      *env.Env
	Cron     *cron.Cron
	Db       *gorm.DB
	Gemini   *gemini.Gemini
	Handlers []Handler `group:"handlers"`
}

func Run(params BootstrapParams) {
	ctx := context.Background()

	// Middleware for logging requests and responses
	params.Srv.Use(func(ctx iris.Context) {
		start := time.Now()

		params.Zap.Sugar().Infof("request: method=%s, path=%s, ip=%s, user-agent=%s",
			ctx.Method(), ctx.Path(), ctx.RemoteAddr(), ctx.GetHeader("User-Agent"))

		ctx.Next()

		duration := time.Since(start)
		params.Zap.Sugar().Infof("response: status=%d, duration=%s, path=%s",
			ctx.GetStatusCode(), duration, ctx.Path())
	})

	// Define routes
	params.Srv.Get("/", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello, world🌍"})
	})

	group := params.Srv.Party("/api/v1")
	for _, handler := range params.Handlers {
		handler.InitRoutes(group)
	}

	// Initialize cron scheduler
	c := cron.New()

	// Add cron job for GenerateReport - Every Monday at 9:00 AM
	c.AddFunc("0 9 * * 1", func() {
		params.Zap.Sugar().Info("Running GenerateReport cron job...")
		if err := cronx.GenerateReport(ctx, params.Db, params.Gemini); err != nil {
			params.Zap.Sugar().Errorf("GenerateReport failed: %v", err)
		}
	})

	// Add cron job for GenerateMission - Every day at midnight
	c.AddFunc("0 0 * * *", func() {
		params.Zap.Sugar().Info("Running GenerateMission cron job...")
		if err := cronx.GenerateMission(ctx, params.Db); err != nil {
			params.Zap.Sugar().Errorf("GenerateMission failed: %v", err)
		}
	})

	// Add cron job for CreateTracker - Every day at midnight
	c.AddFunc("0 0 * * *", func() {
		params.Zap.Sugar().Info("Running CreateTracker cron job...")
		if err := cronx.CreateTracker(ctx, params.Db); err != nil {
			params.Zap.Sugar().Errorf("CreateTracker failed: %v", err)
		}
	})

	// Start the cron scheduler
	c.Start()
	defer c.Stop()

	params.Zap.Sugar().Info("Server is running on port 8080...")
	if err := params.Srv.Run(iris.Addr(fmt.Sprintf("%s:%s", params.Env.AppAddr, params.Env.AppPort))); err != nil {
		params.Zap.Sugar().Fatal("Server failed to start:", err)
	}
}
