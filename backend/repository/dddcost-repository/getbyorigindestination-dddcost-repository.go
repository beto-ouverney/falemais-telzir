package dddcost_repository

import (
	"context"
	"github.com/beto-ouverney/falemais-telzir/customerror"
	"github.com/beto-ouverney/falemais-telzir/entity"
)

// GetByOriginDestination gets the cost of a call by origin and destination
func (r *dDDCostRepository) GetByOriginDestination(ctx context.Context, origin, destination string) (*entity.DDDCost, *customerror.CustomError) {
	var dddcost entity.DDDCost
	err := r.sqlx.GetContext(ctx, &dddcost, "SELECT origin, destination, cost FROM dddcost WHERE origin = ? AND destination = ?", origin, destination)
	if err != nil {
		return nil, customerror.NewError(customerror.EINTERNAL, "Internal error", "dddcost_repository.GetByOriginDestination", err, nil)
	}
	return &dddcost, nil
}
