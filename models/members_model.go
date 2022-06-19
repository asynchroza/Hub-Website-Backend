package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Member struct {
	Id         primitive.ObjectID `json:"id, omitempty"`
	UserKey    string             `json:"user_key, omitempty" validate:"required"` // will be used to delete user
	Firstname  string             `json:"first_name, omitempty" validate:"required"`
	Lastname   string             `json:"last_name, omitempty" validate:"required"`
	Department string             `json:"department, omitempty" validate:"required"`
	Position   string             `json:"job_position, omitempty" validate:"required"`
	LinkedIn   string             `json:"social_link, omitempty" validate:"required"`
}
