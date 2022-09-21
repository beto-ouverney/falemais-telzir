package dddcost_usecase

import (
	"context"
	"github.com/beto-ouverney/falemais-telzir/customerror"
	"github.com/beto-ouverney/falemais-telzir/entity"
)

// GetAllDDDCodes gets all codes available in the database
func (u *dDDCostUseCase) GetAllDDDCodes(ctx context.Context) (*[]entity.DDDCost, *customerror.CustomError) {
	return u.dDDCostRepository.GetAllDDDCodes(ctx)
}
