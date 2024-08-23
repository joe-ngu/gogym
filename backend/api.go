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
	router.HandleFunc("GET /user", handler.JWTAuthMiddleware(handler.Make(userHandler.Get), s.store))
	router.HandleFunc("PUT /user", handler.JWTAuthMiddleware(handler.Make(userHandler.Update), s.store))
	router.HandleFunc("DELETE /user", handler.JWTAuthMiddleware(handler.Make(userHandler.Delete), s.store))

	// Workout Routes
	router.HandleFunc("POST /workout", handler.Make(workoutHandler.Create))
	router.HandleFunc("GET /workout", handler.JWTAuthMiddleware(handler.Make(workoutHandler.Get), s.store))
	router.HandleFunc("PUT /workout", handler.JWTAuthMiddleware(handler.Make(workoutHandler.Update), s.store))
	router.HandleFunc("DELETE /workout", handler.JWTAuthMiddleware(handler.Make(workoutHandler.Delete), s.store))

	// Exercise Routes
	router.HandleFunc("POST /exercise", handler.JWTAuthMiddleware(handler.Make(exerciseHandler.Create), s.store))
	router.HandleFunc("GET /exercise", handler.JWTAuthMiddleware(handler.Make(exerciseHandler.Get), s.store))
	router.HandleFunc("PUT /exercise", handler.JWTAuthMiddleware(handler.Make(exerciseHandler.Update), s.store))
	router.HandleFunc("DELETE /exercise", handler.JWTAuthMiddleware(handler.Make(exerciseHandler.Delete), s.store))

	log.Println("API server running on port: ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}
