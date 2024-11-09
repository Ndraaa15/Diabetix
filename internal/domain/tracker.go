package domain

type Tracker struct {
	ID             uint64          `json:"id" gorm:"primaryKey;autoIncrement"`
	TotalGlucose   float64         `json:"totalGlucose" gorm:"type:numeric;not null"`
	Status         string          `json:"status" gorm:"type:varchar(255);not null"`
	CreatedAt      string          `json:"createdAt" gorm:"type:timestamp;not null;autoCreateTime"`
	UpdatedAt      string          `json:"updatedAt" gorm:"type:timestamp;not null;autoUpdateTime"`
	TrackerDetails []TrackerDetail `json:"glucoseTrackerDetails" gorm:"foreignKey:TrackerID;references:ID"`
	ReportID       uint64          `json:"reportID" gorm:"type:integer;not null"`
}

type TrackerDetail struct {
	ID           uint64  `json:"id" gorm:"primaryKey;autoIncrement"`
	TrackerID    uint64  `json:"trackerID" gorm:"type:integer;not null"`
	FoodImage    string  `json:"foodImage" gorm:"type:varchar(255);not null"`
	FoodName     string  `json:"foodName" gorm:"type:varchar(255);not null"`
	Glucose      float64 `json:"glucose" gorm:"type:numeric;not null"`
	Calory       float64 `json:"calory" gorm:"type:numeric;not null"`
	Fat          float64 `json:"fat" gorm:"type:numeric;not null"`
	Protein      float64 `json:"protein" gorm:"type:numeric;not null"`
	Carbohydrate float64 `json:"carbohydrate" gorm:"type:numeric;not null"`
	CreatedAt    string  `json:"createdAt" gorm:"type:timestamp;not null;autoCreateTime"`
	UpdatedAt    string  `json:"updatedAt" gorm:"type:timestamp;not null;autoUpdateTime"`
}
