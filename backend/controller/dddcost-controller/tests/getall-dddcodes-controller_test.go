package tests

import (
	"context"
	"encoding/json"
	"github.com/beto-ouverney/falemais-telzir/controller/dddcost-controller/mocks"
	"github.com/beto-ouverney/falemais-telzir/customerror"
	"github.com/beto-ouverney/falemais-telzir/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_dDDCostController_GetAllDDDCodes(t *testing.T) {
	assertions := assert.New(t)

	dddCodesMock := []entity.DDDCost{
		{
			Origin:      11,
			Destination: 16,
		},
		{
			Origin:      16,
			Destination: 11,
		},
	}

	jsonMock, errJ := json.MarshalIndent(dddCodesMock, "", "  ")
	if errJ != nil {
		t.Fatal("Error to marshal dddCodesMock: ", errJ)
	}

	tests := []struct {
		describe string
		want     *[]byte
		want1    *customerror.CustomError
		msg      string
		msg1     string
	}{
		{
			describe: "Should be able return a json with a list of dddcosts without a error",
			want:     &jsonMock,
			want1:    nil,
			msg:      "Must be Equal",
			msg1:     "Must be nil",
		},
	}
	for _, tt := range tests {
		t.Run(tt.describe, func(t *testing.T) {
			ctx := context.Background()
			m := new(mocks.IDDDCostController)
			m.On("GetAllDDDCodes", mock.AnythingOfType("*context.emptyCtx")).Return(tt.want, tt.want1)

			got, got1 := m.GetAllDDDCodes(ctx)

			assertions.Equal(tt.want, got, tt.msg)
			assertions.Equal(tt.want1, got1, tt.msg1)

		})
	}
}
