package domain

import "time"

type BMI struct {
	ID        uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	Height    float64   `json:"height" gorm:"type:numeric;not null"`
	Weight    float64   `json:"weight" gorm:"type:numeric;not null"`
	Status    BMIStatus `json:"status" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"type:timestamp;not null;autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"type:timestamp;not null;autoUpdateTime"`
}

type BMIStatus uint8

const (
	BMIStatusUnknown BMIStatus = 0
	BMIStatusUnder   BMIStatus = 1
	BMIStatusNormal  BMIStatus = 2
	BMIStatusOver    BMIStatus = 3
)

var BMIStatusMap = map[BMIStatus]string{
	BMIStatusUnder:  "Under",
	BMIStatusNormal: "Normal",
	BMIStatusOver:   "Over",
}

func (s BMIStatus) String() string {
	return BMIStatusMap[s]
}

func (s BMIStatus) Value() uint8 {
	return uint8(s)
}
