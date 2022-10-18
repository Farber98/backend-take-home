package models

type Response struct {
	Score              int8     `json:"score,omitempty"`
	SelectedProject    string   `json:"selected_project,omitempty"`
	EligibleProjects   []string `json:"eligible_projects,omitempty"`
	InelegibleProjects []string `json:"ineligible_projects,omitempty"`
}

type ErrorResponse struct {
	Message string `json:"message,omitempty"`
}

func NewErrorResponse(msg string) ErrorResponse {
	return ErrorResponse{
		Message: msg,
	}
}
