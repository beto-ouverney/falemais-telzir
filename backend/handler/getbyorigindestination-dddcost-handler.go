package handler

import (
	"encoding/json"
	dddcostcontroller "github.com/beto-ouverney/falemais-telzir/controller/dddcost-controller"
	"log"
	"net/http"
)

// GetCost is the handler for the route /dddcost/planscost, return a plans cost comparations
// @Summary      Show plans cost comparations
// @Description  get plan cost comparations by origin and destination DDD
// @Tags         plancost
// @Accept       json
// @Produce      json
// @Param        origin   path      int  true  "Account ID"
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
func GetCost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	response := []byte{}
	status := http.StatusOK

	var jsonBody struct {
		Origin      int `json:"origin,omitempty"`
		Destination int `json:"destination,omitempty"`
		Minutes     int `json:"minutes,omitempty"`
	}

	errJ := json.NewDecoder(r.Body).Decode(&jsonBody)
	if errJ != nil {
		response = []byte(errJ.Error())
		status = http.StatusInternalServerError
	}

	cost, err := dddcostcontroller.New().GetCostByOriginDestination(r.Context(), &jsonBody.Origin, &jsonBody.Destination, &jsonBody.Minutes)

	if err != nil {
		response, status = errorHandler(err)
	} else {
		response = *cost
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	_, errW := w.Write(response)
	if errW != nil {
		log.Printf("Write failed: %v", err)
	}

}
