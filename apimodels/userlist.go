package apimodels

import "time"

type UserList []struct {
	ID             int    `json:"id,omitempty"`
	Email          string `json:"email,omitempty"`
	Login          string `json:"login,omitempty"`
	FirstName      string `json:"first_name,omitempty"`
	LastName       string `json:"last_name,omitempty"`
	UsualFullName  string `json:"usual_full_name,omitempty"`
	UsualFirstName any    `json:"usual_first_name,omitempty"`
	URL            string `json:"url,omitempty"`
	Phone          string `json:"phone,omitempty"`
	Displayname    string `json:"displayname,omitempty"`
	Kind           string `json:"kind,omitempty"`
	Image          struct {
		Link     string `json:"link,omitempty"`
		Versions struct {
			Large  string `json:"large,omitempty"`
			Medium string `json:"medium,omitempty"`
			Small  string `json:"small,omitempty"`
			Micro  string `json:"micro,omitempty"`
		} `json:"versions,omitempty"`
	} `json:"image,omitempty"`
	Staff           bool      `json:"staff?,omitempty"`
	CorrectionPoint int       `json:"correction_point,omitempty"`
	PoolMonth       string    `json:"pool_month,omitempty"`
	PoolYear        string    `json:"pool_year,omitempty"`
	Location        any       `json:"location,omitempty"`
	Wallet          int       `json:"wallet,omitempty"`
	AnonymizeDate   time.Time `json:"anonymize_date,omitempty"`
	DataErasureDate time.Time `json:"data_erasure_date,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
	AlumnizedAt     any       `json:"alumnized_at,omitempty"`
	Alumni          bool      `json:"alumni?,omitempty"`
	Active          bool      `json:"active?,omitempty"`
}
