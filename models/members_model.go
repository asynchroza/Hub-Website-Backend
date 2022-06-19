package models

type Member struct {
	Firstname  string `json:"first_name, omitempty" validate:"required"`
	Lastname   string `json:"last_name, omitempty" validate:"required"`
	Department string `json:"department, omitempty" validate:"required"`
	Position   string `json:"job_position, omitempty" validate:"required"`
	LinkedIn   string `json:"social_link, omitempty" validate:"required"`
}
