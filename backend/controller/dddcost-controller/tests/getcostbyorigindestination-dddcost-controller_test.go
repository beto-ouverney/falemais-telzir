package tests

import (
	"context"
	"encoding/json"
	"github.com/beto-ouverney/falemais-telzir/controller/dddcost-controller/mocks"
	"github.com/beto-ouverney/falemais-telzir/customerror"
	"github.com/beto-ouverney/falemais-telzir/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_dDDCostController_GetCostByOriginDestination(t *testing.T) {
	assertions := assert.New(t)

	plansComparationsMock := entity.PlanComparation{
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
	}

	jsonMock, errJ := json.MarshalIndent(&plansComparationsMock, "", "  ")
	if errJ != nil {
		t.Errorf("Error to marshal plansComparations: %v", errJ)
	}

	type args struct {
		origin      int
		destination int
		min         int
	}

	tests := []struct {
		describe string
		args     args
		want     *[]byte
		want1    *customerror.CustomError
		msg      string
		msg1     string
	}{
		{
			describe: "Should be able return a plans comparation without a error",
			args: args{
				origin:      11,
				destination: 16,
				min:         20,
			},
			want:  &jsonMock,
			want1: nil,
			msg:   "Must be equal",
			msg1:  "Must be nil",
		},
	}
	for _, tt := range tests {
		t.Run(tt.describe, func(t *testing.T) {
			ctx := context.Background()
			m := new(mocks.IDDDCostController)
			m.On("GetCostByOriginDestination", ctx, &tt.args.origin, &tt.args.destination, &tt.args.min).Return(tt.want, tt.want1)

			got, got1 := m.GetCostByOriginDestination(ctx, &tt.args.origin, &tt.args.destination, &tt.args.min)

			assertions.Equal(tt.want, got, tt.msg)
			assertions.Equal(tt.want1, got1, tt.msg1)

		})
	}
}
