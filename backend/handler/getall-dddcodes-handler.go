package handler

import (
	"fmt"
	dddcost_controller "github.com/beto-ouverney/falemais-telzir/controller/dddcost-controller"
	"log"
	"net/http"
)

// GetAllDDDCodes is the handler for the route /dddcodes, return dddcodes avalible
func GetAllDDDCodes(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	dddCodes, err := dddcost_controller.New().GetAllDDDCodes(r.Context())
	response := *dddCodes
	status := http.StatusOK
	if err != nil {
		response, status = errorHandler(err)
	}
	fmt.Println(response)

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	_, errW := w.Write(response)
	if errW != nil {
		log.Printf("Write failed: %v", err)
	}
}
