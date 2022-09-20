package dddcost_repository

import (
	"context"
	"github.com/beto-ouverney/falemais-telzir/customerror"
	"github.com/beto-ouverney/falemais-telzir/entity"
)

func (r *dDDCostRepository) GetAll(ctx context.Context) (*[]entity.DDDCost, *customerror.CustomError) {
	var dddcosts []entity.DDDCost
	err := r.sqlx.SelectContext(ctx, &dddcosts, "SELECT * FROM dddcost")
	if err != nil {
		return nil, customerror.NewError(customerror.EINTERNAL, "Internal error", "dddcost_repository.GetAll", err, nil)
	}
	return &dddcosts, nil
}
