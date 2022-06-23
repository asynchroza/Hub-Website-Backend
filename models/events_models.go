package models

import "time"

type Event struct {
	EventID      string    `json:"eventid"`
	Title        string    `json:"title"`
	StartDate    time.Time `json:"startdate"`    // should be converted to ISO in frontend
	EndDate      time.Time `json:"enddate"`      // --||--
	Description  string    `json:"description"`  // long string
	Location     string    `json:"location"`     // location in text
	LocationLink string    `json:"locationlink"` // location google maps link
	Banner       string    `json:"banner"`       // banner image google drive link
}
