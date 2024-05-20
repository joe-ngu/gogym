package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/joe-ngu/gogym/storage"
	"github.com/joe-ngu/gogym/types"
)

type WorkoutPayload struct {
	UserID    uuid.UUID `json:"user"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	Exercises []string  `json:"exercises"`
}

func (p *WorkoutPayload) validate() map[string]string {
	errs := make(map[string]string)
	return errs
}

type WorkoutHandler struct {
	db storage.DB
}

func (h *WorkoutHandler) Create(w http.ResponseWriter, r *http.Request) error {
	log.Println("Handling CREATE request - Method:", r.Method)
	var req WorkoutPayload
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return InvalidJSON()
	}
	defer r.Body.Close()

	if errs := req.validate(); len(errs) > 0 {
		return InvalidRequestData(errs)
	}

	workout, err := types.NewWorkout(req.UserID, req.Name, req.Exercises)
	if err != nil {
		return err
	}

	if err := h.db.CreateWorkout(workout); err != nil {
		return err
	}
	return writeJSON(w, http.StatusCreated, workout)
}

func (h *WorkoutHandler) GetAll(w http.ResponseWriter, r *http.Request) error {
	log.Println("Handling READ ALL request - Method:", r.Method)

	workouts, err := h.db.GetWorkouts()
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, workouts)
}

func (h *WorkoutHandler) Get(w http.ResponseWriter, r *http.Request) error {
	log.Println("Handling READ request - Method:", r.Method)
	id := r.PathValue("id")

	workoutID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	workout, err := h.db.GetWorkout(workoutID)
	if err != nil {
		return err
	}
	if workout == nil {
		return writeJSON(w, http.StatusNotFound, workout)
	}
	return writeJSON(w, http.StatusOK, workout)
}

func (h *WorkoutHandler) Update(w http.ResponseWriter, r *http.Request) error {
	log.Println("Handling UPDATE request - Method:", r.Method)
	id := r.PathValue("id")
	workoutID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	exists, _ := h.db.GetWorkout(workoutID)
	if exists == nil {
		return errors.New("Workout to update does not exist")
	}

	var req WorkoutPayload
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return InvalidJSON()
	}
	defer r.Body.Close()

	if errors := req.validate(); len(errors) > 0 {
		return InvalidRequestData(errors)
	}

	workout, err := types.NewWorkout(req.UserID, req.Name, req.Exercises)
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, workout)
}

func (h *WorkoutHandler) Delete(w http.ResponseWriter, r *http.Request) error {
	log.Println("Handling DELETE request - Method:", r.Method)
	id := r.PathValue("id")
	workoutID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	err = h.db.DeleteWorkout(workoutID)
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusNoContent, nil)
}
