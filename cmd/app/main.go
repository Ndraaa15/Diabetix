package main

import (
	"flag"
	"os"

	"github.com/Ndraaa15/diabetix-server/cmd/bootstrap"
	"github.com/Ndraaa15/diabetix-server/db/migration"
	"github.com/Ndraaa15/diabetix-server/db/seed"
	"github.com/Ndraaa15/diabetix-server/internal/handler"
	"github.com/Ndraaa15/diabetix-server/internal/store"
	"github.com/Ndraaa15/diabetix-server/internal/usecase"
	"github.com/Ndraaa15/diabetix-server/pkg/cloudinary"
	"github.com/Ndraaa15/diabetix-server/pkg/config"
	"github.com/Ndraaa15/diabetix-server/pkg/env"
	"github.com/Ndraaa15/diabetix-server/pkg/gemini"
	"github.com/Ndraaa15/diabetix-server/pkg/gomail"
	"github.com/Ndraaa15/diabetix-server/pkg/midtrans"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

func main() {
	c := dig.New()

	mustProvide := func(constructor interface{}, opts ...dig.ProvideOption) {
		if err := c.Provide(constructor, opts...); err != nil {
			zap.S().Fatalf("Unable to start application: %v", err)
		}
	}

	mustProvide(config.NewIris)
	mustProvide(config.NewSnowFlake)
	mustProvide(config.NewZap)
	mustProvide(config.NewGorm)
	mustProvide(config.NewBigCache)
	mustProvide(config.NewCron)
	mustProvide(gomail.NewGomail)
	mustProvide(env.New)
	mustProvide(gemini.NewGemini)
	mustProvide(cloudinary.NewCloudinary)
	mustProvide(config.NewValidator)
	mustProvide(midtrans.NewMidtrans)

	mustProvide(store.NewAuthStore)
	mustProvide(usecase.NewAuthUsecase)

	mustProvide(store.NewTrackerStore)
	mustProvide(usecase.NewTrackerUsecase)

	mustProvide(store.NewArticleStore)
	mustProvide(usecase.NewArticleUsecase)

	mustProvide(store.NewMissionStore)
	mustProvide(usecase.NewMissionUsecase)

	mustProvide(store.NewBMIStore)
	mustProvide(usecase.NewBMIUsecase)

	mustProvide(store.NewUserStore)
	mustProvide(usecase.NewUserUsecase)

	mustProvide(store.NewReportStore)
	mustProvide(usecase.NewReportUsecase)

	mustProvide(usecase.NewFileUploadUsecase)

	mustProvide(store.NewDoctorStore)
	mustProvide(usecase.NewDoctorUsecase)

	mustProvide(handler.NewReportHandler, dig.Group("handlers"))
	mustProvide(handler.NewMissionHandler, dig.Group("handlers"))
	mustProvide(handler.NewBMIHandler, dig.Group("handlers"))
	mustProvide(handler.NewArticleHandler, dig.Group("handlers"))
	mustProvide(handler.NewFileUploadHandler, dig.Group("handlers"))
	mustProvide(handler.NewAuthHandler, dig.Group("handlers"))
	mustProvide(handler.NewTrackerHandler, dig.Group("handlers"))
	mustProvide(handler.NewUserHandler, dig.Group("handlers"))
	mustProvide(handler.NewDoctorHandler, dig.Group("handlers"))

	if err := c.Invoke(func(e *env.Env) {
		handleArgs(e)
	}); err != nil {
		panic(err)
	}

	if err := c.Invoke(bootstrap.Run); err != nil {
		zap.S().Fatalf("Unable to start application: %v", err)
	}
}

func handleArgs(env *env.Env) {
	migrateCmd := flag.NewFlagSet("migrate", flag.ExitOnError)
	seedCmd := flag.NewFlagSet("seed", flag.ExitOnError)

	migrateAction := migrateCmd.String("action", "", "Specify 'up' or 'down' for migration")
	seedDomain := seedCmd.String("name", "", "Specify a domain for seeding (optional)")

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "migrate":
			if err := migrateCmd.Parse(os.Args[2:]); err != nil {
				zap.S().Fatalf("Unable to parse migrate command: %v", err)
			}

			if *migrateAction == "" {
				zap.S().Fatal("Action is required")
			}
			migration.Migrate(env, *migrateAction)
			os.Exit(1)
		case "seed":
			if err := seedCmd.Parse(os.Args[2:]); err != nil {
				zap.S().Fatalf("Unable to parse seed command: %v", err)
			}
			seed.Execute(env, *seedDomain)
			os.Exit(1)
		}
	}
}
