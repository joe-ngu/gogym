package storage

import (
	"github.com/google/uuid"
	"github.com/joe-ngu/gogym/types"
)

type DB interface {
	//Exercise
	CreateExercise(*types.Exercise) error
	GetExercises() ([]*types.Exercise, error)
	GetExercise(string) (*types.Exercise, error)
	UpdateExercise(*types.Exercise) error
	DeleteExercise(string) error

	//Workout
	CreateWorkout(*types.Workout) error
	GetWorkouts() ([]*types.Workout, error)
	GetWorkout(uuid.UUID) (*types.Workout, error)
	UpdateWorkout(*types.Workout) error
	DeleteWorkout(uuid.UUID) error

	// User
	CreateUser(*types.User) error
	GetUsers() ([]*types.User, error)
	GetUser(uuid.UUID) (*types.User, error)
	UpdateUser(*types.User) error
	DeleteUser(uuid.UUID) error
}
