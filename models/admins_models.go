package models

type Admin struct {
	Username string `json:"user, omitempty" validate: "required"`
	Password string `json:"password, omitempty" validate: "required"`
}
