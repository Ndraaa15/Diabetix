package domain

type Doctor struct {
	ID              uint64           `json:"id" gorm:"primaryKey;autoIncrement"`
	Name            string           `json:"name" gorm:"type:varchar(255);not null"`
	Image           string           `json:"image" gorm:"type:varchar(255);not null"`
	Description     string           `json:"description" gorm:"type:text;not null"`
	YearExperience  uint64           `json:"yearExperience" gorm:"type:integer;not null"`
	Location        string           `json:"location" gorm:"type:varchar(255);not null"`
	Specialist      string           `json:"specialist" gorm:"type:varchar(255);not null"`
	DoctorSchedules []DoctorSchedule `json:"doctorSchedules,omitempty" gorm:"references:ID;foreignKey:DoctorID"`
	CreatedAt       string           `json:"createdAt" gorm:"type:timestamp;not null;autoCreateTime"`
	UpdatedAt       string           `json:"updatedAt" gorm:"type:timestamp;not null;autoUpdateTime"`
}

type DoctorSchedule struct {
	ID        uint64 `json:"id" gorm:"primaryKey;autoIncrement"`
	DoctorID  uint64 `json:"doctorID" gorm:"type:integer;not null"`
	StartTime string `json:"startTime" gorm:"type:varchar(255);not null"`
	EndTime   string `json:"endTime" gorm:"type:varchar(255);not null"`
	IsOpen    bool   `json:"isOpen" gorm:"type:boolean;not null;default:false"`
	CreatedAt string `json:"createdAt" gorm:"type:timestamp;not null;autoCreateTime"`
	UpdatedAt string `json:"updatedAt" gorm:"type:timestamp;not null;autoUpdateTime"`
}
