package types

import (
	"time"

	"github.com/google/uuid"
)

// Exercise details model
type ExerciseDetail struct {
	ID   uuid.UUID `json:"id"`
	Sets int       `json:"sets"`
	Reps int       `json:"reps"`
	Load int       `json:"load"`
}

// Workout model
type Workout struct {
	ID        uuid.UUID        `json:"id"`
	UserID    uuid.UUID        `json:"user_id"`
	Name      string           `json:"name"`
	CreatedAt time.Time        `json:"created_at"`
	Date      time.Time        `json:"date"`
	Exercises []ExerciseDetail `json:"exercises"`
}

func NewWorkout(userID uuid.UUID, name string, date time.Time, exercises []ExerciseDetail) (*Workout, error) {
	return &Workout{
		UserID:    userID,
		Name:      name,
		Date:      date,
		Exercises: exercises,
	}, nil
}

type WorkoutPayload struct {
	UserID    uuid.UUID        `json:"user"`
	Name      string           `json:"name"`
	Date      time.Time        `json:"date"`
	Exercises []ExerciseDetail `json:"exercises"`
}

func (p *WorkoutPayload) Validate() map[string]string {
	errs := make(map[string]string)
	return errs
}
