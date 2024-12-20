package seed

import (
	"github.com/Ndraaa15/diabetix-server/pkg/config"
	"github.com/Ndraaa15/diabetix-server/pkg/env"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Seeder func(db *gorm.DB) error

var seeders = map[string]Seeder{}

func Execute(env *env.Env, name string) {
	db := config.NewGorm(env)

	tx := db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			zap.S().Fatal(r)
		}
	}()

	RegisterSeeder()

	if name == "" {
		seederOrder := []string{
			"level",
			"user",
			"report",
			"article",
			"mission",
			"user_mission",
			"bmi",
			"tracker",
			"tracker_detail",
			"personalization",
			"doctor",
			"doctor_schedule",
		}

		for _, seederName := range seederOrder {
			seedFunc, exists := seeders[seederName]
			if !exists {
				zap.S().Fatalf("Seeder %s not found", seederName)
				return
			}
			if err := seedFunc(tx); err != nil {
				tx.Rollback()
				zap.S().Fatal(err)
				return
			}
			zap.S().Infof("Seeder %s done", seederName)
		}

		if err := tx.Commit().Error; err != nil {
			zap.S().Fatal(err)
		}
		return
	}

	seederFunc, exists := seeders[name]
	if !exists {
		zap.S().Fatal("Seeder not found")
		return
	}

	if err := seederFunc(tx); err != nil {
		tx.Rollback()
		zap.S().Fatal(err)
		return
	}

	if err := tx.Commit().Error; err != nil {
		zap.S().Fatal(err)
	}
}

func RegisterSeeder() {
	seeders["level"] = LevelSeeder()
	seeders["user"] = UserSeeder()
	seeders["report"] = ReportSeeder()
	seeders["article"] = ArticleSeeder()
	seeders["mission"] = MissionSeeder()
	seeders["user_mission"] = UserMissionSeeder()
	seeders["bmi"] = BMISeeder()
	seeders["tracker"] = TrackerSeeder()
	seeders["tracker_detail"] = TrackerDetailSeeder()
	seeders["personalization"] = PersonalizationSeeder()
	seeders["doctor"] = DoctorSeeder()
	seeders["doctor_schedule"] = DoctorScheduleSeeder()
}
