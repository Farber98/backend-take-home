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

func (ps *ProjectsService) CalculateEligibilityScore(req *models.Request) (score uint8) {

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
	case req.PastExperiences.Support:
		score += 3
		fallthrough
	case req.PastExperiences.Sales:
		score += 5
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
