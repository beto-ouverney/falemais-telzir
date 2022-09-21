package handler

import (
	"errors"
	"github.com/beto-ouverney/falemais-telzir/customerror"
	"reflect"
	"testing"
)

func Test_errorHandler(t *testing.T) {

	tests := []struct {
		describe     string
		err          *customerror.CustomError
		wantResponse []byte
		wantStatus   int
	}{
		{
			describe: "Should return a 404 status code and a json with the message",
			err: &customerror.CustomError{
				Code: customerror.ENOTFOUND,
				Err:  errors.New("not found"),
			},
			wantResponse: []byte("{\"message\":\"not found\"}"),
			wantStatus:   404,
		},
		{
			describe: "Should return a 400 status code and a json with the message",
			err: &customerror.CustomError{
				Code: customerror.ECONFLICT,
				Err:  errors.New("conflict"),
			},
			wantResponse: []byte("{\"message\":\"conflict\"}"),
			wantStatus:   400,
		},
		{
			describe: "Should return a 500 status code and a json with the message",
			err: &customerror.CustomError{
				Code: customerror.EINTERNAL,
				Err:  errors.New("internal"),
			},
			wantResponse: []byte("{\"message\":\"internal\"}"),
			wantStatus:   500,
		},
	}

	for _, tt := range tests {
		t.Run(tt.describe, func(t *testing.T) {
			gotResponse, gotStatus := errorHandler(tt.err)
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("errorHandler() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
			if gotStatus != tt.wantStatus {
				t.Errorf("errorHandler() gotStatus = %v, want %v", gotStatus, tt.wantStatus)
			}
		})
	}
}
