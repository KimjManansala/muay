package models

type AthleteGym struct {
	AthleteId int     `json:"athleteId" gorm:"primaryKey"`
	Athlete   Athlete `json:"athlete" gorm:"foreignKey:AthleteId"`
	GymId     int     `json:"gymId" gorm:"primaryKey"`
	Gym       Gym     `json:"gym" gorm:"foreignKey:GymId"`
}
