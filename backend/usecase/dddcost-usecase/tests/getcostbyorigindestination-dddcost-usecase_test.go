package tests

import (
	"context"
	"github.com/beto-ouverney/falemais-telzir/customerror"
	"github.com/beto-ouverney/falemais-telzir/entity"
	"github.com/beto-ouverney/falemais-telzir/usecase/dddcost-usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_GetCostByOriginDestination(t *testing.T) {
	assertions := assert.New(t)

	type args struct {
		origin      int
		destination int
		minutes     int
	}

	tests := []struct {
		describe string
		args     args
		want     *entity.PlanComparation
		want1    *customerror.CustomError
		msg      string
		msg1     string
	}{
		{
			describe: "Should be able return a plans comparation without a error",
			args: args{
				origin:      11,
				destination: 16,
				minutes:     20,
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
					{
						Name:    "FaleMais 60",
						With:    0.00,
						Without: 38.00,
					},
					{
						Name:    "FaleMais 120",
						With:    0.00,
						Without: 38.00,
					},
				},
			},
			want1: nil,
			msg:   "Must be equal",
			msg1:  "Must be nil",
		},
		{
			describe: "Should be able return a plans comparation without a error",
			args: args{
				origin:      18,
				destination: 11,
				minutes:     200,
			},
			want: &entity.PlanComparation{
				Origin:      18,
				Destination: 11,
				Minutes:     200,
				Plans: []entity.Plan{
					{
						Name:    "FaleMais 30",
						With:    355.30,
						Without: 380.00,
					},
					{
						Name:    "FaleMais 60",
						With:    292.60,
						Without: 380.00,
					},
					{
						Name:    "FaleMais 120",
						With:    0.00,
						Without: 380.00,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.describe, func(t *testing.T) {
			ctx := context.Background()
			m := new(mocks.IDDDCostUseCase)
			m.On("GetCostByOriginDestination", mock.AnythingOfType("*context.emptyCtx"), &tt.args.origin, &tt.args.destination, &tt.args.minutes).Return(tt.want, tt.want1)

			got, got1 := m.GetCostByOriginDestination(ctx, &tt.args.origin, &tt.args.destination, &tt.args.minutes)
			assertions.Equal(tt.want, got, tt.msg)
			assertions.Equal(tt.want1, got1, tt.msg1)
		})
	}
}
