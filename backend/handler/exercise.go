package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/joe-ngu/gogym/store"
	"github.com/joe-ngu/gogym/types"
)

type ExerciseHandler struct {
	db store.DB
}

func NewExerciseHandler(db store.DB) *ExerciseHandler {
	return &ExerciseHandler{db: db}
}

func (h *ExerciseHandler) Create(w http.ResponseWriter, r *http.Request) error {
	log.Println("Handling Exercise CREATE request - Method:", r.Method)

	var req types.ExercisePayload
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return InvalidJSON()
	}
	defer r.Body.Close()

	if errs := req.Validate(); len(errs) > 0 {
		return InvalidRequestData(errs)
	}

	exists, _ := h.db.GetExercise(req.Name)
	if exists != nil {
		return errors.New("exercise already exists")
	}

	exercise, err := types.NewExercise(req.Name, req.MuscleGroup)
	if err != nil {
		return err
	}

	if err := h.db.CreateExercise(exercise); err != nil {
		return err
	}

	return writeJSON(w, http.StatusCreated, exercise)
}

func (h *ExerciseHandler) Get(w http.ResponseWriter, r *http.Request) error {
	log.Println("Handling Exercise READ request - Method:", r.Method)

	name := r.URL.Query().Get("name")
	if name == "" {
		exercises, err := h.db.GetExercises()
		if err != nil {
			return err
		}
		return writeJSON(w, http.StatusOK, exercises)
	}

	exercise, err := h.db.GetExercise(name)
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, exercise)
}

func (h *ExerciseHandler) Update(w http.ResponseWriter, r *http.Request) error {
	log.Println("Handling Exercise UPDATE request - Method:", r.Method)
	var req types.ExercisePayload

	name := r.URL.Query().Get("name")
	if name == "" {
		return InvalidQueryParams()
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return InvalidJSON()
	}
	defer r.Body.Close()

	if errs := req.Validate(); len(errs) > 0 {
		return InvalidRequestData(errs)
	}

	exists, _ := h.db.GetExercise(name)
	if exists == nil {
		return errors.New("exercise to update does not exist")
	}

	exercise, err := types.NewExercise(req.Name, req.MuscleGroup)
	if err != nil {
		return err
	}

	updated_exercise, err := h.db.UpdateExercise(name, exercise)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, updated_exercise)
}

func (h *ExerciseHandler) Delete(w http.ResponseWriter, r *http.Request) error {
	log.Println("Handling Exercise DELETE request - Method:", r.Method)

	name := r.URL.Query().Get("name")
	if name == "" {
		return InvalidQueryParams()
	}

	err := h.db.DeleteExercise(name)
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusNoContent, nil)
}
