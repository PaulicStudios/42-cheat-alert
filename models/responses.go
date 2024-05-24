package models

import "time"

type ProjectsUsers []struct {
	ID            int       `json:"id"`
	Occurrence    int       `json:"occurrence"`
	FinalMark     int       `json:"final_mark"`
	Status        string    `json:"status"`
	Validated     bool      `json:"validated?"`
	CurrentTeamID int       `json:"current_team_id"`
	Project       Project   `json:"project"`
	CursusIds     []int     `json:"cursus_ids"`
	MarkedAt      time.Time `json:"marked_at"`
	Marked        bool      `json:"marked"`
	RetriableAt   time.Time `json:"retriable_at"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	User          User      `json:"user"`
	Teams         []Teams   `json:"teams"`
}
