package seed

import (
	"time"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"gorm.io/gorm"
)

func TrackerSeeder() Seeder {
	return func(db *gorm.DB) error {
		trackers := []domain.Tracker{
			{
				ID:           1,
				UserID:       "1854723870678847488",
				TotalGlucose: 30,
				Status:       "Normal",
				MaxGlucose:   50,
				CreatedAt:    time.Now().AddDate(0, 0, -1),
				UpdatedAt:    time.Now(),
				ReportID:     1,
			},
			{
				ID:           2,
				UserID:       "1854723870678847488",
				TotalGlucose: 45,
				Status:       "Tinggi",
				MaxGlucose:   75,
				CreatedAt:    time.Now().AddDate(0, 0, -2),
				UpdatedAt:    time.Now(),
				ReportID:     1,
			},
			{
				ID:           3,
				UserID:       "1854723870678847488",
				TotalGlucose: 55,
				Status:       "Tinggi",
				MaxGlucose:   75,
				CreatedAt:    time.Now().AddDate(0, 0, -3),
				UpdatedAt:    time.Now(),
				ReportID:     1,
			},
			{
				ID:           4,
				UserID:       "1854723870678847488",
				TotalGlucose: 0,
				Status:       "Rendah",
				MaxGlucose:   50,
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
				ReportID:     1,
			},
		}

		if err := db.Model(&domain.Tracker{}).CreateInBatches(&trackers, len(trackers)).Error; err != nil {
			return err
		}

		return nil
	}
}
