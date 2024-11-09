package seed

import (
	"time"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"gorm.io/gorm"
)

func UserMissionSeeder() Seeder {
	return func(db *gorm.DB) error {
		userMissions := []domain.UserMission{
			{
				UserID:    "1854723870678847488",
				MissionID: 1,
				IsDone:    true,
				ReportID:  1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				UserID:    "1854723870678847488",
				MissionID: 2,
				IsDone:    false,
				ReportID:  1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				UserID:    "1854723870678847488",
				MissionID: 3,
				IsDone:    false,
				ReportID:  1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}

		if err := db.Model(&domain.UserMission{}).CreateInBatches(&userMissions, len(userMissions)).Error; err != nil {
			return err
		}

		return nil
	}
}
