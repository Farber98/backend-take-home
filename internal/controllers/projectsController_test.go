package controllers

import (
	"cloudhumans/internal/models"
	"cloudhumans/internal/services"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestAllocateErrors(t *testing.T) {

	projectService := &services.ProjectsService{}
	projectController := &ProjectsController{Service: projectService}

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
			name:             "ERR_VALIDATE_EMPTY_OBJECT",
			body:             `{}`,
			expectStatusCode: http.StatusBadRequest,
			expectError:      services.ERR_AGE,
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

				errorResponseModel := &models.MsgResponse{}
				err := json.Unmarshal(rec.Body.Bytes(), errorResponseModel)

				if assert.NoError(t, err) {
					assert.Equal(t, tc.expectStatusCode, rec.Code)
					assert.Equal(t, tc.expectError, errorResponseModel.Message)
				}
			}
		})
	}
}
func TestAllocateSuccess(t *testing.T) {

	projectService := &services.ProjectsService{}
	projectController := &ProjectsController{Service: projectService}

	testCases := []struct {
		name             string
		body             string
		expectStatusCode int
		out              *models.Response
	}{
		{
			name: "SUCCESS_EXAMPLE",
			body: `{
						"age": 35, 
						"education_level": "high_school",
						"past_experiences": {
							"sales": false,
							"support": true
						},
						"internet_test": {
							"download_speed": 50.4,
							"upload_speed": 40.2
						},
						"writing_score": 0.6,
						"referral_code": "token1234"
					}`,
			expectStatusCode: http.StatusOK,
			out: &models.Response{
				Score:           7,
				SelectedProject: services.PROJECT_SCHRODINGER,
				EligibleProjects: []string{
					services.PROJECT_SCHRODINGER,
					services.PROJECT_YXZ,
					services.PROJECT_XPTO,
				},
				InelegibleProjects: []string{
					services.PROJECT_NASA,
				},
			},
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

				realOutputModel := &models.Response{}
				realOutputUnmarshall := json.Unmarshal(rec.Body.Bytes(), realOutputModel)

				if assert.NoError(t, realOutputUnmarshall) {
					assert.Equal(t, tc.expectStatusCode, rec.Code)
					assert.Equal(t, tc.out, realOutputModel)
				}

			}
		})
	}
}
