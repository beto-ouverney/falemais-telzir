package dddcost_usecase

import (
	"context"
	"github.com/beto-ouverney/falemais-telzir/customerror"
	"github.com/beto-ouverney/falemais-telzir/entity"
	dddcost_repository "github.com/beto-ouverney/falemais-telzir/repository/dddcost-repository"
)

// IDDDCostUseCase presents the interface for the dddcost usecase
type IDDDCostUseCase interface {
	GetAllDDDCodes(ctx context.Context) (*[]entity.DDDCost, *customerror.CustomError)
	GetCostByOriginDestination(ctx context.Context, origin, destination, min *int) (*entity.PlanComparation, *customerror.CustomError)
}

// dDDCostUseCase presents the usecase for the dddcost
type dDDCostUseCase struct {
	dDDCostRepository dddcost_repository.IDDDCostRepository
}

// New creates a new dddcost usecase
func New() IDDDCostUseCase {
	return &dDDCostUseCase{
		dddcost_repository.New(),
	}
}
