package services

import (
	"cloudhumans/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateErrors(t *testing.T) {
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

func TestValidateSuccess(t *testing.T) {
	testCases := []struct {
		name           string
		req            *models.Request
		expectedOutput string
	}{
		{
			name: "SUCCESS_VALIDATE",
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
				WritingScore: 0.5,
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		projectService := &ProjectsService{}
		t.Run(tc.name, func(t *testing.T) {
			serviceError := projectService.Validate(tc.req)
			if assert.NoError(t, serviceError) {
				assert.Nil(t, nil, serviceError)
			}
		})
	}
}

func TestCalculateEligibilityScoreSuccess(t *testing.T) {
	testCases := []struct {
		name          string
		req           *models.Request
		expectedScore int8
	}{
		{
			name: "SUCCESS_UNDER_AGE",
			req: &models.Request{
				Age: 15, EducationLevel: "no_education",
				PastExperiences: &models.PastExperiences{
					Sales:   true,
					Support: false,
				},
				InternetTest: &models.InternetTest{
					DownloadSpeed: 50.0,
					UploadSpeed:   50.0,
				},
				WritingScore: 0.5,
				ReferralCode: "token1234",
			},
			expectedScore: 0,
		},
		{
			name: "SUCCESS_MINIMUM_SCORE",
			req: &models.Request{
				Age: 18, EducationLevel: "no_education",
				PastExperiences: &models.PastExperiences{
					Sales:   false,
					Support: false,
				},
				InternetTest: &models.InternetTest{
					DownloadSpeed: 1.0,
					UploadSpeed:   1.0,
				},
				WritingScore: 0.2,
				ReferralCode: "not valid",
			},
			expectedScore: -3,
		},
		{
			name: "SUCCESS_ZERO_SCORE",
			req: &models.Request{
				Age: 18, EducationLevel: "no_education",
				PastExperiences: &models.PastExperiences{
					Support: true,
					Sales:   false,
				},
				InternetTest: &models.InternetTest{
					DownloadSpeed: 1.0,
					UploadSpeed:   1.0,
				},
				WritingScore: 0.2,
				ReferralCode: "not valid",
			},
			expectedScore: 0,
		},
		{
			name: "SUCCESS_MAXIMUM_SCORE",
			req: &models.Request{
				Age: 18, EducationLevel: "bachelors_degree_or_high",
				PastExperiences: &models.PastExperiences{
					Support: true,
					Sales:   true,
				},
				InternetTest: &models.InternetTest{
					DownloadSpeed: 51.0,
					UploadSpeed:   51.0,
				},
				WritingScore: 0.8,
				ReferralCode: "token1234",
			},
			expectedScore: 15,
		},
		{
			name: "SUCCESS_HIGH_SCHOOL_SALES_AVG_INTERNET_AVG_WRITING_NO_REFERRAL",
			req: &models.Request{
				Age: 18, EducationLevel: "high_school",
				PastExperiences: &models.PastExperiences{
					Support: false,
					Sales:   true,
				},
				InternetTest: &models.InternetTest{
					DownloadSpeed: 40.0,
					UploadSpeed:   40.0,
				},
				WritingScore: 0.4,
				ReferralCode: "not valid",
			},
			expectedScore: 7,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		projectService := &ProjectsService{}
		t.Run(tc.name, func(t *testing.T) {
			score := projectService.CalculateEligibilityScore(tc.req)
			assert.Equal(t, tc.expectedScore, score)
		})
	}
}
func TestProjectsSuccess(t *testing.T) {
	testCases := []struct {
		name       string
		score      int8
		selected   string
		eligible   []string
		ineligible []string
	}{
		{
			name:       "SUCCESS_ALL_NOT_ELIGIBLE_NEGATIVE_SCORE",
			score:      -3,
			selected:   "",
			eligible:   []string{},
			ineligible: []string{PROJECT_NASA, PROJECT_SCHRODINGER, PROJECT_YXZ, PROJECT_XPTO},
		},
		{
			name:       "SUCCESS_ALL_NOT_ELIGIBLE_ZERO_SCORE",
			score:      0,
			selected:   "",
			eligible:   []string{},
			ineligible: []string{PROJECT_NASA, PROJECT_SCHRODINGER, PROJECT_YXZ, PROJECT_XPTO},
		},
		{
			name:       "SUCCESS_ALL_NOT_ELIGIBLE_TWO_SCORE",
			score:      2,
			selected:   "",
			eligible:   []string{},
			ineligible: []string{PROJECT_NASA, PROJECT_SCHRODINGER, PROJECT_YXZ, PROJECT_XPTO},
		},
		{
			name:       "SUCCESS_XPTO_ELIGIBLE_THREE_SCORE",
			score:      3,
			selected:   PROJECT_XPTO,
			eligible:   []string{PROJECT_XPTO},
			ineligible: []string{PROJECT_NASA, PROJECT_SCHRODINGER, PROJECT_YXZ},
		},
		{
			name:       "SUCCESS_XPTO_YXZ_ELIGIBLE_FOUR_SCORE",
			score:      4,
			selected:   PROJECT_YXZ,
			eligible:   []string{PROJECT_YXZ, PROJECT_XPTO},
			ineligible: []string{PROJECT_NASA, PROJECT_SCHRODINGER},
		},
		{
			name:       "SUCCESS_XPTO_YXZ_ELIGIBLE_FIVE_SCORE",
			score:      5,
			selected:   PROJECT_YXZ,
			eligible:   []string{PROJECT_YXZ, PROJECT_XPTO},
			ineligible: []string{PROJECT_NASA, PROJECT_SCHRODINGER},
		},
		{
			name:       "SUCCESS_XPTO_YXZ_SCHRODINGER_ELIGIBLE_SIX_SCORE",
			score:      6,
			selected:   PROJECT_SCHRODINGER,
			eligible:   []string{PROJECT_SCHRODINGER, PROJECT_YXZ, PROJECT_XPTO},
			ineligible: []string{PROJECT_NASA},
		},
		{
			name:       "SUCCESS_XPTO_YXZ_SCHRODINGER_ELIGIBLE_TEN_SCORE",
			score:      10,
			selected:   PROJECT_SCHRODINGER,
			eligible:   []string{PROJECT_SCHRODINGER, PROJECT_YXZ, PROJECT_XPTO},
			ineligible: []string{PROJECT_NASA},
		},
		{
			name:       "SUCCESS_ALL_ELIGIBLE_ELEVEN_SCORE",
			score:      11,
			selected:   PROJECT_NASA,
			eligible:   []string{PROJECT_NASA, PROJECT_SCHRODINGER, PROJECT_YXZ, PROJECT_XPTO},
			ineligible: []string{},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		projectService := &ProjectsService{}
		t.Run(tc.name, func(t *testing.T) {
			selected, eligible, not_eligible := projectService.Projects(tc.score)
			assert.Equal(t, tc.selected, selected)
			assert.Equal(t, tc.eligible, eligible)
			assert.Equal(t, tc.ineligible, not_eligible)
		})
	}
}
