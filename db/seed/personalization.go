package seed

import (
	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"gorm.io/gorm"
)

func PersonalizationSeeder() Seeder {
	return func(db *gorm.DB) error {
		personalization := []domain.Personalization{
			{
				UserID:              "1854723870678847488",
				Gender:              domain.PersonalizationGenderMale,
				Age:                 18,
				FrequenceSport:      domain.PersonalizationFrequenceSportOncePerWeek,
				MaxGlucose:          50.0,
				DiabetesInheritance: false,
			},
		}

		if err := db.Model(&domain.Personalization{}).CreateInBatches(&personalization, len(personalization)).Error; err != nil {
			return err
		}

		return nil
	}
}
