package dddcost_usecase

import (
	"context"
	"errors"
	"github.com/beto-ouverney/falemais-telzir/customerror"
	"github.com/beto-ouverney/falemais-telzir/entity"
	"math"
	"sync"
)

// validateFields validates the fields
func validateFields(origin, destination, min *int) *customerror.CustomError {
	if origin == nil {
		return customerror.NewError(customerror.ECONFLICT, "Origin is required",
			"dddcost_usecase.GetCostByOriginDestination", errors.New("origin is required"), nil)
	}
	if *origin < 10 {
		return customerror.NewError(customerror.ECONFLICT, "Origin must be greater than 10",
			"dddcost_usecase.GetCostByOriginDestination", errors.New("origin must be greater than 10"), nil)
	}
	if destination == nil {
		return customerror.NewError(customerror.ECONFLICT, "Destination is required",
			"dddcost_usecase.GetCostByOriginDestination", errors.New("destination is required"), nil)
	}

	if *destination < 10 {
		return customerror.NewError(customerror.ECONFLICT, "Destination must be greater than 10",
			"dddcost_usecase.GetCostByOriginDestination", errors.New("destination must be greater than 10"), nil)
	}

	if min == nil {
		return customerror.NewError(customerror.ECONFLICT, "Minutes is required",
			"dddcost_usecase.GetCostByOriginDestination", errors.New("minutes is required"), nil)
	}

	if *min <= 0 {
		return customerror.NewError(customerror.ECONFLICT, "Minutes must be greater than 0",
			"dddcost_usecase.GetCostByOriginDestination", errors.New("minutes must be greater than 0"), nil)
	}
	return nil
}

// roundFloat is a function to round float
func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

// calculateCost calculates the cost of a call by minutes by plan
func calculateCost(planComparation *entity.PlanComparation, planName string, freeMinutes int, costByMinute float64) {
	plan := entity.Plan{}
	plan.Name = planName
	plan.With = 0
	plan.Without = roundFloat(costByMinute*float64(planComparation.Minutes), 2)

	if planComparation.Minutes > freeMinutes {
		with := float64(planComparation.Minutes-freeMinutes) * (costByMinute * 1.10)
		plan.With = roundFloat(with, 2)
	}
	planComparation.Plans = append(planComparation.Plans, plan)
}

// GetCostByOriginDestination gets the dddcost by origin and destination
func (u *dDDCostUseCase) GetCostByOriginDestination(ctx context.Context, origin, destination, min *int) (*entity.PlanComparation, *customerror.CustomError) {

	errV := validateFields(origin, destination, min)
	if errV != nil {
		return nil, errV
	}

	dddcost, err := u.dDDCostRepository.GetByOriginDestination(ctx, origin, destination)
	if err != nil {
		return nil, err
	}

	plansComparations := entity.PlanComparation{}
	plansComparations.Origin = *origin
	plansComparations.Destination = *destination
	plansComparations.Minutes = *min

	wg := sync.WaitGroup{}
	wg.Add(3)

	go calculateCost(&plansComparations, "FaleMais 30", 30, dddcost.Cost)
	go calculateCost(&plansComparations, "FaleMais 60", 60, dddcost.Cost)
	go calculateCost(&plansComparations, "FaleMais 120", 120, dddcost.Cost)
	wg.Wait()

	return &plansComparations, nil

}
