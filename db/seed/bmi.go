package seed

import (
	"time"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"gorm.io/gorm"
)

func BMISeeder() Seeder {
	return func(db *gorm.DB) error {
		bmis := []domain.BMI{
			{UserID: "1854723870678847488", Height: 170, Weight: 70, Status: "Underweight", BMI: 25.0, CreatedAt: time.Now().AddDate(0, 0, -7), UpdatedAt: time.Now().AddDate(0, 0, -7)},
			{UserID: "1854723870678847488", Height: 165, Weight: 68, Status: "Normal", BMI: 24.98, CreatedAt: time.Now().AddDate(0, 0, -14), UpdatedAt: time.Now().AddDate(0, 0, -14)},
			{UserID: "1854723870678847488", Height: 180, Weight: 75, Status: "Normal", BMI: 23.15, CreatedAt: time.Now().AddDate(0, 0, -21), UpdatedAt: time.Now().AddDate(0, 0, -21)},
			{UserID: "1854723870678847488", Height: 160, Weight: 60, Status: "Normal", BMI: 23.44, CreatedAt: time.Now().AddDate(0, 0, -28), UpdatedAt: time.Now().AddDate(0, 0, -28)},
			{UserID: "1854723870678847488", Height: 175, Weight: 80, Status: "Overweight", BMI: 26.12, CreatedAt: time.Now().AddDate(0, 0, -35), UpdatedAt: time.Now().AddDate(0, 0, -35)},
			{UserID: "1854723870678847488", Height: 155, Weight: 50, Status: "Underweight", BMI: 20.81, CreatedAt: time.Now().AddDate(0, 0, -42), UpdatedAt: time.Now().AddDate(0, 0, -42)},
			{UserID: "1854723870678847488", Height: 168, Weight: 72, Status: "Overweight", BMI: 25.51, CreatedAt: time.Now().AddDate(0, 0, -49), UpdatedAt: time.Now().AddDate(0, 0, -49)},
			{UserID: "1854723870678847488", Height: 178, Weight: 85, Status: "Overweight", BMI: 26.83, CreatedAt: time.Now().AddDate(0, 0, -56), UpdatedAt: time.Now().AddDate(0, 0, -56)},
			{UserID: "1854723870678847488", Height: 162, Weight: 58, Status: "Normal", BMI: 22.1, CreatedAt: time.Now().AddDate(0, 0, -63), UpdatedAt: time.Now().AddDate(0, 0, -63)},
			{UserID: "1854723870678847488", Height: 170, Weight: 65, Status: "Normal", BMI: 22.49, CreatedAt: time.Now().AddDate(0, 0, -70), UpdatedAt: time.Now().AddDate(0, 0, -70)},
		}

		if err := db.Model(&domain.BMI{}).CreateInBatches(&bmis, len(bmis)).Error; err != nil {
			return err
		}

		return nil
	}
}
