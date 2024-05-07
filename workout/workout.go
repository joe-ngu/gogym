package workout

import "time"

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
