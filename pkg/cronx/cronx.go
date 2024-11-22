package cronx

import (
	"context"
	"fmt"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"github.com/Ndraaa15/diabetix-server/pkg/gemini"
	"gorm.io/gorm"
)

func GenerateReport(ctx context.Context, db *gorm.DB, gemini *gemini.Gemini) error {
	// Fetch reports from the database
	var reports []domain.Report
	if err := db.WithContext(ctx).Model(&domain.Report{}).Find(&reports).Error; err != nil {
		return fmt.Errorf("failed to fetch reports: %w", err)
	}

	if len(reports) == 0 {
		fmt.Println("No reports found to process.")
		return nil
	}

	// Process each report to generate advice
	for i, report := range reports {
		res, err := gemini.GenerateReportAdvice(ctx, report)
		if err != nil {
			return fmt.Errorf("failed to generate advice for report ID %d: %w", report.ID, err)
		}
		reports[i].Advice = res.Advice
	}

	// Save updated reports back to the database in a single batch
	if err := db.WithContext(ctx).Save(&reports).Error; err != nil {
		return fmt.Errorf("failed to save updated reports: %w", err)
	}

	fmt.Println("Reports successfully updated with advice.")
	return nil
}

// CreateTracker creates daily trackers for all active users.
func CreateTracker(ctx context.Context, db *gorm.DB) error {
	// Fetch only active users
	var users []domain.User
	if err := db.WithContext(ctx).Model(&domain.User{}).Where("is_active = ?", true).Find(&users).Error; err != nil {
		return fmt.Errorf("failed to fetch active users: %w", err)
	}

	if len(users) == 0 {
		fmt.Println("No active users found.")
		return nil
	}

	// Prepare tracker records
	trackers := make([]domain.Tracker, len(users))
	for i, user := range users {
		trackers[i] = domain.Tracker{
			UserID: user.ID,
			Status: "Pending",
		}
	}

	// Batch insert trackers using a transaction
	tx := db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&trackers).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create trackers: %w", err)
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	fmt.Println("Successfully created daily trackers.")
	return nil
}

// GenerateMission updates all user missions and assigns new missions daily.
func GenerateMission(ctx context.Context, db *gorm.DB) error {
	// Start a database transaction
	tx := db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Mark all active missions as inactive
	if err := tx.Model(&domain.UserMission{}).Where("is_active = ?", true).Update("is_active", false).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to mark active missions as inactive: %w", err)
	}

	// Fetch all users
	var users []domain.User
	if err := tx.Model(&domain.User{}).Where("is_active = ?", true).Find(&users).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to fetch users: %w", err)
	}

	// Fetch all missions
	var missions []domain.Mission
	if err := tx.Model(&domain.Mission{}).Find(&missions).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to fetch missions: %w", err)
	}

	// Prepare new user missions
	newUserMissions := make([]domain.UserMission, 0)
	for _, user := range users {
		for _, mission := range missions {
			newUserMissions = append(newUserMissions, domain.UserMission{
				UserID:    user.ID,
				MissionID: mission.ID,
				IsActive:  true,
			})
		}
	}

	// Insert new user missions
	if err := tx.Create(&newUserMissions).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create new user missions: %w", err)
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	fmt.Println("Successfully generated new user missions.")
	return nil
}
