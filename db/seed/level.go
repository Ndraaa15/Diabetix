package seed

import (
	"time"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"gorm.io/gorm"
)

func LevelSeeder() Seeder {
	return func(db *gorm.DB) error {
		levels := []domain.Level{
			{ID: 1, Name: "Level 1", TotalExp: 100, NextLevel: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: 2, Name: "Level 2", TotalExp: 200, NextLevel: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: 3, Name: "Level 3", TotalExp: 300, NextLevel: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: 4, Name: "Level 4", TotalExp: 400, NextLevel: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: 5, Name: "Level 5", TotalExp: 500, NextLevel: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: 6, Name: "Level 6", TotalExp: 600, NextLevel: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: 7, Name: "Level 7", TotalExp: 700, NextLevel: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: 8, Name: "Level 8", TotalExp: 800, NextLevel: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: 9, Name: "Level 9", TotalExp: 900, NextLevel: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: 10, Name: "Level 10", TotalExp: 1000, NextLevel: 0, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		}

		if err := db.Model(&domain.Level{}).CreateInBatches(&levels, len(levels)).Error; err != nil {
			return err
		}

		return nil
	}
}
