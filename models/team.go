package models

import "time"

type Team struct {
	ID                int
	Name              string
	FinalMark         int
	ProjectID         int
	CreatedAt         time.Time
	UpdatedAt         time.Time
	Status            string
	Locked            bool
	Validated         bool
	Closed            bool
	RepoURL           string
	RepoUUID          string
	LockedAt          time.Time
	ProjectSessionID  int
	ProjectGitlabPath string
	Users             []*User `gorm:"many2many:user_teams;"`
}
