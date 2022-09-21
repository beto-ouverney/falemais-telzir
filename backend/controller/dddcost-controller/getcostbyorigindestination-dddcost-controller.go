package dddcost_controller

import (
	"context"
	"encoding/json"
	"github.com/beto-ouverney/falemais-telzir/customerror"
)

// GetCostByOriginDestination gets all codes available make a json and return to handler
func (c *dDDCostController) GetCostByOriginDestination(ctx context.Context, origin, destination, min *int) (*[]byte, *customerror.CustomError) {
	cost, err := c.u.GetCostByOriginDestination(ctx, origin, destination, min)
	if err != nil {
		return nil, err
	}

	resJson, errJ := json.MarshalIndent(&cost, "", "  ")
	if errJ != nil {
		return nil, customerror.NewError(customerror.EINTERNAL, "Json marshal error",
			"getcostbyorigindestination-dddcost-controller.GetByID", errJ)
	}
	return &resJson, nil
}
