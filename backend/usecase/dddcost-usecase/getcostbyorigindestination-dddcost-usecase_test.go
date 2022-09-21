package dddcost_usecase

import (
	"errors"
	"github.com/beto-ouverney/falemais-telzir/customerror"
	"github.com/beto-ouverney/falemais-telzir/entity"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func Test_calculateCost(t *testing.T) {
	assertions := assert.New(t)

	type args struct {
		planComparation *entity.PlanComparation
		planName        string
		freeMinutes     int
		costByMinute    float64
	}
	tests := []struct {
		describe string
		args     args
		want     *entity.PlanComparation
		msg      string
	}{
		{
			describe: "Should be able return a cost plan comparation",
			args: args{
				planComparation: &entity.PlanComparation{
					Origin:      11,
					Destination: 16,
					Minutes:     20,
					Plans:       []entity.Plan{},
				},
				planName:     "FaleMais 30",
				freeMinutes:  30,
				costByMinute: 1.90,
			},
			want: &entity.PlanComparation{
				Origin:      11,
				Destination: 16,
				Minutes:     20,
				Plans: []entity.Plan{
					{
						Name:    "FaleMais 30",
						With:    0.00,
						Without: 38.00,
					},
				},
			},
			msg: "Must be equal",
		},
		{
			describe: "Should be able return a cost plan comparation",
			args: args{
				planComparation: &entity.PlanComparation{
					Origin:      11,
					Destination: 17,
					Minutes:     80,
					Plans:       []entity.Plan{},
				},
				planName:     "FaleMais 60",
				freeMinutes:  60,
				costByMinute: 1.70,
			},
			want: &entity.PlanComparation{
				Origin:      11,
				Destination: 17,
				Minutes:     80,
				Plans: []entity.Plan{
					{
						Name:    "FaleMais 60",
						With:    37.40,
						Without: 136.00,
					},
				},
			},
			msg: "Must be equal",
		},
		{
			describe: "Should be able return a cost plan comparation",
			args: args{
				planComparation: &entity.PlanComparation{
					Origin:      18,
					Destination: 11,
					Minutes:     200,
					Plans:       []entity.Plan{},
				},
				planName:     "FaleMais 120",
				freeMinutes:  120,
				costByMinute: 1.90,
			},
			want: &entity.PlanComparation{
				Origin:      18,
				Destination: 11,
				Minutes:     200,
				Plans: []entity.Plan{
					{
						Name:    "FaleMais 120",
						With:    167.20,
						Without: 380.00,
					},
				},
			},
			msg: "Must be equal",
		},
	}
	for _, tt := range tests {
		t.Run(tt.describe, func(t *testing.T) {
			wg := sync.WaitGroup{}
			wg.Add(1)
			go calculateCost(tt.args.planComparation, tt.args.planName, tt.args.freeMinutes, tt.args.costByMinute, &wg)
			wg.Wait()
			assertions.Equal(tt.want, tt.args.planComparation, tt.msg)
		})
	}
}

func Test_validateFields(t *testing.T) {
	assertions := assert.New(t)

	type args struct {
		origin      *int
		destination *int
		min         *int
	}

	origin9 := 9
	origin11 := 11
	destination9 := 9
	destination16 := 16
	min0 := 0
	min20 := 20

	tests := []struct {
		describe string
		args     args
		want     *customerror.CustomError
		msg      string
	}{
		{
			describe: "Should be able return a custom error if origin is nil",
			args: args{
				origin:      nil,
				destination: &destination16,
				min:         &min20,
			},
			want: customerror.NewError(customerror.ECONFLICT, "Origin is required",
				"dddcost_usecase.GetCostByOriginDestination", errors.New("origin is required")),
			msg: "Must be equal",
		},
		{
			describe: "Should be able return a custom error if origin is not be greater than 10",
			args: args{
				origin:      &origin9,
				destination: &destination16,
				min:         &min20,
			},
			want: customerror.NewError(customerror.ECONFLICT, "Origin must be greater than 10",
				"dddcost_usecase.GetCostByOriginDestination", errors.New("origin must be greater than 10")),
			msg: "Must be equal",
		},
		{
			describe: "Should be able return a custom error if destination is nil",
			args: args{
				origin:      &origin11,
				destination: nil,
				min:         &min20,
			},
			want: customerror.NewError(customerror.ECONFLICT, "Destination is required",
				"dddcost_usecase.GetCostByOriginDestination", errors.New("destination is required")),
			msg: "Must be equal",
		},
		{
			describe: "Should be able return a custom error if destination is not be greater than 10",
			args: args{
				origin:      &origin11,
				destination: &destination9,
				min:         &min20,
			},
			want: customerror.NewError(customerror.ECONFLICT, "Destination must be greater than 10",
				"dddcost_usecase.GetCostByOriginDestination", errors.New("destination must be greater than 10")),
			msg: "Must be equal",
		},
		{
			describe: "Should be able return a custom error if minutes is nil",
			args: args{
				origin:      &origin11,
				destination: &destination16,
				min:         nil,
			},
			want: customerror.NewError(customerror.ECONFLICT, "Minutes is required",
				"dddcost_usecase.GetCostByOriginDestination", errors.New("minutes is required")),
			msg: "Must be equal",
		},
		{
			describe: "Should be able return a custom error if minutes is not be greater than 0",
			args: args{
				origin:      &origin11,
				destination: &destination16,
				min:         &min0,
			},
			want: customerror.NewError(customerror.ECONFLICT, "Minutes must be greater than 0",
				"dddcost_usecase.GetCostByOriginDestination", errors.New("minutes must be greater than 0")),
			msg: "Must be equal",
		},
	}
	for _, tt := range tests {
		t.Run(tt.describe, func(t *testing.T) {
			got := validateFields(tt.args.origin, tt.args.destination, tt.args.min)
			assertions.Equal(tt.want, got, tt.msg)
		})
	}
}
