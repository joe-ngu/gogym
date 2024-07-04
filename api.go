package main

import (
	"log"
	"net/http"

	"github.com/joe-ngu/gogym/handler"
	"github.com/joe-ngu/gogym/store"
)

type APIServer struct {
	listenAddr string
	store      store.DB
}

func NewAPIServer(listenAddr string, store store.DB) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {
	router := http.NewServeMux()

	exerciseHandler := handler.NewExerciseHandler(s.store)
	workoutHandler := handler.NewWorkoutHandler(s.store)
	userHandler := handler.NewUserHandler(s.store)
	authHandler := handler.NewAuthHandler(s.store)

	// authRoutes
	router.HandleFunc("POST /login", handler.Make(authHandler.HandleLogin))

	// User Routes
	router.HandleFunc("POST /user", handler.Make(userHandler.Create))
	router.HandleFunc("GET /users", handler.Make(userHandler.GetAll))
	router.HandleFunc("GET /user/{id}", handler.JWTAuthMiddleware(handler.Make(userHandler.Get), s.store))
	router.HandleFunc("PUT /user/{id}", handler.JWTAuthMiddleware(handler.Make(userHandler.Update), s.store))
	router.HandleFunc("DELETE /user/{id}", handler.JWTAuthMiddleware(handler.Make(userHandler.Delete), s.store))

	// Workout Routes
	router.HandleFunc("POST /workout", handler.Make(workoutHandler.Create))
	router.HandleFunc("GET /workouts", handler.Make(workoutHandler.GetAll))
	router.HandleFunc("GET /workout/{id}", handler.JWTAuthMiddleware(handler.Make(workoutHandler.Get), s.store))
	router.HandleFunc("PUT /workout/{id}", handler.JWTAuthMiddleware(handler.Make(workoutHandler.Update), s.store))
	router.HandleFunc("DELETE /workout/{id}", handler.JWTAuthMiddleware(handler.Make(workoutHandler.Delete), s.store))

	// Exercise Routes
	router.HandleFunc("POST /exercise", handler.Make(exerciseHandler.Create))
	router.HandleFunc("GET /exercises", handler.Make(exerciseHandler.GetAll))
	router.HandleFunc("GET /exercise/{name}", handler.Make(exerciseHandler.Get))
	router.HandleFunc("PUT /exercise/{name}", handler.Make(exerciseHandler.Update))
	router.HandleFunc("DELETE /exercise/{name}", handler.Make(exerciseHandler.Delete))

	log.Println("API server running on port: ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}
