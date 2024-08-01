package models

type Exercise struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	Sport       string `json:"sport"` // TODO SHOULD MAKE ENUM
	Description string `json:"description"`
}
