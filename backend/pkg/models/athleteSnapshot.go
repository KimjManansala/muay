package models

import "time"

type AthleteSnapshot struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	AthleteId int       `json:"athleteId" gorm:"primaryKey"`
	Athlete   Athlete   `json:"athlete" gorm:"foreignKey:AthleteId"`
	Date      time.Time `json:"date"`
	Weight    int       `json:"weight"`
}
