package seed

import (
	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"gorm.io/gorm"
)

func PersonalizationSeeder() Seeder {
	return func(db *gorm.DB) error {
		personalization := []domain.Personalization{
			{
				UserID:         "1854723870678847488",
				Gender:         "Male",
				Age:            18,
				FrequenceSport: 2,
				MaxGlucose:     50.0,
			},
		}

		if err := db.Model(&domain.Personalization{}).CreateInBatches(&personalization, len(personalization)).Error; err != nil {
			return err
		}

		return nil
	}
}
