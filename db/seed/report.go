package seed

import (
	"time"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"gorm.io/gorm"
)

func ReportSeeder() Seeder {
	return func(db *gorm.DB) error {
		startDate, err := time.Parse("02-01-2006", "04-11-2024")
		if err != nil {
			return err
		}

		endDate, err := time.Parse("02-01-2006", "04-11-2024")
		if err != nil {
			return err
		}

		reports := []domain.Report{
			{
				ID:        1,
				UserID:    "1854723870678847488",
				StartDate: startDate,
				EndDate:   endDate,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}

		if err := db.Model(&domain.Report{}).CreateInBatches(&reports, len(reports)).Error; err != nil {
			return err
		}

		return nil
	}
}
