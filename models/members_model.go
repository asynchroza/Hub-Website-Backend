package models

type Member struct {
	MemberID       string `json:"memberid, omitempty" validate:"required"`
	Firstname      string `json:"firstname, omitempty" validate:"required"`
	Lastname       string `json:"lastname, omitempty" validate:"required"`
	Department     string `json:"department, omitempty" validate:"required"`
	Position       string `json:"position, omitempty" validate:"required"`
	SocialLink     string `json:"sociallink, omitempty"`
	ProfilePicture string `json:"profilepicture, omitempty" validate:"required"`
}

type EditMember struct {
	MemberID       string `json:"memberid"`
	Firstname      string `json:"firstname"`
	Lastname       string `json:"lastname"`
	Department     string `json:"department"`
	Position       string `json:"position"`
	SocialLink     string `json:"sociallink"`
	ProfilePicture string `json:"profilepicture"`
}
