package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/joe-ngu/gogym/store"
	"github.com/joe-ngu/gogym/types"
)

type WorkoutHandler struct {
	db store.DB
}

func NewWorkoutHandler(db store.DB) *WorkoutHandler {
	return &WorkoutHandler{db: db}
}

func (h *WorkoutHandler) Create(w http.ResponseWriter, r *http.Request) error {
	log.Println("Handling CREATE request - Method:", r.Method)
	var req types.WorkoutPayload
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return InvalidJSON()
	}
	defer r.Body.Close()

	if errs := req.Validate(); len(errs) > 0 {
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

	var req types.WorkoutPayload
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return InvalidJSON()
	}
	defer r.Body.Close()

	if errs := req.Validate(); len(errs) > 0 {
		return InvalidRequestData(errs)
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
	if err := h.db.DeleteWorkout(workoutID); err != nil {
		return err
	}
	return writeJSON(w, http.StatusNoContent, nil)
}
