package workout

import (
	"log"
	"net/http"
)


type Handler struct{}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusCreated)
  log.Println("Handling CREATE request - Method:", r.Method)
}  

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
  log.Println("Handling READ ALL request - Method:", r.Method)
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
  log.Println("Handling READ request - Method:", r.Method)
}

func (h *Handler) UpdateByID(w http.ResponseWriter, r *http.Request) {
  log.Println("Handling UPDATE request - Method:", r.Method)
}


func (h *Handler) DeleteByID(w http.ResponseWriter, r *http.Request) {
  log.Println("Handling DELETE request - Method:", r.Method)
}
