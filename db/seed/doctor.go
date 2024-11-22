package seed

import (
	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"gorm.io/gorm"
)

func DoctorSeeder() Seeder {
	return func(db *gorm.DB) error {
		doctors := []domain.Doctor{
			{
				ID:             1,
				Name:           "Dr. John Doe",
				Specialist:     "General Practitioner",
				Image:          "",
				Description:    "Dr. John Doe is a general practitioner with over 10 years of experience. He is passionate about providing quality healthcare services to his patients.",
				YearExperience: 10,
				Location:       "Jakarta",
			},
			{
				ID:             2,
				Name:           "Dr. Jane Smith",
				Specialist:     "Dermatologist",
				Image:          "",
				Description:    "Dr. Jane Smith is a dermatologist with over 5 years of experience. She is dedicated to helping her patients achieve healthy and beautiful skin.",
				YearExperience: 5,
				Location:       "Jakarta",
			},
			{
				ID:             3,
				Name:           "Dr. Alice Cooper",
				Specialist:     "Pediatrician",
				Image:          "",
				Description:    "Dr. Alice Cooper has been providing exceptional care for children for over 8 years.",
				YearExperience: 8,
				Location:       "Surabaya",
			},
			{
				ID:             4,
				Name:           "Dr. Michael Lee",
				Specialist:     "Orthopedic Surgeon",
				Image:          "",
				Description:    "Dr. Michael Lee specializes in orthopedic surgeries with over 12 years of expertise.",
				YearExperience: 12,
				Location:       "Bandung",
			},
			{
				ID:             5,
				Name:           "Dr. Emily Brown",
				Specialist:     "Cardiologist",
				Image:          "",
				Description:    "Dr. Emily Brown is a cardiologist committed to ensuring heart health for her patients.",
				YearExperience: 15,
				Location:       "Jakarta",
			},
			{
				ID:             6,
				Name:           "Dr. Robert Wilson",
				Specialist:     "Neurologist",
				Image:          "",
				Description:    "Dr. Robert Wilson is a neurologist with a decade of experience in treating neurological disorders.",
				YearExperience: 10,
				Location:       "Medan",
			},
			{
				ID:             7,
				Name:           "Dr. Sophia Taylor",
				Specialist:     "Ophthalmologist",
				Image:          "",
				Description:    "Dr. Sophia Taylor specializes in eye care and has been serving her patients for 7 years.",
				YearExperience: 7,
				Location:       "Semarang",
			},
			{
				ID:             8,
				Name:           "Dr. William Scott",
				Specialist:     "Oncologist",
				Image:          "",
				Description:    "Dr. William Scott is dedicated to providing top-notch care for cancer patients with over 14 years of experience.",
				YearExperience: 14,
				Location:       "Yogyakarta",
			},
			{
				ID:             9,
				Name:           "Dr. Laura Johnson",
				Specialist:     "Psychiatrist",
				Image:          "",
				Description:    "Dr. Laura Johnson helps her patients manage mental health issues with a compassionate approach.",
				YearExperience: 9,
				Location:       "Bali",
			},
			{
				ID:             10,
				Name:           "Dr. Oliver Evans",
				Specialist:     "Gastroenterologist",
				Image:          "",
				Description:    "Dr. Oliver Evans specializes in digestive health with a focus on patient care and well-being.",
				YearExperience: 6,
				Location:       "Malang",
			},
		}

		if err := db.Model(&domain.Doctor{}).Preload("DoctorSchedules").CreateInBatches(&doctors, len(doctors)).Error; err != nil {
			return err
		}

		return nil
	}
}
