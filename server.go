package main

//import "net/http"

type APIServer struct {
	listenAddr string
	// store Storage
}

func NewAPIServer(listenAddr string,// store Storage
) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		//store:      Store,
	}
}

func (s *APIServer) Run() {
	//router := http.NewServeMux()

	// User Routes

	// Workout Routes

	// Exercise Routes

}
