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

func Test_dDDCostRepository_GetByOriginDestination(t *testing.T) {
	assertions := assert.New(t)

	type args struct {
		origin      int
		destination int
	}

	tests := []struct {
		describe string
		args     args
		want     *entity.DDDCost
		want1    *customerror.CustomError
		msg      string
		msg1     string
	}{
		{
			describe: "Should be able return a dddcosts without a error",
			args: args{
				origin:      11,
				destination: 16,
			},
			want: &entity.DDDCost{
				ID:          1,
				Origin:      11,
				Destination: 16,
				Cost:        1.90,
			},
			want1: nil,
			msg:   "Should be able return a dddcosts without a error",
			msg1:  "Should be able return a dddcosts with a cost field",
		},
	}
	for _, tt := range tests {
		t.Run(tt.describe, func(t *testing.T) {
			ctx := context.Background()
			m := new(mocks.IDDDCostRepository)
			m.On("GetByOriginDestination", mock.AnythingOfType("*context.emptyCtx"), &tt.args.origin, &tt.args.destination).Return(tt.want, tt.want1)

			got, got1 := m.GetByOriginDestination(ctx, &tt.args.origin, &tt.args.destination)
			assertions.Equal(tt.want, got, tt.msg)
			assertions.Equal(tt.want1, got1, tt.msg1)
		})
	}
}
