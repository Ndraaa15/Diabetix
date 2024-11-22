package domain

import "time"

type Personalization struct {
	UserID              string                        `json:"userID" gorm:"type:varchar(19);not null;primaryKey"`
	Gender              PersonalizationGender         `json:"gender" gorm:"type:varchar(255);not null"`
	Age                 uint8                         `json:"age" gorm:"type:integer;not null"`
	FrequenceSport      PersonalizationFrequenceSport `json:"frequenceSport" gorm:"type:varchar(255);not null"`
	MaxGlucose          float64                       `json:"maxGlucose" gorm:"type:integer;not null"`
	DiabetesInheritance bool                          `json:"diabetesInheritance" gorm:"type:boolean;not null"`
	CreatedAt           time.Time                     `json:"createdAt" gorm:"type:timestamp;not null;autoCreateTime"`
	UpdatedAt           time.Time                     `json:"updatedAt" gorm:"type:timestamp;not null;autoUpdateTime"`
}

type PersonalizationFrequenceSport string

const (
	PersonalizationFrequenceSportOncePerWeek             PersonalizationFrequenceSport = "OncePerWeek"
	PersonalizationFrequenceSportOnceToThreePerWeek      PersonalizationFrequenceSport = "OnceToThreePerWeek"
	PersonalizationFrequenceSportFourToFiveTimesPerWeek  PersonalizationFrequenceSport = "FourToFiveTimesPerWeek"
	PersonalizationFrequenceSportFiveToSevenTimesPerWeek PersonalizationFrequenceSport = "FiveToSevenTimesPerWeek"
)

type PersonalizationGender string

const (
	PersonalizationGenderMale   PersonalizationGender = "Male"
	PersonalizationGenderFemale PersonalizationGender = "Female"
)
