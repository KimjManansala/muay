package models

import "time"

type Workout struct {
	Id      int       `json:"id" gorm:"primaryKey"`
	Sport   string    `json:"sport"` // TODO SHOULD MAKE ENUM
	Date    time.Time `json:"date"`
	CoachId int       `json:"coachId"`
	Coach   Coach     `json:"coach" gorm:"foreignKey:CoachId"`
}
