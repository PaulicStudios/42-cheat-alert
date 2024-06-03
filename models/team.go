package models

import (
	"time"
)

type Team struct {
	ID                int
	Name              string
	FinalMark         int
	ProjectID         int
	CreatedAt         *time.Time
	UpdatedAt         *time.Time
	Status            string
	Locked            bool
	Validated         bool
	Closed            bool
	RepoURL           string
	RepoUUID          string
	LockedAt          *time.Time
	ProjectSessionID  int
	ProjectGitlabPath string
	Users             []*User      `gorm:"many2many:user_teams;"`
	ScaleTeams        []*ScaleTeam `gorm:"foreignKey:TeamID;references:ID"`
}

type ScaleTeam struct {
	ID        int
	TeamID    int
	ScaleID   int
	Comment   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Feedback  string
	FinalMark int
	BeginAt   time.Time
	FilledAt  *time.Time
	Flag      Flag `gorm:"foreignKey:ScaleID;references:ID"`
}

type Flag struct {
	ScaleID   int `gorm:"primaryKey"`
	ID        int
	Name      string
	Positive  bool
	Icon      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
