package models

type Response struct {
	Score              uint8    `json:"score,omitempty"`
	SelectedProject    string   `json:"selected_project,omitempty"`
	EligibleProjects   []string `json:"eligible_projects,omitempty"`
	InelegibleProjects []string `json:"ineligible_projects,omitempty"`
}
