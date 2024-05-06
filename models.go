package main

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User model
type User struct {
	ID                uint
	FirstName         string
	LastName          string
	UserName          string
	EncryptedPassword string
	CreatedAt         time.Time
	Workouts          []uint
}

func (u *User) ValidPassword(pw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(pw)) == nil
}

func NewUser(firstName string, lastName string, userName string, password string) (*User, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		FirstName:         firstName,
		LastName:          lastName,
		UserName:          userName,
		EncryptedPassword: string(encpw),
		CreatedAt:         time.Now().UTC(),
	}, nil
}

// Workout model
type Workout struct {
	ID        uint
	Name      string
	Exercises []uint
	CreatedAt time.Time
	Date      time.Time
}

func NewWorkout(name string, exercises []uint) (*Workout, error) {
	return &Workout{
		Name:      name,
		Exercises: exercises,
		CreatedAt: time.Now().UTC(),
	}, nil
}

// Exercise model
type Exercise struct {
	ID          uint        `json:"id"`
	Name        string      `json:"name"`
	MuscleGroup MuscleGroup `json:"muscle_group"`
}

type LoginResponse struct {
	UserName string `json:"user_name"`
	Token    string `json:"token"`
}

type LoginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
}

// A muscle group enum using a struct for safer representation
type MuscleGroup struct {
	slug string
}

// Converts the value of the enum into a string
func (m MuscleGroup) String() string {
	return m.slug
}

// list of valid muscle group enum values
var (
	Unknown    = MuscleGroup{""}
	Chest      = MuscleGroup{"chest"}
	Back       = MuscleGroup{"back"}
	Shoulders  = MuscleGroup{"shoulders"}
	Arms       = MuscleGroup{"arms"}
	Abs        = MuscleGroup{"abs"}
	Glutes     = MuscleGroup{"glutes"}
	Quads      = MuscleGroup{"quads"}
	Hamstrings = MuscleGroup{"hamstrings"}
	Calves     = MuscleGroup{"calves"}
)

// Assigns a string to one of the enums if valid, else throws error at unknown enum value
func GetMuscleGroup(s string) (MuscleGroup, error) {
	switch s {
	case Chest.slug:
		return Chest, nil
	case Back.slug:
		return Back, nil
	case Shoulders.slug:
		return Shoulders, nil
	case Arms.slug:
		return Arms, nil
	case Abs.slug:
		return Abs, nil
	case Glutes.slug:
		return Glutes, nil
	case Quads.slug:
		return Quads, nil
	case Hamstrings.slug:
		return Hamstrings, nil
	case Calves.slug:
		return Calves, nil
	}

	return Unknown, errors.New("unknown muscle group: " + s)
}
