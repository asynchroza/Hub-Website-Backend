package models

import "time"

type Event struct {
	EventID      string    `json:"eventid" validate: "required"`
	Title        string    `json:"name" validate: "required"`
	StartDate    time.Time `json:"startdate" validate: "required"`    // should be converted to ISO in frontend
	EndDate      time.Time `json:"enddate" validate: "required"`      // --||--
	Description  string    `json:"description" validate: "required"`  // long string
	Location     string    `json:"location" validate: "required"`     // location in text
	LocationLink string    `json:"locationlink" validate: "required"` // location google maps link
	Banner       string    `json:"banner" validate: "required"`       // banner image google drive link
}
