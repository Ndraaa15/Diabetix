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
	CreateConsultation(ctx context.Context, data domain.Consultation) error
	GetDoctorScheduleByID(ctx context.Context, doctorScheduleID uint64) (domain.DoctorSchedule, error)
	UpdateDoctorSchedule(ctx context.Context, doctorSchedule domain.DoctorSchedule) error
	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
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
	// Todo : if consultation is already exist with the same doctor schedule id and user id and status is not done, return error
	if err := r.db.WithContext(ctx).Model(&domain.Consultation{}).Create(&data).Error; err != nil {
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

func (r *DoctorStore) GetDoctorScheduleByID(ctx context.Context, doctorScheduleID uint64) (domain.DoctorSchedule, error) {
	var doctorSchedule domain.DoctorSchedule
	if err := r.db.WithContext(ctx).Model(&domain.DoctorSchedule{}).Where("id = ?", doctorScheduleID).First(&doctorSchedule).Error; err != nil {
		return doctorSchedule, err
	}

	return doctorSchedule, nil
}

func (r *DoctorStore) UpdateDoctorSchedule(ctx context.Context, doctorSchedule domain.DoctorSchedule) error {
	if err := r.db.WithContext(ctx).Model(&domain.DoctorSchedule{}).Updates(&doctorSchedule).Error; err != nil {
		return err
	}

	return nil
}

func (r *DoctorStore) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	tx := r.db.WithContext(ctx).Begin()
	if err := fn(ctx); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
