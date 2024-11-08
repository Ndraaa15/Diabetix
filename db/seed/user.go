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
				Name:     "Gede Indra Adi Brata",
				Email:    "indrabrata599@gmail.com",
				Password: hashedPassword,
				Birth:    birth,
				IsActive: true,
			},
			{
				ID:       "d296f037-e55d-400f-8fb6-21dc97cc8fad",
				Name:     "Paula Sugiarto",
				Email:    "paulaaaa@gmail.com",
				Password: hashedPassword,
				Birth:    birth,
				IsActive: true,
			},
			{
				ID:       "5d4bc029-2d99-4fa2-b853-e24febafee1d",
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
