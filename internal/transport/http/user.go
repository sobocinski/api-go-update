package http

import (
	"encoding/json"
	"fmt"
	"github.com/sobocinski/api-go-update/internal/user/command"
	"github.com/sobocinski/api-go-update/internal/user/query"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// GetUser - retrieve a user by id
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	i, _ := strconv.ParseUint(vars["id"], 10, 64)
	query := query.GetUserQuery{
		Id: uint(i),
	}
	user, err := h.userService.GetUser(query)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// CreateUser - create new user
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var command command.CreateUserCommand
	err := json.NewDecoder(r.Body).Decode(&command)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.userService.CreateUser(command)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetUserCars - create new user
func (h *Handler) GetUserCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	i, _ := strconv.ParseUint(vars["id"], 10, 64)
	query := query.GetUserCars{
		UserId: uint(i),
	}

	cars, err := h.carService.GetUserCars(query)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cars)
}

// GetUserCars - create new user
func (h *Handler) GetUserWithCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	i, _ := strconv.ParseUint(vars["id"], 10, 64)

	err := h.userService.UserWithCars(uint(i))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "user cars..")
}