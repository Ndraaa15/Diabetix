package bootstrap

import (
	"net/http"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

type Handler interface {
	InitRoutes(app router.Party)
}

type HandlerParams struct {
	dig.In
	Handlers []Handler `group:"handlers"`
}

func Run(iris *iris.Application, zap *zap.Logger, params HandlerParams) {
	group := iris.Party("/api/v1")
	for _, handler := range params.Handlers {
		handler.InitRoutes(group)
	}

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      iris,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	zap.Sugar().Info("Server is running on port 8080...")
	if err := srv.ListenAndServe(); err != nil {
		zap.Sugar().Fatal("Server failed to start:", err)
	}
}
