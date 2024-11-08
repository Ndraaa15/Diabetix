package domain

type Personalization struct {
	UserID         string         `json:"userID" gorm:"type:varchar(255);not null;primaryKey"`
	Gender         UserGender     `json:"gender" gorm:"type:integer;not null"`
	Age            uint8          `json:"age" gorm:"type:integer;not null"`
	FrequenceSport FrequenceSport `json:"frequenceSport" gorm:"type:integer;not null"`
}

type UserGender uint8

const (
	GenderUnknown UserGender = 0
	GenderMale    UserGender = 1
	GenderFemale  UserGender = 2
)

var GenderMap = map[UserGender]string{
	GenderMale:   "Male",
	GenderFemale: "Female",
}

func (s UserGender) String() string {
	return GenderMap[s]
}

func (s UserGender) Value() uint8 {
	return uint8(s)
}

type FrequenceSport uint8

const (
	FrequenceSportUnknown          FrequenceSport = 0
	FrequenceSportOnceAWeek        FrequenceSport = 1
	FrequenceSportOnceToThreeAWeek FrequenceSport = 2
	FrequenceSportFourToFiveAWeek  FrequenceSport = 3
	FrequenceSportFiveToSevenAWeek FrequenceSport = 4
)
