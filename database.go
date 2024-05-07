package main

import "database/sql"

type Storage interface {
	//Exercise
	CreateExercise(*Exercise) error
	GetExercises() ([]*Exercise, error)
	GetExercise(string) (*Exercise, error)
	UpdateExercise(*Exercise) error
	DeleteExercise(string) error

	//Workout
	CreateWorkout(*Workout)
	GetWorkouts() ([]*Workout, error)
	GetWorkout(uint) (*Workout, error)
	UpdateWorkout(*Workout) error
	DeleteWorkout(uint) error

	// User
	CreateUser(*User) error
	GetUsers() ([]*User, error)
	GetUser(uint) (*User, error)
	UpdateUser(*User) error
	DeleteUser(uint) error
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := ""
}
