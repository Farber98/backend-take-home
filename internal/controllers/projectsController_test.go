package controllers

import (
	"cloudhumans/internal/services"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestAllocateErrors(t *testing.T) {

	projectService := &services.ProjectsService{}
	projectController := &ProjectsController{service: projectService}

	testCases := []struct {
		name             string
		body             string
		expectStatusCode int
		expectError      string
	}{
		{
			name:             "ERR_BINDING",
			body:             `NOT A JSON`,
			expectStatusCode: http.StatusInternalServerError,
			expectError:      ERR_BINDING,
		},
		{
			name:             "ERR_VALIDATE_EMPTY_AGE",
			body:             `{}`,
			expectStatusCode: http.StatusBadRequest,
			expectError:      services.ERR_AGE,
		},
		{
			name:             "ERR_VALIDATE_WRONG_AGE",
			body:             `{"age":0}`,
			expectStatusCode: http.StatusBadRequest,
			expectError:      services.ERR_AGE,
		},
		{
			name:             "ERR_VALIDATE_EMPTY_EDUCATION_LEVEL",
			body:             `{"age": 18}`,
			expectStatusCode: http.StatusBadRequest,
			expectError:      services.ERR_EDUCATION_LEVEL,
		},
		{
			name:             "ERR_VALIDATE_WRONG_EDUCATION_LEVEL",
			body:             `{"age": 18, "education_level": "kindergarden"}`,
			expectStatusCode: http.StatusBadRequest,
			expectError:      services.ERR_EDUCATION_LEVEL,
		},
		{
			name: "ERR_VALIDATE_EMPTY_PAST_EXPERIENECS",
			body: `{
						"age": 18, 
						"education_level": "no_education"
					}`,
			expectStatusCode: http.StatusBadRequest,
			expectError:      services.ERR_PAST_EXPERIENCES,
		},
		{
			name: "ERR_VALIDATE_EMPTY_INTERNET_TEST",
			body: `{
						"age": 18, 
						"education_level": "no_education",
						"past_experiences": {
							"sales": true,
							"support": false
						}
					}`,
			expectStatusCode: http.StatusBadRequest,
			expectError:      services.ERR_INTERNET_TEST,
		},
		{
			name: "ERR_VALIDATE_EMPTY_WRITING_SCORE",
			body: `{
						"age": 18, 
						"education_level": "no_education",
						"past_experiences": {
							"sales": true,
							"support": false
						},
						"internet_test": {
							"download_speed": 50.4,
							"upload_speed": 40.2
						}
					}`,
			expectStatusCode: http.StatusBadRequest,
			expectError:      services.ERR_WRITING_SCORE,
		},
		{
			name: "ERR_VALIDATE_WRONG_WRITING_SCORE",
			body: `{
						"age": 18, 
						"education_level": "no_education",
						"past_experiences": {
							"sales": true,
							"support": false
						},
						"internet_test": {
							"download_speed": 50.4,
							"upload_speed": 40.2
						},
						"writing_score": 1.1
					}`,
			expectStatusCode: http.StatusBadRequest,
			expectError:      services.ERR_WRITING_SCORE,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/projects/allocate", strings.NewReader(tc.body))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if assert.NoError(t, projectController.Allocate(c)) {
				newl := len(rec.Body.String()) - 2

				assert.Equal(t, tc.expectStatusCode, rec.Code)
				assert.Equal(t, tc.expectError, rec.Body.String()[1:newl])
			}
		})
	}
}
