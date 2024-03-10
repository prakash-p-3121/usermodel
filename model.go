package usermodel

import "time"

type User struct {
	ID             string    `json:"id"`
	FirstName      string    `json:"first-name"`
	LastName       string    `json:"last-name"`
	EmailID        string    `json:"email-id"`
	CountryCode    string    `json:"country-code"`
	PhoneNumberStr string    `json:"phone-number"`
	CreatedAt      time.Time `json:"created-at"`
	UpdatedAt      time.Time `json:"updated-at"`
}

type Password struct {
	UserID    string    `json:"id"` // unique-key
	Passwd    string    `json:"password"`
	UpdatedAt time.Time `json:"updated-at"`
}

type ShippingAddress struct {
	ID                  string    `json:"id"`
	UserID              string    `json:"user-id"`
	PropertyDescription string    `json:"property-desc"` /* Property Number or House Number or Building Number */
	Area                string    `json:"area"`
	CityOrTown          string    `json:"city-or-town"`
	State               string    `json:"state"`
	Country             string    `json:"country"`
	Pincode             string    `json:"pincode"`
	IsDefault           bool      `json:"is-default"`
	CreatedAt           time.Time `json:"created-at"`
	UpdatedAt           time.Time `json:"updated-at"`
}
