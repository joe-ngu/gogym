package main

import (
	"net/http"

	"github.com/joe-ngu/gogym/handler"
	"github.com/joe-ngu/gogym/store"
)

type APIServer struct {
	*http.Server
	store store.DB
}

func NewAPIServer(addr string, store store.DB) *APIServer {
	router := http.NewServeMux()
	return &APIServer{
		Server: &http.Server{Addr: addr, Handler: router},
		store:  store,
	}
}

func (s *APIServer) setupRoutes() {
	router := s.Server.Handler.(*http.ServeMux)
	userRouter := http.NewServeMux()

	exerciseHandler := handler.NewExerciseHandler(s.store)
	workoutHandler := handler.NewWorkoutHandler(s.store)
	userHandler := handler.NewUserHandler(s.store)
	authHandler := handler.NewAuthHandler(s.store)

	// authRoutes
	router.HandleFunc("POST /login", handler.Make(authHandler.HandleLogin))
	router.HandleFunc("POST /signup", handler.Make(userHandler.Create))

	// User Routes
	userRouter.HandleFunc("GET /users", handler.Make(userHandler.GetAll))
	userRouter.HandleFunc("GET /user", handler.Make(userHandler.Get))
	userRouter.HandleFunc("PUT /user", handler.Make(userHandler.Update))
	userRouter.HandleFunc("DELETE /user", handler.Make(userHandler.Delete))

	// Workout Routes
	userRouter.HandleFunc("POST /workout", handler.Make(workoutHandler.Create))
	userRouter.HandleFunc("GET /workout", handler.Make(workoutHandler.Get))
	userRouter.HandleFunc("PUT /workout", handler.Make(workoutHandler.Update))
	userRouter.HandleFunc("DELETE /workout", handler.Make(workoutHandler.Delete))

	// Exercise Routes
	userRouter.HandleFunc("POST /exercise", handler.Make(exerciseHandler.Create))
	userRouter.HandleFunc("GET /exercise", handler.Make(exerciseHandler.Get))
	userRouter.HandleFunc("PUT /exercise", handler.Make(exerciseHandler.Update))
	userRouter.HandleFunc("DELETE /exercise", handler.Make(exerciseHandler.Delete))

	userRouterWithMiddleware := handler.JWTAuthMiddlewareFactory(s.store)(userRouter)
	router.Handle("/user", userRouterWithMiddleware)
	router.Handle("/users", userRouterWithMiddleware)
	router.Handle("/workout", userRouterWithMiddleware)
	router.Handle("/exercise", userRouterWithMiddleware)

	routerWithCors := handler.CorsMiddleware(router)
	s.Server.Handler = routerWithCors
}
