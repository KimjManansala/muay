package models

type AthleteWorkout struct {
	AthleteId int     `json:"athleteId" gorm:"primaryKey"`
	WorkoutId int     `json:"workoutId" gorm:"primaryKey"`
	Athlete   Athlete `json:"athlete" gorm:"foreignKey:athleteId"`
	Workout   Workout `json:"workout" gorm:"foreignKey:workoutId"`
}
