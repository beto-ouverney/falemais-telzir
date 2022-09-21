package handler

import (
	"encoding/json"
	dddcostcontroller "github.com/beto-ouverney/falemais-telzir/controller/dddcost-controller"
	"log"
	"net/http"
)

// GetCost is the handler for the route /dddcost/planscost, return a plans cost comparations
func GetCost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	response := []byte{}
	status := http.StatusOK

	var jsonBody struct {
		Origin      int `json:"origin"`
		Destination int `json:"destination"`
		Minutes     int `json:"minutes"`
	}

	errJ := json.NewDecoder(r.Body).Decode(&jsonBody)
	if errJ != nil {
		response = []byte(errJ.Error())
		status = http.StatusInternalServerError
	}

	cost, err := dddcostcontroller.New().GetCostByOriginDestination(r.Context(), &jsonBody.Origin, &jsonBody.Destination, &jsonBody.Minutes)
	response = *cost
	if err != nil {
		response, status = errorHandler(err)
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	_, errW := w.Write(response)
	if errW != nil {
		log.Printf("Write failed: %v", err)
	}

}
