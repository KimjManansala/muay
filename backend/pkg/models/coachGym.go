package models

type CoachGym struct {
	CoachId int   `json:"coachId" gorm:"primaryKey"`
	Coach   Coach `json:"coach" gorm:"foreignKey:CoachId"`
	GymId   int   `json:"gymId" gorm:"primaryKey"`
	Gym     Gym   `json:"gym" gorm:"foreignKey:GymId"`
}
