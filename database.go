package main

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/joe-ngu/gogym/exercise"
	"github.com/joe-ngu/gogym/user"
	"github.com/joe-ngu/gogym/workout"
)

type Exercise exercise.Exercise
type Workout workout.Workout
type User user.User

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
	user, err := Getenv("DB_USER")
	if err != nil {
		return nil, errors.New("failed to create PostgresStore: " + err.Error())
	}
	dbname, err := Getenv("DB_NAME")
	if err != nil {
		return nil, errors.New("failed to create PostgresStore: " + err.Error())
	}
	dbpwd, err := Getenv("DB_PWD")
	if err != nil {
		return nil, errors.New("failed to create PostgresStore: " + err.Error())
	}

	connStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", user, dbname, dbpwd)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	if err := s.createExerciseTables(); err != nil {
		return err
	}
	if err := s.createWorkoutTables(); err != nil {
		return err
	}
	if err := s.createUserTables(); err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) createExerciseTables() error {
	createMuscleGroups := `
  CREATE TABLE IF NOT EXISTS muscle_groups (
    id SERIAL PRIMARY KEY,
    name TEXT,
  )`

	createExercise := `
  CREATE TABLE IF NOT EXISTS exercise (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT,
    muscle_group INTEGER REFERENCES muscle_groups(id),
  )`

	if _, err := s.db.Exec(createMuscleGroups); err != nil {
		return err
	}
	if _, err := s.db.Exec(createExercise); err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) CreateExercise(exercise *Exercise) error {
	query := `INSERT INTO exercise
  (name, muscle_group) 
  values ($1, $2)`

	if _, err := s.db.Query(
		query,
		exercise.Name,
		exercise.MuscleGroup,
	); err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) GetExercises() error {
	return nil
}

func (s *PostgresStore) GetExercise(id uuid.UUID) error {
	return nil
}

func (s *PostgresStore) UpdateExercise(exercise *Exercise) error {
	return nil
}

func (s *PostgresStore) DeleteExercise(id uuid.UUID) error {
	return nil
}

func (s *PostgresStore) createWorkoutTables() error {
	createWorkout := `
  CREATE TABLE IF NOT EXISTS workout (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT,
    created_at TIMESTAMP,
    date TIMESTAMP, 
  )`

	createWorkoutExercise := `
  CREATE TABLE IF NOT EXISTS workout_exercise (
    workout_id UUID REFERENCES workout(id)
    exercise_id UUID REFERENCES exercise(id)
    set INTEGER,
    reps INTEGER,
    load NUMERIC(6, 2),
  )`

	if _, err := s.db.Exec(createWorkout); err != nil {
		return err
	}
	if _, err := s.db.Exec(createWorkoutExercise); err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) createUserTables() error {
	createUser := `
  CREATE TABLE IF NOT EXISTS user (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    first_name TEXT,
    last_name TEXT,
    user_name TEXT,
    encrypted_password TEXT,
    created_at TIMESTAMP,
  );`

	createUserWorkout := `
  CREATE TABLe IF NOT EXISTS user_workout (
    user_id UUID REFERENCES User(id),
    workout_id UUID FOREIGN_KEY REFERENCES Workout(id)
  )`

	if _, err := s.db.Exec(createUser); err != nil {
		return err
	}
	if _, err := s.db.Exec(createUserWorkout); err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) CreateUser(user *User) error {

}
