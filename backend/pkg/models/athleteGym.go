package models

type AthleteGym struct {
	AthleteId int     `json:"athleteId" gorm:"primaryKey"`
	Athlete   Athlete `json:"athlete" gorm:"foreignKey:athleteId"`
	GymId     int     `json:"gymId" gorm:"primaryKey"`
	Gym       Gym     `json:"gym" gorm:"foreignKey:gymId"`
}
