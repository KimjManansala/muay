package models

type Athlete struct {
	Id     int  `json:"id" gorm:"primaryKey"`
	UserId int  `json:"userId"`
	User   User `json:"user" gorm:"foreignKey:UserId"`
}
