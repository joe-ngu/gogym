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
	UpdateExercise(string, *types.Exercise) (*types.Exercise, error)
	DeleteExercise(string) error

	//Workout
	CreateWorkout(uuid.UUID, *types.Workout) error
	GetWorkouts(uuid.UUID) ([]*types.Workout, error)
	GetWorkout(uuid.UUID, uuid.UUID) (*types.Workout, error)
	UpdateWorkout(uuid.UUID, *types.Workout) (*types.Workout, error)
	DeleteWorkout(uuid.UUID, uuid.UUID) error

	// User
	CreateUser(*types.User) error
	GetUsers() ([]*types.User, error)
	GetUserByID(uuid.UUID) (*types.User, error)
	GetUserByUsername(string) (*types.User, error)
	UpdateUser(uuid.UUID, *types.User) (*types.User, error)
	DeleteUser(uuid.UUID) error
}
