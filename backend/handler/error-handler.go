package handler

import (
	"github.com/beto-ouverney/falemais-telzir/customerror"
	"log"
)

// errorHandler handles the error and return a json to the client
func errorHandler(err *customerror.CustomError) (response []byte, status int) {

	if err.Code == customerror.ENOTFOUND {
		status = 404
	} else if err.Code == customerror.ECONFLICT {
		status = 400
	} else {
		log.Println(err)
		status = 500
	}
	response = []byte("{\"message\":\"" + err.Error() + "\"}")

	return
}
