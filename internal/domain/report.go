package domain

import "time"

type Report struct {
	ID           uint64        `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID       string        `json:"userID" gorm:"type:varchar(255);not null"`
	StartDate    time.Time     `json:"startDate" gorm:"type:timestamp;not null"`
	EndDate      time.Time     `json:"endDate" gorm:"type:timestamp;not null"`
	Trackers     []Tracker     `json:"trackers" gorm:"foreignKey:ReportID;references:ID"`
	UserMissions []UserMission `json:"missions" gorm:"foreignKey:ReportID;references:ID"`
	Advice       string        `json:"advice" gorm:"type:text;not null"`
	CreatedAt    time.Time     `json:"createdAt" gorm:"type:timestamp;not null;autoCreateTime"`
	UpdatedAt    time.Time     `json:"updatedAt" gorm:"type:timestamp;not null;autoUpdateTime"`
}
