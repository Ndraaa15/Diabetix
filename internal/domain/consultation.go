package domain

import "time"

type Consultation struct {
	ID               uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID           string    `json:"userID" gorm:"type:varchar(19);not null"`
	DoctorScheduleID uint64    `json:"doctorScheduleID" gorm:"type:integer;not null"`
	IsDone           bool      `json:"isDone" gorm:"type:boolean;not null;default:false"`
	MeetURL          string    `json:"meetURL" gorm:"type:varchar(255);not null"`
	CreatedAt        time.Time `json:"createdAt" gorm:"type:timestamp;not null;autoCreateTime"`
	UpdatedAt        time.Time `json:"updatedAt" gorm:"type:timestamp;not null;autoUpdateTime"`
}
