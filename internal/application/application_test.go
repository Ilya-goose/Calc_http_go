package application_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Ilya-goose/Calc_http_go/internal/application"
)

func TestCalcHandlerBadCase(t *testing.T) {
	testCases := []struct {
		name           string
		expression     string
		expectedBody   string
		expectedStatus int
	}{
		{
			name:           "Empty JSON",
			expression:     ``,
			expectedBody:   `{"error":"internal server error"}`,
			expectedStatus: 500,
		},
		{
			name:           "No expression field",
			expression:     `{"sometext": "2+2"}`,
			expectedBody:   `{"error":"expression is not valid"}`,
			expectedStatus: 422,
		},
		{
			name:           "Empty expression",
			expression:     `{"expression": ""}`,
			expectedBody:   `{"error":"expression is not valid"}`,
			expectedStatus: 422,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewBuffer([]byte(tc.expression)))
			w := httptest.NewRecorder()
			application.CalcHandler(w, req)
			res := w.Result()

			var responceBody map[string]interface{}
			json.NewDecoder(res.Body).Decode(&responceBody)
			body, _ := json.Marshal(responceBody)

			if tc.expectedBody != string(body) {
				t.Errorf("expected body %s, got %s", tc.expectedBody, body)
			}

			if res.StatusCode != tc.expectedStatus {
				t.Errorf("expected status %d, got %d", tc.expectedStatus, res.StatusCode)
			}
		})
	}
}

func TestCalcHandlerSuccessCase(t *testing.T) {
	testCases := []struct {
		name           string
		expression     string
		expectedBody   string
		expectedStatus int
	}{
		{
			name:           "First simple expression",
			expression:     `{"expression": "2+2*2"}`,
			expectedBody:   `{"result":6}`,
			expectedStatus: 200,
		},
		{
			name:           "Second simple expression",
			expression:     `{"expression": "1-2+3*4/5"}`,
			expectedBody:   `{"result":1.4}`,
			expectedStatus: 200,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewBuffer([]byte(tc.expression)))
			w := httptest.NewRecorder()
			application.CalcHandler(w, req)
			res := w.Result()

			var responceBody map[string]interface{}
			json.NewDecoder(res.Body).Decode(&responceBody)
			body, _ := json.Marshal(responceBody)

			if tc.expectedBody != string(body) {
				t.Errorf("expected body %s, got %s", tc.expectedBody, body)
			}

			if res.StatusCode != tc.expectedStatus {
				t.Errorf("expected status %d, got %d", tc.expectedStatus, res.StatusCode)
			}
		})
	}
}
