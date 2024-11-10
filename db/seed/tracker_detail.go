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
				FoodImage:    "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/diabetix/green-salad.jpg",
				FoodName:     "Salad",
				Glucose:      5.6,
				Calory:       120.0,
				Fat:          1.5,
				Protein:      2.3,
				Carbohydrate: 18.5,
				CreatedAt:    time.Now().AddDate(0, 0, -1),
				UpdatedAt:    time.Now(),
			},
			{
				ID:           2,
				TrackerID:    1,
				FoodImage:    "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/diabetix/apel4jpg-20231211013431.jpg",
				FoodName:     "Apel",
				Glucose:      3.1,
				Calory:       95.0,
				Fat:          0.3,
				Protein:      0.5,
				Carbohydrate: 25.1,
				CreatedAt:    time.Now().AddDate(0, 0, -1),
				UpdatedAt:    time.Now(),
			},
			{
				ID:           3,
				TrackerID:    1,
				FoodImage:    "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/diabetix/0a0717810b73a1672a029c29788e557b_creamy_alfredo_pasta_long_left_1080_850.jpg",
				FoodName:     "Pasta",
				Glucose:      8.7,
				Calory:       220.0,
				Fat:          5.0,
				Protein:      7.5,
				Carbohydrate: 30.0,
				CreatedAt:    time.Now().AddDate(0, 0, -2),
				UpdatedAt:    time.Now(),
			},
			{
				ID:           4,
				TrackerID:    2,
				FoodImage:    "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/diabetix/mixed-berry-breakfast-smoothie-7959466-1x1-e0ad2304222e49508cda7b73b21de921.jpg",
				FoodName:     "Smoothie",
				Glucose:      6.2,
				Calory:       150.0,
				Fat:          0.7,
				Protein:      1.2,
				Carbohydrate: 30.5,
				CreatedAt:    time.Now().AddDate(0, 0, -2),
				UpdatedAt:    time.Now(),
			},
			{
				ID:           5,
				TrackerID:    3,
				FoodImage:    "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/diabetix/thumb.jpg",
				FoodName:     "Rice and Vegetables",
				Glucose:      7.5,
				Calory:       300.0,
				Fat:          2.1,
				Protein:      5.3,
				Carbohydrate: 50.7,
				CreatedAt:    time.Now().AddDate(0, 0, -3),
				UpdatedAt:    time.Now(),
			},
		}

		if err := db.Model(&domain.TrackerDetail{}).CreateInBatches(&trackerDetails, len(trackerDetails)).Error; err != nil {
			return err
		}

		return nil
	}
}
