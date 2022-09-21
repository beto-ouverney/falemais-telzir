package dddcost_repository

import (
	"context"
	"github.com/beto-ouverney/falemais-telzir/customerror"
	"github.com/beto-ouverney/falemais-telzir/entity"
)

// GetByOriginDestination gets the cost of a call by origin and destination
func (r *dDDCostRepository) GetByOriginDestination(ctx context.Context, origin, destination *int) (*entity.DDDCost, *customerror.CustomError) {
	var dddcost entity.DDDCost

	err := r.sqlx.GetContext(ctx, &dddcost, "SELECT origin, destination, cost FROM dddcost WHERE origin = $1 AND destination = $2", *origin, *destination)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, customerror.NewError(customerror.ENOTFOUND, "Not found", "dddcost_repository.GetByOriginDestination", err)
		}
		return nil, customerror.NewError(customerror.EINVALID, "Internal error", "dddcost_repository.GetByOriginDestination", err)
	}
	return &dddcost, nil
}
