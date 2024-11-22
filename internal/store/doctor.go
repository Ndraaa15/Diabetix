package store

import (
	"context"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"github.com/Ndraaa15/diabetix-server/internal/dto"
	"gorm.io/gorm"
)

type IDoctorStore interface {
	GetAllDoctor(ctx context.Context, filter dto.GetDoctorsFilter) ([]domain.Doctor, error)
	GetDoctorByID(ctx context.Context, doctorID uint64) (domain.Doctor, error)
}

type DoctorStore struct {
	db *gorm.DB
}

func NewDoctorStore(db *gorm.DB) IDoctorStore {
	return &DoctorStore{
		db: db,
	}
}

func (r *DoctorStore) GetAllDoctor(ctx context.Context, filter dto.GetDoctorsFilter) ([]domain.Doctor, error) {
	var doctors []domain.Doctor

	queryBuilder := r.db.Model(&domain.Doctor{}).Find(&doctors)
	if filter.Keyword != "" {
		queryBuilder.Where("LOWER(description) LIKE ?", "%"+filter.Keyword+"%")
		queryBuilder.Where("LOWER(specialist) LIKE ?", "%"+filter.Keyword+"%")
	}

	return doctors, nil
}

func (r *DoctorStore) GetDoctorByID(ctx context.Context, doctorID uint64) (domain.Doctor, error) {
	var doctor domain.Doctor
	if err := r.db.WithContext(ctx).Model(&domain.Doctor{}).Preload("DoctorSchedules").Where("id = ?", doctorID).First(&doctor).Error; err != nil {
		return doctor, err
	}

	return doctor, nil
}

func (r *DoctorStore) CreateConsultation(ctx context.Context, data domain.Consultation) error {
	if err := r.db.WithContext(ctx).Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func (r *DoctorStore) GetConsultationByUserID(ctx context.Context, userID string) ([]domain.Consultation, error) {
	var consultations []domain.Consultation
	if err := r.db.WithContext(ctx).Model(&domain.Consultation{}).Where("user_id = ?", userID).Find(&consultations).Error; err != nil {
		return consultations, err
	}

	return consultations, nil
}
