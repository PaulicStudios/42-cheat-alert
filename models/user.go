package models

import "time"

type User struct {
	ID              int
	Email           string
	Login           string
	FirstName       string
	LastName        string
	UsualFullName   string
	Phone           string
	Displayname     string
	Kind            string
	Staff           bool
	CorrectionPoint int
	PoolMonth       string
	PoolYear        string
	Wallet          int
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Alumni          bool
	Active          bool
	XP              *float64
	CurrentProject  *string
	Teams           []*Team `gorm:"many2many:user_teams;"`
}
