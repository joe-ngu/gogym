package store

import (
	"github.com/google/uuid"
	"github.com/joe-ngu/gogym/types"
)

type DB interface {
	//Exercise
	CreateExercise(*types.Exercise) error
	GetExercises() ([]*types.Exercise, error)
	GetExercise(string) (*types.Exercise, error)
	UpdateExercise(*types.Exercise) (*types.Exercise, error)
	DeleteExercise(string) error

	//Workout
	CreateWorkout(*types.Workout) error
	GetWorkouts() ([]*types.Workout, error)
	GetWorkout(uuid.UUID) (*types.Workout, error)
	UpdateWorkout(*types.Workout) (*types.Workout, error)
	DeleteWorkout(uuid.UUID) error

	// User
	CreateUser(*types.User) error
	GetUsers() ([]*types.User, error)
	GetUserByID(uuid.UUID) (*types.User, error)
	GetUserByUsername(string) (*types.User, error)
	UpdateUser(*types.User) (*types.User, error)
	DeleteUser(uuid.UUID) error
}
