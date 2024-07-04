package types

import (
	"time"

	"github.com/google/uuid"
)

// Workout model
type Workout struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Name      string    `json:"name"`
	Exercises []string    `json:"exercises"`
	CreatedAt time.Time `json:"created_at"`
}

func NewWorkout(userID uuid.UUID, name string, exercises []string) (*Workout, error) {
	return &Workout{
    ID: uuid.New(),
    UserID: userID,
		Name:      name,
		Exercises: exercises,
		CreatedAt: time.Now().UTC(),
	}, nil
}

type WorkoutPayload struct {
	UserID    uuid.UUID `json:"user"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	Exercises []string  `json:"exercises"`
}

func (p *WorkoutPayload) Validate() map[string]string {
	errs := make(map[string]string)
	return errs
}
