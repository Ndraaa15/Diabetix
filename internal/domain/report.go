package domain

import "time"

type Report struct {
	ID        uint64        `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    string        `json:"userID" gorm:"type:varchar(255);not null"`
	StartDate time.Time     `json:"startDate" gorm:"type:timestamp;not null"`
	EndDate   time.Time     `json:"endDate" gorm:"type:timestamp;not null"`
	Trackers  []Tracker     `json:"trackers" gorm:"foreignKey:ReportID;references:ID"`
	Missions  []UserMission `json:"missions" gorm:"foreignKey:ReportID;references:ID"`
}
