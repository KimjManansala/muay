package db

import (
	"log"

	"github.com/KimjManansala/muay/backend/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := "postgres://pg:pass@localhost:5432/crud"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(
		&models.User{},
		&models.Athlete{},
		&models.Coach{},
		&models.Gym{},
		&models.CoachGym{},
		&models.AthleteGym{},
		&models.Workout{},
		&models.AthleteSnapshot{},
		&models.AthleteWorkout{},
		&models.Exercise{},
		&models.WorkoutExercise{},
	)

	return db
}
