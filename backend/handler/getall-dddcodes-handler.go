package handler

import (
	"fmt"
	dddcost_controller "github.com/beto-ouverney/falemais-telzir/controller/dddcost-controller"
	"log"
	"net/http"
)

// GetAllDDDCodes is the handler for the route /dddcodes, return dddcodes avaliable
// @Summary      Get all DDD codes avaliable in database and return in JSON format
// @Description  Get all DDD codes avaliable in database and return in JSON format
// @Tags         dddcodes, ddd
// @Produce      json
// @Success      200  {string} json "{"version":"0.0.1"}"
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router / [get]
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
