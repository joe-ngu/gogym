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
