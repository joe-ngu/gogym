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
	log.Println("Handling Workout CREATE request - Method:", r.Method)

	userID, ok := r.Context().Value(UserIDKey).(uuid.UUID)
	if !ok {
		return InvalidPermissions()
	}

	var req types.WorkoutPayload
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return InvalidJSON()
	}
	defer r.Body.Close()

	if errs := req.Validate(); len(errs) > 0 {
		return InvalidRequestData(errs)
	}

	workout, err := types.NewWorkout(req.UserID, req.Name, req.Date, req.Exercises)
	if err != nil {
		return err
	}

	workoutID, err := h.db.CreateWorkout(userID, workout)
	if err != nil {
		return err
	}

	workout, err = h.db.GetWorkout(userID, workoutID)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusCreated, workout)
}

func (h *WorkoutHandler) Get(w http.ResponseWriter, r *http.Request) error {
	log.Println("Handling Workout GET request - Method:", r.Method)

	userID, ok := r.Context().Value(UserIDKey).(uuid.UUID)
	if !ok {
		return InvalidPermissions()
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		workouts, err := h.db.GetWorkouts(userID)
		if err != nil {
			return err
		}
		return writeJSON(w, http.StatusOK, workouts)
	}

	workoutID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	workout, err := h.db.GetWorkout(userID, workoutID)
	if err != nil {
		return err
	}
	if workout == nil {
		return writeJSON(w, http.StatusNotFound, workout)
	}
	return writeJSON(w, http.StatusOK, workout)
}

func (h *WorkoutHandler) Update(w http.ResponseWriter, r *http.Request) error {
	log.Println("Handling Workout UPDATE request - Method:", r.Method)

	userID, ok := r.Context().Value(UserIDKey).(uuid.UUID)
	if !ok {
		return InvalidPermissions()
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		return InvalidQueryParams()
	}

	workoutID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	exists, _ := h.db.GetWorkout(userID, workoutID)
	if exists == nil {
		return errors.New("workout to update does not exist")
	}

	var req types.WorkoutPayload
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return InvalidJSON()
	}
	defer r.Body.Close()

	if errs := req.Validate(); len(errs) > 0 {
		return InvalidRequestData(errs)
	}

	workout, err := types.NewWorkout(req.UserID, req.Name, req.Date, req.Exercises)
	if err != nil {
		return err
	}
	workout.ID = exists.ID

	updatedWorkout, err := h.db.UpdateWorkout(userID, workout)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, updatedWorkout)
}

func (h *WorkoutHandler) Delete(w http.ResponseWriter, r *http.Request) error {
	log.Println("Handling Workout DELETE request - Method:", r.Method)

	userID, ok := r.Context().Value(UserIDKey).(uuid.UUID)
	if !ok {
		return InvalidPermissions()
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		return InvalidQueryParams()
	}

	workoutID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	if err := h.db.DeleteWorkout(userID, workoutID); err != nil {
		return err
	}
	return writeJSON(w, http.StatusNoContent, nil)
}
