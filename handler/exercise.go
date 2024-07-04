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
	log.Println("Handling CREATE request - Method:", r.Method)
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
		return errors.New("Exercise already exists")
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

func (h *ExerciseHandler) GetAll(w http.ResponseWriter, r *http.Request) error {
	log.Println("Handling READ ALL request - Method:", r.Method)

	exercises, err := h.db.GetExercises()
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, exercises)
}

func (h *ExerciseHandler) Get(w http.ResponseWriter, r *http.Request) error {
	log.Println("Handling READ request - Method:", r.Method)
	name := r.PathValue("name")
	exercise, err := h.db.GetExercise(name)
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, exercise)
}

func (h *ExerciseHandler) Update(w http.ResponseWriter, r *http.Request) error {
	log.Println("Handling UPDATE request - Method:", r.Method)
	var req types.ExercisePayload

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return InvalidJSON()
	}
	defer r.Body.Close()

	if errs := req.Validate(); len(errs) > 0 {
		return InvalidRequestData(errs)
	}

	exists, _ := h.db.GetExercise(req.Name)
	if exists == nil {
		return errors.New("Exercise to update does not exist")
	}

	exercise, err := types.NewExercise(req.Name, req.MuscleGroup)
	if err != nil {
		return err
	}

	updated_exercise, err := h.db.UpdateExercise(exercise)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, updated_exercise)
}

func (h *ExerciseHandler) Delete(w http.ResponseWriter, r *http.Request) error {
	log.Println("Handling DELETE request - Method:", r.Method)
	name := r.PathValue("name")
	err := h.db.DeleteExercise(name)
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusNoContent, nil)
}
