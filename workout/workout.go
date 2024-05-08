package workout

import (
	"time"

	"github.com/google/uuid"
)

// Workout model
type Workout struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Exercises []uint    `json:"exercises"`
	CreatedAt time.Time `json:"created_at"`
	Date      time.Time `json:"date"`
}

func NewWorkout(name string, exercises []uint) (*Workout, error) {
	return &Workout{
		Name:      name,
		Exercises: exercises,
		CreatedAt: time.Now().UTC(),
	}, nil
}
