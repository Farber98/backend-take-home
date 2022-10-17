package services

import (
	"cloudhumans/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	testCases := []struct {
		name           string
		req            *models.Request
		expectedOutput string
	}{
		{
			name:           "ERR_VALIDATE_EMPTY_OR_WRONG_AGE",
			req:            &models.Request{Age: 0},
			expectedOutput: ERR_AGE,
		},
		{
			name:           "ERR_VALIDATE_EMPTY_OR_WRONG_EDUCATION_LEVEL",
			req:            &models.Request{Age: 18, EducationLevel: "kindergarden"},
			expectedOutput: ERR_EDUCATION_LEVEL,
		},
		{
			name:           "ERR_VALIDATE_EMPTY_PAST_EXPERIENECS",
			req:            &models.Request{Age: 18, EducationLevel: "high_school", PastExperiences: nil},
			expectedOutput: ERR_PAST_EXPERIENCES,
		},
		{
			name: "ERR_VALIDATE_EMPTY_INTERNET_TEST",
			req: &models.Request{
				Age: 18, EducationLevel: "high_school",
				PastExperiences: &models.PastExperiences{
					Sales:   true,
					Support: false,
				},
			},
			expectedOutput: ERR_INTERNET_TEST,
		},
		{
			name: "ERR_VALIDATE_WRONG_INTERNET_TEST_DOWNLOAD",
			req: &models.Request{
				Age: 18, EducationLevel: "bachelors_degree_or_high",
				PastExperiences: &models.PastExperiences{
					Sales:   true,
					Support: false,
				},
				InternetTest: &models.InternetTest{
					DownloadSpeed: 0.0,
					UploadSpeed:   50.0,
				},
			},
			expectedOutput: ERR_INTERNET_TEST,
		},
		{
			name: "ERR_VALIDATE_WRONG_INTERNET_TEST_UPLOAD",
			req: &models.Request{
				Age: 18, EducationLevel: "bachelors_degree_or_high",
				PastExperiences: &models.PastExperiences{
					Sales:   true,
					Support: false,
				},
				InternetTest: &models.InternetTest{
					DownloadSpeed: 50.0,
					UploadSpeed:   0.0,
				},
			},
			expectedOutput: ERR_INTERNET_TEST,
		},
		{
			name: "ERR_VALIDATE_EMPTY_OR_WRONG_WRITING_SCORE",
			req: &models.Request{
				Age: 18, EducationLevel: "no_education",
				PastExperiences: &models.PastExperiences{
					Sales:   true,
					Support: false,
				},
				InternetTest: &models.InternetTest{
					DownloadSpeed: 50.0,
					UploadSpeed:   50.0,
				},
				WritingScore: 0,
			},
			expectedOutput: ERR_WRITING_SCORE,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		projectService := &ProjectsService{}
		t.Run(tc.name, func(t *testing.T) {
			serviceError := projectService.Validate(tc.req)
			if assert.Error(t, serviceError) {
				assert.Equal(t, tc.expectedOutput, serviceError.Error())
			}
		})
	}
}
