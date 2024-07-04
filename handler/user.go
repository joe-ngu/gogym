package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/joe-ngu/gogym/store"
	"github.com/joe-ngu/gogym/types"
)

type UserHandler struct {
	db store.DB
}

func NewUserHandler(db store.DB) *UserHandler {
	return &UserHandler{db: db}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) error {
	log.Println("Handling CREATE request - Method:", r.Method)
	var req types.UserPayload
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return InvalidJSON()
	}
	defer r.Body.Close()

	if errs := req.Validate(); len(errs) > 0 {
		return InvalidRequestData(errs)
	}

	user, err := types.NewUser(req.FirstName, req.LastName, req.UserName, req.Password)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusCreated, user)
}

func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) error {
	log.Println("Handling GET ALL request - Method:", r.Method)

	users, err := h.db.GetUsers()
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, users)
}

func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) error {
	log.Println("Handling Get request - Method:", r.Method)
	id := r.PathValue("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	user, err := h.db.GetUserByID(userID)
	if err != nil {
		return err
	}
	if user == nil {
		return writeJSON(w, http.StatusNotFound, user)
	}
	return writeJSON(w, http.StatusOK, user)
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) error {
	log.Println("Handling CREATE request - Method:", r.Method)
	id := r.PathValue("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	exists, _ := h.db.GetUserByID(userID)
	if exists == nil {
		return writeJSON(w, http.StatusNotFound, nil)
	}

	var req types.UserPayload
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return InvalidJSON()
	}
	defer r.Body.Close()

	if errs := req.Validate(); len(errs) > 0 {
		return InvalidRequestData(errs)
	}
	user, err := types.NewUser(req.FirstName, req.LastName, req.UserName, req.Password)
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, user)
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) error {
	log.Println("Handling CREATE request - Method:", r.Method)
	id := r.PathValue("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	if err := h.db.DeleteUser(userID); err != nil {
		return err
	}
	return writeJSON(w, http.StatusNoContent, nil)
}
