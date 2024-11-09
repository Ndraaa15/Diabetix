package domain

import "time"

type Mission struct {
	ID        uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	Image     string    `json:"image" gorm:"type:text;not null"`
	Exp       uint64    `json:"exp" gorm:"type:integer;not null"`
	Calory    uint64    `json:"calory" gorm:"type:integer;not null"`
	Title     string    `json:"title" gorm:"type:varchar(255);not null"`
	Body      string    `json:"body" gorm:"type:text;not null"`
	Category  string    `json:"category" gorm:"type:varchar(255);not null"`
	Duration  uint64    `json:"duration" gorm:"type:integer;not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"type:timestamp;not null;autoCreateTime"`
	UpdateAt  time.Time `json:"updateAt" gorm:"type:timestamp;not null;autoUpdateTime"`
}

type UserMission struct {
	UserID    string    `json:"userID" gorm:"type:varchar(255);primaryKey"`
	MissionID uint64    `json:"missionID" gorm:"type:integer;primaryKey"`
	Mission   Mission   `json:"mission" gorm:"foreignKey:MissionID;references:ID"`
	IsDone    bool      `json:"isDone" gorm:"type:boolean;not null;default:false"`
	ReportID  uint64    `json:"reportID" gorm:"type:integer;not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"type:timestamp;not null"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"type:timestamp;not null"`
}
