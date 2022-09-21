package dddcost_repository

import (
	"context"
	"github.com/beto-ouverney/falemais-telzir/customerror"
	"github.com/beto-ouverney/falemais-telzir/entity"
)

// GetAllDDDCodes gets all codes available in the database
func (r *dDDCostRepository) GetAllDDDCosts(ctx context.Context) (*[]entity.DDDCost, *customerror.CustomError) {
	var dddcosts []entity.DDDCost
	err := r.sqlx.SelectContext(ctx, &dddcosts, "SELECT * FROM dddcost ORDER BY origin ASC")
	if err != nil {
		return nil, customerror.NewError(customerror.EINVALID, "Internal error", "dddcost_repository.GetAllDDDCosts", err)
	}
	return &dddcosts, nil
}
