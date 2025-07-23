package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/AnshSinghSonkhia/go-rest-api-postgresql/services/users"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

// CORS middleware
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next.ServeHTTP(w, r)
	})
}

func (s *APIServer) Run() error {
	// create router
	router := mux.NewRouter()

	// register services
	userStore := users.NewUserStore(s.db)
	userHandler := users.NewHandler(userStore)
	userHandler.RegisterRoutes(router)

	// apply CORS middleware
	handler := enableCORS(router)

	log.Println("Listening on port", s.addr)

	return http.ListenAndServe(s.addr, handler)
}
