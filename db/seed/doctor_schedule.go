package seed

import (
	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"gorm.io/gorm"
)

func DoctorScheduleSeeder() Seeder {
	return func(db *gorm.DB) error {
		schedules := []domain.DoctorSchedule{
			{ID: 1, DoctorID: 1, StartTime: "08:00", EndTime: "09:00", IsOpen: true},
			{ID: 2, DoctorID: 1, StartTime: "10:00", EndTime: "11:00", IsOpen: true},
			{ID: 3, DoctorID: 1, StartTime: "13:00", EndTime: "14:00", IsOpen: true},
			{ID: 4, DoctorID: 2, StartTime: "09:00", EndTime: "10:00", IsOpen: true},
			{ID: 5, DoctorID: 2, StartTime: "11:00", EndTime: "12:00", IsOpen: true},
			{ID: 6, DoctorID: 2, StartTime: "14:00", EndTime: "15:00", IsOpen: true},
			{ID: 7, DoctorID: 3, StartTime: "08:30", EndTime: "09:30", IsOpen: true},
			{ID: 8, DoctorID: 3, StartTime: "10:30", EndTime: "11:30", IsOpen: true},
			{ID: 9, DoctorID: 3, StartTime: "14:30", EndTime: "15:30", IsOpen: true},
			{ID: 10, DoctorID: 4, StartTime: "08:00", EndTime: "09:00", IsOpen: true},
			{ID: 11, DoctorID: 4, StartTime: "10:00", EndTime: "11:00", IsOpen: true},
			{ID: 12, DoctorID: 4, StartTime: "15:00", EndTime: "16:00", IsOpen: true},
			{ID: 13, DoctorID: 5, StartTime: "09:00", EndTime: "10:00", IsOpen: true},
			{ID: 14, DoctorID: 5, StartTime: "11:00", EndTime: "12:00", IsOpen: true},
			{ID: 15, DoctorID: 5, StartTime: "14:00", EndTime: "15:00", IsOpen: true},
			{ID: 16, DoctorID: 6, StartTime: "07:00", EndTime: "08:00", IsOpen: true},
			{ID: 17, DoctorID: 6, StartTime: "10:00", EndTime: "11:00", IsOpen: true},
			{ID: 18, DoctorID: 6, StartTime: "15:00", EndTime: "16:00", IsOpen: true},
			{ID: 19, DoctorID: 7, StartTime: "08:30", EndTime: "09:30", IsOpen: true},
			{ID: 20, DoctorID: 7, StartTime: "11:30", EndTime: "12:30", IsOpen: true},
			{ID: 21, DoctorID: 7, StartTime: "13:30", EndTime: "14:30", IsOpen: true},
			{ID: 22, DoctorID: 8, StartTime: "08:00", EndTime: "09:00", IsOpen: true},
			{ID: 23, DoctorID: 8, StartTime: "10:00", EndTime: "11:00", IsOpen: true},
			{ID: 24, DoctorID: 8, StartTime: "14:00", EndTime: "15:00", IsOpen: true},
			{ID: 25, DoctorID: 9, StartTime: "08:30", EndTime: "09:30", IsOpen: true},
			{ID: 26, DoctorID: 9, StartTime: "11:30", EndTime: "12:30", IsOpen: true},
			{ID: 27, DoctorID: 9, StartTime: "15:30", EndTime: "16:30", IsOpen: true},
			{ID: 28, DoctorID: 10, StartTime: "09:00", EndTime: "10:00", IsOpen: true},
			{ID: 29, DoctorID: 10, StartTime: "13:00", EndTime: "14:00", IsOpen: true},
			{ID: 30, DoctorID: 10, StartTime: "16:00", EndTime: "17:00", IsOpen: true},
		}

		if err := db.Model(&domain.DoctorSchedule{}).CreateInBatches(&schedules, len(schedules)).Error; err != nil {
			return err
		}

		return nil
	}
}
