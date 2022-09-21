package dddcost_controller

import (
	"context"
	"github.com/beto-ouverney/falemais-telzir/customerror"
	dddcost_usecase "github.com/beto-ouverney/falemais-telzir/usecase/dddcost-usecase"
)

// IDDDCostController presents the interface for the dddcost controller
type IDDDCostController interface {
	GetAllDDDCodes(ctx context.Context) (*[]byte, *customerror.CustomError)
	GetCostByOriginDestination(ctx context.Context, origin, destination, min *int) (*[]byte, *customerror.CustomError)
}

// dDDCostController presents the controller for the dddcost
type dDDCostController struct {
	u dddcost_usecase.IDDDCostUseCase
}

// New creates a new dddcost controller
func New() IDDDCostController {
	return &dDDCostController{
		dddcost_usecase.New(),
	}
}
