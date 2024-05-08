package exercise

import (
	"errors"

	"github.com/google/uuid"
)

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

// Exercise model
type Exercise struct {
	ID          uuid.UUID   `json:"id"`
	Name        string      `json:"name"`
	MuscleGroup MuscleGroup `json:"muscle_group"`
}

func NewExercise(name string, muscleGroup MuscleGroup) (*Exercise, error) {
	return &Exercise{
		Name:        name,
		MuscleGroup: muscleGroup,
	}, nil
}
