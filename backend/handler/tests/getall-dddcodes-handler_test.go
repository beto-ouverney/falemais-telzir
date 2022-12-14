package handler_test

import (
	"encoding/json"
	"github.com/beto-ouverney/falemais-telzir/entity"
	"github.com/beto-ouverney/falemais-telzir/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestGetAllDDDCodes(t *testing.T) {
	assertions := assert.New(t)
	const DATABASE_URL = "postgres://root:password@telzir_db_test:5432/telzir_db_test?sslmode=disable"
	t.Setenv("POSTGRES_USER", "root")
	t.Setenv("POSTGRES_PASSWORD", "password")
	t.Setenv("POSTGRES_DB", "telzir_db_test")
	t.Setenv("DB_CONNECTION", "postgres://root:password@telzir_db_test:5432/telzir_db_test?sslmode=disable")

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
		r.Get("/planscost", handler.GetCost)
	})

	tests := []struct {
		describe        string
		expectedStatus  int
		expectedMessage interface{}
		msg             string
		msg1            string
	}{
		{
			describe:        "Should be able to return all ddd codes",
			expectedStatus:  200,
			expectedMessage: allDDDCodesMock,
			msg:             "Status code must be 200",
			msg1:            "Must be equal",
		},
	}
	for _, tt := range tests {
		t.Run(tt.describe, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/api/v1/dddcost", nil)
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			assertions.Equal(tt.expectedStatus, rr.Code, tt.msg)
			var actual []entity.DDDCost
			err := json.Unmarshal(rr.Body.Bytes(), &actual)
			if err != nil {
				t.Errorf("Error unmarshalling response: %v", err)
			}
			t.Log(actual)
			t.Log(tt.expectedMessage)
			assertions.Equal(tt.expectedMessage, actual, tt.msg1)
		})
	}
}
