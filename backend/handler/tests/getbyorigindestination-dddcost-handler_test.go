package handler_test

import (
	"bytes"
	"encoding/json"
	"github.com/beto-ouverney/falemais-telzir/entity"
	"github.com/beto-ouverney/falemais-telzir/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestGetCost(t *testing.T) {
	assertions := assert.New(t)

	t.Setenv("POSTGRES_USER", "root")
	t.Setenv("POSTGRES_PASSWORD", "password")
	t.Setenv("POSTGRES_DB", "telzir_db_test")
	t.Setenv("DB_CONNECTION", "user=root password=password dbname=telzir_db_test sslmode=disable")

	// initialize the database if in the test environment
	if strings.Contains(os.Getenv("POSTGRES_DB"), "test") {
		t.Log("Initializing the database for testing")
		initDBTest(t)
	} else {
		t.Skip("Skipping test because it is not a test environment, the port number is not 6306")
	}

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Route("/api/v1/dddcost", func(r chi.Router) {
		r.Get("/", handler.GetAllDDDCodes)
		r.Post("/", handler.GetCost)
	})

	jsonBodyMock := struct {
		Origin      int `json:"origin,omitempty"`
		Destination int `json:"destination,omitempty"`
		Minutes     int `json:"minutes,omitempty"`
	}{
		18,
		11,
		200,
	}

	tests := []struct {
		describe       string
		args           interface{}
		expectedStatus int
		expectedBody   entity.PlanComparation
		msgStatus      string
		msgBody        string
	}{
		{
			describe:       "Should be able to return the cost of a call with plans_comparations",
			args:           jsonBodyMock,
			expectedStatus: 200,
			expectedBody: entity.PlanComparation{
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
						With:    167.20,
						Without: 380.00,
					},
				},
			},
			msgStatus: "The status code should be 200",
			msgBody:   "Should be the same as the mock",
		},
	}
	for _, tt := range tests {
		t.Run(tt.describe, func(t *testing.T) {
			res := httptest.NewRecorder()
			bodyMock, errj := json.Marshal(&tt.args)
			if errj != nil {
				t.Fatal(errj)
			}

			req := httptest.NewRequest("POST", "/api/v1/dddcost", bytes.NewBuffer(bodyMock))
			router.ServeHTTP(res, req)

			assertions.Equal(tt.expectedStatus, res.Code, tt.msgStatus)
			var actual entity.PlanComparation
			err := json.Unmarshal(res.Body.Bytes(), &actual)
			if err != nil {
				t.Errorf("Error unmarshalling response: %v", err)
			}
			assertions.EqualValues(tt.expectedBody.Origin, actual.Origin, tt.msgBody)
			assertions.EqualValues(tt.expectedBody.Destination, actual.Destination, tt.msgBody)
			assertions.EqualValues(tt.expectedBody.Minutes, actual.Minutes, tt.msgBody)
			for _, plan := range tt.expectedBody.Plans {
				for _, pActual := range actual.Plans {
					if plan.Name == pActual.Name {
						assertions.EqualValues(plan.With, pActual.With, tt.msgBody)
						assertions.EqualValues(plan.Without, pActual.Without, tt.msgBody)
					}

				}

			}

		})
	}
}
