package seed

import (
	"time"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"gorm.io/gorm"
)

func TrackerDetailSeeder() Seeder {
	return func(db *gorm.DB) error {
		trackerDetails := []domain.TrackerDetail{
			{
				ID:           1,
				TrackerID:    1,
				FoodImage:    "https://example.com/food1.jpg",
				FoodName:     "Salad",
				Glucose:      5.6,
				Calory:       120.0,
				Fat:          1.5,
				Protein:      2.3,
				Carbohydrate: 18.5,
				CreatedAt:    time.Now().Format(time.RFC3339),
				UpdatedAt:    time.Now().Format(time.RFC3339),
			},
			{
				ID:           2,
				TrackerID:    2,
				FoodImage:    "https://example.com/food2.jpg",
				FoodName:     "Apple",
				Glucose:      3.1,
				Calory:       95.0,
				Fat:          0.3,
				Protein:      0.5,
				Carbohydrate: 25.1,
				CreatedAt:    time.Now().Format(time.RFC3339),
				UpdatedAt:    time.Now().Format(time.RFC3339),
			},
			{
				ID:           3,
				TrackerID:    3,
				FoodImage:    "https://example.com/food3.jpg",
				FoodName:     "Pasta",
				Glucose:      8.7,
				Calory:       220.0,
				Fat:          5.0,
				Protein:      7.5,
				Carbohydrate: 30.0,
				CreatedAt:    time.Now().Format(time.RFC3339),
				UpdatedAt:    time.Now().Format(time.RFC3339),
			},
			{
				ID:           4,
				TrackerID:    4,
				FoodImage:    "https://example.com/food4.jpg",
				FoodName:     "Smoothie",
				Glucose:      6.2,
				Calory:       150.0,
				Fat:          0.7,
				Protein:      1.2,
				Carbohydrate: 30.5,
				CreatedAt:    time.Now().Format(time.RFC3339),
				UpdatedAt:    time.Now().Format(time.RFC3339),
			},
			{
				ID:           5,
				TrackerID:    5,
				FoodImage:    "https://example.com/food5.jpg",
				FoodName:     "Rice and Vegetables",
				Glucose:      7.5,
				Calory:       300.0,
				Fat:          2.1,
				Protein:      5.3,
				Carbohydrate: 50.7,
				CreatedAt:    time.Now().Format(time.RFC3339),
				UpdatedAt:    time.Now().Format(time.RFC3339),
			},
		}

		if err := db.Model(&domain.TrackerDetail{}).CreateInBatches(&trackerDetails, len(trackerDetails)).Error; err != nil {
			return err
		}

		return nil
	}
}
