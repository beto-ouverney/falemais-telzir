package tests

import (
	"context"
	"github.com/beto-ouverney/falemais-telzir/customerror"
	"github.com/beto-ouverney/falemais-telzir/entity"
	"github.com/beto-ouverney/falemais-telzir/repository/dddcost-repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_dDDCostRepository_GetAllDDDCosts(t *testing.T) {
	assertions := assert.New(t)

	tests := []struct {
		describe string
		want     *[]entity.DDDCost
		want1    *customerror.CustomError
		msg      string
		msg1     string
	}{
		{
			describe: "Should be able return a list of dddcosts in order by origin without a error",
			want: &[]entity.DDDCost{
				{
					ID:          1,
					Origin:      11,
					Destination: 16,
					Cost:        3.50,
				},
				{
					ID:          2,
					Origin:      16,
					Destination: 11,
					Cost:        3.90,
				},
			},
			want1: nil,
			msg:   "Should be able return a list of dddcosts",
			msg1:  "Should be able return a list of dddcosts without a error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.describe, func(t *testing.T) {
			ctx := context.Background()
			m := new(mocks.IDDDCostRepository)
			m.On("GetAllDDDCodes", mock.AnythingOfType("*context.emptyCtx")).Return(tt.want, tt.want1)

			got, got1 := m.GetAllDDDCodes(ctx)
			assertions.Equal(tt.want, got, tt.msg)
			assertions.Equal(tt.want1, got1, tt.msg1)
		})
	}
}
