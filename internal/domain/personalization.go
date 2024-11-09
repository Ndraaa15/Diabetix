package domain

type Personalization struct {
	UserID         string `json:"userID" gorm:"type:varchar(255);not null;primaryKey"`
	Gender         string `json:"gender" gorm:"type:varchar(255);not null"`
	Age            uint8  `json:"age" gorm:"type:integer;not null"`
	FrequenceSport string `json:"frequenceSport" gorm:"type:varchar(255);not null"`
}
