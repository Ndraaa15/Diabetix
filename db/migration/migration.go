package migration

import (
	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"github.com/Ndraaa15/diabetix-server/pkg/config"
	"github.com/Ndraaa15/diabetix-server/pkg/env"
	"go.uber.org/zap"
)

func Migrate(env *env.Env, action string) {
	db := config.NewGorm(env)

	switch action {
	case "up":
		if err := db.AutoMigrate(
			&domain.Level{},
			&domain.User{},
			&domain.Report{},
			&domain.Article{},
			&domain.BMI{},
			&domain.Tracker{},
			&domain.TrackerDetail{},
			&domain.Mission{},
			&domain.UserMission{},
			&domain.Doctor{},
			&domain.Consultation{},
			&domain.DoctorSchedule{},
			&domain.Personalization{},
		); err != nil {
			zap.S().Fatal(err)
		}
	case "down":
		if err := db.Migrator().DropTable(
			&domain.Level{},
			&domain.User{},
			&domain.Report{},
			&domain.Article{},
			&domain.BMI{},
			&domain.Tracker{},
			&domain.TrackerDetail{},
			&domain.Mission{},
			&domain.UserMission{},
			&domain.Doctor{},
			&domain.Consultation{},
			&domain.DoctorSchedule{},
			&domain.Personalization{},
		); err != nil {
			zap.S().Fatal(err)
		}
	}

	zap.S().Info("Migration done")
}
