package services

import (
	"cloudhumans/internal/models"
	"errors"
)

const (
	ERR_AGE              = "ERR. Age must be an integer equal or greater than 0."
	ERR_EDUCATION_LEVEL  = "ERR. Education level must be no_education, high_school or bachelors_degree_or_high."
	ERR_PAST_EXPERIENCES = "ERR. Past experiences must be provided."
	ERR_INTERNET_TEST    = "ERR. Internet test must be provided and be greater than 0.0."
	ERR_WRITING_SCORE    = "ERR. Writing score must be a float between 0.0 and 1.0."
	ERR_UNDER_AGE        = "ERR. Writing score must be a float between 0.0 and 1.0."
	PROJECT_NASA         = "Calculate the Dark Matter of the universe for Nasa"
	PROJECT_SCHRODINGER  = "Determine if the Schrodinger's cat is alive"
	PROJECT_YXZ          = "Attend to users support for a YXZ Company"
	PROJECT_XPTO         = "Collect specific people information from their social media for XPTO Company"
)

type ProjectsService struct {
}

func (ps *ProjectsService) Validate(req *models.Request) error {

	if req.Age <= 0 {
		return errors.New(ERR_AGE)
	}

	if req.EducationLevel != "no_education" && req.EducationLevel != "high_school" && req.EducationLevel != "bachelors_degree_or_high" {
		return errors.New(ERR_EDUCATION_LEVEL)
	}

	if req.PastExperiences == nil {
		return errors.New(ERR_PAST_EXPERIENCES)

	}

	if req.InternetTest == nil || req.InternetTest.DownloadSpeed <= 0 || req.InternetTest.UploadSpeed <= 0 {
		return errors.New(ERR_INTERNET_TEST)

	}

	if req.WritingScore <= 0.0 || req.WritingScore > 1 {
		return errors.New(ERR_WRITING_SCORE)
	}

	return nil
}

func (ps *ProjectsService) CalculateEligibilityScore(req *models.Request) (score int8) {

	if req.Age < 18 {
		return 0
	}

	switch req.EducationLevel {
	case "high_school":
		score++
	case "bachelors_degree_or_high":
		score += 2
	}

	switch {
	case req.PastExperiences.Support && req.PastExperiences.Sales:
		score += 8
	case req.PastExperiences.Sales:
		score += 5
	case req.PastExperiences.Support:
		score += 3
	}

	switch {
	case req.InternetTest.DownloadSpeed < 5:
		score--
	case req.InternetTest.DownloadSpeed > 50:
		score++
	}

	switch {
	case req.InternetTest.UploadSpeed < 5:
		score--
	case req.InternetTest.UploadSpeed > 50:
		score++
	}

	switch {
	case req.WritingScore < 0.3:
		score--
	case req.WritingScore >= 0.3 && req.WritingScore <= 0.7:
		score++
	case req.WritingScore > 0.7:
		score += 2
	}

	if req.ReferralCode == "token1234" {
		score++
	}

	return score
}

func (ps *ProjectsService) Projects(score int8) (selected string, eligible, ineligible []string) {
	switch {
	case score > 10:
		selected = PROJECT_NASA
		eligible = append(eligible, PROJECT_NASA, PROJECT_SCHRODINGER, PROJECT_YXZ, PROJECT_XPTO)
		ineligible = make([]string, 0, 0)
	case score > 5:
		selected = PROJECT_SCHRODINGER
		eligible = append(eligible, PROJECT_SCHRODINGER, PROJECT_YXZ, PROJECT_XPTO)
		ineligible = append(ineligible, PROJECT_NASA)
	case score > 3:
		selected = PROJECT_YXZ
		eligible = append(eligible, PROJECT_YXZ, PROJECT_XPTO)
		ineligible = append(ineligible, PROJECT_NASA, PROJECT_SCHRODINGER)
	case score > 2:
		selected = PROJECT_XPTO
		eligible = append(eligible, PROJECT_XPTO)
		ineligible = append(ineligible, PROJECT_NASA, PROJECT_SCHRODINGER, PROJECT_YXZ)
	default:
		eligible = make([]string, 0, 0)
		ineligible = append(ineligible, PROJECT_NASA, PROJECT_SCHRODINGER, PROJECT_YXZ, PROJECT_XPTO)
	}

	return selected, eligible, ineligible
}
