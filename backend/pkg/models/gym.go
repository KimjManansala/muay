package models

type Gym struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
