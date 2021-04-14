package http

import (
	"encoding/json"
	"github.com/sobocinski/api-go-update/internal/car"
	"github.com/sobocinski/api-go-update/internal/user"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// Handler - stores pointer to our services
type Handler struct {
	Router  *mux.Router
	userService *user.Service
	carService *car.Service
}

// NewHandler - returns a pointer to a Handler
func NewHandler(us *user.Service, cs *car.Service) *Handler {
	return &Handler{
		userService: us,
		carService: cs,
	}
}

// SetupRoutes - sets up all the routes for our application
func (h *Handler) SetupRoutes() {
	log.Info("Setting Up Routes")
	h.Router = mux.NewRouter()
	h.Router.Use(LoggingMiddleware)

	h.Router.HandleFunc("/api/user/{id}", h.GetUser).Methods("GET")
	h.Router.HandleFunc("/api/user", h.CreateUser).Methods("POST")
	h.Router.HandleFunc("/api/user/{id}/cars", h.GetUserCars).Methods("GET")
	h.Router.HandleFunc("/api/user/{id}/cars2", h.GetUserWithCars).Methods("GET")

	h.Router.HandleFunc("/api/car/{id}", h.GetCar).Methods("GET")
	h.Router.HandleFunc("/api/car", h.CreateCar).Methods("POST")

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(struct{ Message string }{Message: "I am Alive!"}); err != nil {
			panic(err)
		}
	})
}

// LoggingMiddleware - a handy middleware function that logs out incoming requests
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(
			log.Fields{
				"Method": r.Method,
				"Path":   r.URL.Path,
			}).
			Info("handled request")
		next.ServeHTTP(w, r)
	})
}
