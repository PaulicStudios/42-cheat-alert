package models

import "time"

type User struct {
	ID              int       `json:"id"`
	Email           string    `json:"email"`
	Login           string    `json:"login"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	UsualFullName   string    `json:"usual_full_name"`
	UsualFirstName  string    `json:"usual_first_name"`
	URL             string    `json:"url"`
	Phone           string    `json:"phone"`
	Displayname     string    `json:"displayname"`
	Kind            string    `json:"kind"`
	Image           Image     `json:"image" gorm:"-"`
	Staff           bool      `json:"staff?"`
	CorrectionPoint int       `json:"correction_point"`
	PoolMonth       string    `json:"pool_month"`
	PoolYear        string    `json:"pool_year"`
	Location        string    `json:"location"`
	Wallet          int       `json:"wallet"`
	AnonymizeDate   time.Time `json:"anonymize_date"`
	DataErasureDate time.Time `json:"data_erasure_date"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	AlumnizedAt     time.Time `json:"alumnized_at"`
	Alumni          bool      `json:"alumni?"`
	Active          bool      `json:"active?"`
}

type Image struct {
	Link     string `json:"link"`
	Versions struct {
		Large  string `json:"large"`
		Medium string `json:"medium"`
		Small  string `json:"small"`
		Micro  string `json:"micro"`
	} `json:"versions"`
}

type Project struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	ParentID any    `json:"parent_id"`
}
