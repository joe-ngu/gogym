package main

import (
	"net/http"

	"github.com/joe-ngu/gogym/handler"
	"github.com/joe-ngu/gogym/storage"
)

type APIServer struct {
	listenAddr string
	store      storage.DB
}

func NewAPIServer(listenAddr string, store storage.DB) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {
	router := http.NewServeMux()

	exerciseHandler := &handler.ExerciseHandler{}
	workoutHandler := &handler.WorkoutHandler{}
	userHandler := &handler.UserHandler{}

	// User Routes
	router.HandleFunc("POST /exercise", handler.Make(exerciseHandler.Create))
	router.HandleFunc("GET /exercises", handler.Make(exerciseHandler.GetAll))
	router.HandleFunc("GET /exercise/{name}", handler.Make(exerciseHandler.Get))
	router.HandleFunc("PUT /exercise/{name}", handler.Make(exerciseHandler.Update))
	router.HandleFunc("DELETE /exercise/{name}", handler.Make(exerciseHandler.Delete))

	// Workout Routes
	router.HandleFunc("POST /workout", handler.Make(workoutHandler.Create))

	// Exercise Routes

}
