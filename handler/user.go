package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/joe-ngu/gogym/storage"
)

type UserPayload struct {
	FirstName         string      `json:"first_name"`
	LastName          string      `json:"last_name"`
	UserName          string      `json:"user_name"`
	EncryptedPassword string      `json:"encrypted_password"`
	CreatedAt         time.Time   `json:"created_at"`
	Workouts          []uuid.UUID `json:"workouts"`
}

func (p *UserPayload) validate() map[string]string {
	errs := make(map[string]string)
	return errs
}

type UserHandler struct {
	db storage.DB
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) error {
	log.Println("Handling CREATE request - Method:", r.Method)
}

func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) error {
	log.Println("Handling GET ALL request - Method:", r.Method)
}

func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) error {
	log.Println("Handling Get request - Method:", r.Method)
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) error {
	log.Println("Handling CREATE request - Method:", r.Method)
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) error {
	log.Println("Handling CREATE request - Method:", r.Method)
}
