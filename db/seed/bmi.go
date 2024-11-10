package seed

import (
	"time"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"gorm.io/gorm"
)

func BMISeeder() Seeder {
	return func(db *gorm.DB) error {
		bmis := []domain.BMI{
			{ID: 1, UserID: "1854723870678847488", Height: 170, Weight: 70, Status: "Underweight", BMI: 25.0, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: 2, UserID: "1854723870678847488", Height: 165, Weight: 68, Status: "Normal", BMI: 24.98, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: 3, UserID: "1854723870678847488", Height: 180, Weight: 75, Status: "Normal", BMI: 23.15, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: 4, UserID: "1854723870678847488", Height: 160, Weight: 60, Status: "Normal", BMI: 23.44, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: 5, UserID: "1854723870678847488", Height: 175, Weight: 80, Status: "Overweight", BMI: 26.12, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: 6, UserID: "1854723870678847488", Height: 155, Weight: 50, Status: "Underweight", BMI: 20.81, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: 7, UserID: "1854723870678847488", Height: 168, Weight: 72, Status: "Overweight", BMI: 25.51, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: 8, UserID: "1854723870678847488", Height: 178, Weight: 85, Status: "Overweight", BMI: 26.83, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: 9, UserID: "1854723870678847488", Height: 162, Weight: 58, Status: "Normal", BMI: 22.1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: 10, UserID: "1854723870678847488", Height: 170, Weight: 65, Status: "Normal", BMI: 22.49, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: 11, UserID: "1854723870678847488", Height: 182, Weight: 90, Status: "Overweight", BMI: 27.17, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		}

		if err := db.Model(&domain.BMI{}).CreateInBatches(&bmis, len(bmis)).Error; err != nil {
			return err
		}

		return nil
	}
}
