package seed

import (
	"time"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"github.com/Ndraaa15/diabetix-server/pkg/bcrypt"
	"gorm.io/gorm"
)

func UserSeeder() Seeder {
	return func(db *gorm.DB) error {
		hashedPassword, err := bcrypt.EncryptPassword("password")
		if err != nil {
			return err
		}

		birth, err := time.Parse("2006-01-02", "2003-12-15")
		if err != nil {
			return err
		}

		users := []domain.User{
			{
				ID:       "1854723870678847488",
				Name:     "Gede Indra Adi Brata",
				Email:    "indrabrata599@gmail.com",
				Password: hashedPassword,
				Birth:    birth,
				IsActive: true,
			},
			{
				ID:       "1854724598575140864",
				Name:     "Paula Sugiarto",
				Email:    "paulaaaa@gmail.com",
				Password: hashedPassword,
				Birth:    birth,
				IsActive: true,
			},
			{
				ID:       "1854763141305470976",
				Name:     "Handedius Sando Sianipar",
				Email:    "sandogi@gmail.com",
				Password: hashedPassword,
				Birth:    birth,
				IsActive: true,
			},
		}

		if err := db.Model(&domain.User{}).CreateInBatches(&users, len(users)).Error; err != nil {
			return err
		}

		return nil
	}
}
