package dddcost_repository

import (
	"context"
	"github.com/beto-ouverney/falemais-telzir/customerror"
	"github.com/beto-ouverney/falemais-telzir/db"
	"github.com/beto-ouverney/falemais-telzir/entity"
	"github.com/jmoiron/sqlx"
)

// IDDDCostRepository presents the interface for the dddcost repository
type IDDDCostRepository interface {
	GetAllDDDCodes(ctx context.Context) (*[]entity.DDDCost, *customerror.CustomError)
	GetAllDDDCosts(ctx context.Context) (*[]entity.DDDCost, *customerror.CustomError)
	GetByOriginDestination(ctx context.Context, origin, destination int) (*entity.DDDCost, *customerror.CustomError)
}

type dDDCostRepository struct {
	sqlx *sqlx.DB
}

// New creates a new dddcost repository
func New() *dDDCostRepository {
	return &dDDCostRepository{
		db.ConnectDB(),
	}
}
