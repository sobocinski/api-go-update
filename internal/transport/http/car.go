package http

import (
	"encoding/json"
	"github.com/sobocinski/api-go-update/internal/car/command"
	"github.com/sobocinski/api-go-update/internal/car/query"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// GetCar - retrieve a car by id
func (h *Handler) GetCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	i, _ := strconv.ParseUint(vars["id"], 10, 64)
	query := query.GetCarQuery{
		Id: uint(i),
	}
	user, err := h.carService.GetCar(query)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// CreateCar - create new car
func (h *Handler) CreateCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var command command.CreateCarCommand

	err := json.NewDecoder(r.Body).Decode(&command)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = h.carService.CreateCar(command)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}


//// GetUserCars - find user cars
//func (h *Handler) GetUserCars(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
//
//	var command command.CreateCarCommand
//
//	err := json.NewDecoder(r.Body).Decode(&command)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//	_, err = h.carService.CreateCar(command)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	w.WriteHeader(http.StatusCreated)
//}