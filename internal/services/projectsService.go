package services

import (
	"cloudhumans/internal/models"
	"errors"
)

const (
	ERR_AGE              = "ERR. Age must be an integer equal or greater than 0."
	ERR_EDUCATION_LEVEL  = "ERR. Education level must be no_education, high_school or bachelors_degree_or_high."
	ERR_PAST_EXPERIENCES = "ERR. Past experiences must be provided."
	ERR_INTERNET_TEST    = "ERR. Internet test must be provided."
	ERR_WRITING_SCORE    = "ERR. Writing score must be a float between 0.0 and 1.0."
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

	if req.InternetTest == nil {
		return errors.New(ERR_INTERNET_TEST)

	}

	if req.WritingScore <= 0.0 || req.WritingScore > 1 {
		return errors.New(ERR_WRITING_SCORE)
	}

	return nil
}
