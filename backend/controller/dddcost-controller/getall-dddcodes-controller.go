package dddcost_controller

import (
	"context"
	"encoding/json"
	"github.com/beto-ouverney/falemais-telzir/customerror"
)

// GetAllDDDCodes gets all codes available make a json and return to handler
func (c *dDDCostController) GetAllDDDCodes(ctx context.Context) (*[]byte, *customerror.CustomError) {
	dDDCodes, err := c.u.GetAllDDDCodes(ctx)
	if err != nil {
		return nil, err
	}

	resJson, errJ := json.MarshalIndent(&dDDCodes, "", "  ")
	if errJ != nil {
		return nil, customerror.NewError(customerror.EINTERNAL, "Json marshal error",
			"getall-dddcodes-controller.GetByID", errJ)
	}
	return &resJson, nil
}
