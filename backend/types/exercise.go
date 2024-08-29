package types

import (
	"encoding/json"
	"errors"
	"regexp"

	"github.com/google/uuid"
)

type MuscleGroup struct {
	slug string
}

func (m MuscleGroup) String() string {
	return m.slug
}

func (m MuscleGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.slug)
}

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

type Exercise struct {
	ID          uuid.UUID   `json:"id"`
	Name        string      `json:"name"`
	MuscleGroup MuscleGroup `json:"muscle_group"`
}

func NewExercise(name string, muscle string) (*Exercise, error) {
	muscleGroup, err := GetMuscleGroup(muscle)
	if err != nil {
		return nil, err
	}

	return &Exercise{
		Name:        name,
		MuscleGroup: muscleGroup,
	}, nil
}

type ExercisePayload struct {
	Name        string `json:"name"`
	MuscleGroup string `json:"muscle_group"`
}

func (p *ExercisePayload) Validate() map[string]string {
	errs := make(map[string]string)
	alphaRegex := regexp.MustCompile(`^[a-zA-Z ]+$`)

	if !alphaRegex.MatchString(p.Name) {
		errs["Name"] = "name must be a string containing only alphabetical characters"
	}
	if len(p.Name) < 2 || 32 < len(p.Name) {
		errs["Name"] = "name must be between 2 and 32 characters long"
	}
	if _, err := GetMuscleGroup(p.MuscleGroup); err != nil {
		errs["MuscleGroup"] = err.Error()
	}
	return errs
}
