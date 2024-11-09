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
				TotalGlucose: 110.5,
				Status:       "Normal",
				CreatedAt:    time.Now().Format(time.RFC3339),
				UpdatedAt:    time.Now().Format(time.RFC3339),
				ReportID:     1,
			},
			{
				ID:           2,
				UserID:       "1854723870678847488",
				TotalGlucose: 90.3,
				Status:       "Low",
				CreatedAt:    time.Now().Format(time.RFC3339),
				UpdatedAt:    time.Now().Format(time.RFC3339),
				ReportID:     1,
			},
			{
				ID:           3,
				UserID:       "1854723870678847488",
				TotalGlucose: 135.2,
				Status:       "High",
				CreatedAt:    time.Now().Format(time.RFC3339),
				UpdatedAt:    time.Now().Format(time.RFC3339),
				ReportID:     1,
			},
			{
				ID:           4,
				UserID:       "1854723870678847488",
				TotalGlucose: 78.9,
				Status:       "Low",
				CreatedAt:    time.Now().Format(time.RFC3339),
				UpdatedAt:    time.Now().Format(time.RFC3339),
				ReportID:     1,
			},
			{
				ID:           5,
				UserID:       "1854723870678847488",
				TotalGlucose: 120.0,
				Status:       "Normal",
				CreatedAt:    time.Now().Format(time.RFC3339),
				UpdatedAt:    time.Now().Format(time.RFC3339),
				ReportID:     1,
			},
		}

		if err := db.Model(&domain.Tracker{}).CreateInBatches(&trackers, len(trackers)).Error; err != nil {
			return err
		}

		return nil
	}
}
