package models

type WorkoutExercise struct {
	WorkoutId  int      `json:"workoutId" gorm:"primaryKey"`
	ExerciseId int      `json:"exerciseId" gorm:"primaryKey"`
	Workout    Workout  `json:"workout" gorm:"foreignKey:WorkoutId"`
	Exercise   Exercise `json:"exercise" gorm:"foreignKey:ExerciseId"`
}
