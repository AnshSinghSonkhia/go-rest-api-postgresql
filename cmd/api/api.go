package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	// create router
	router := mux.NewRouter()

	// register services

	// c := cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"http://localhost:5173"},
	// 	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
	// 	AllowedHeaders:   []string{"Content-Type", "Authorization"},
	// 	AllowCredentials: true,
	// })

	// corsHandler := c.Handler(router)

	log.Println("Listening on port", s.addr)

	return http.ListenAndServe(s.addr, router)
}
