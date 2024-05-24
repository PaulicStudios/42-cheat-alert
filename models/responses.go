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
	RetriableAt   any       `json:"retriable_at"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	User          User      `json:"user"`
	Teams         []struct {
		ID            int       `json:"id"`
		Name          string    `json:"name"`
		URL           string    `json:"url"`
		FinalMark     int       `json:"final_mark"`
		ProjectID     int       `json:"project_id"`
		CreatedAt     time.Time `json:"created_at"`
		UpdatedAt     time.Time `json:"updated_at"`
		Status        string    `json:"status"`
		TerminatingAt time.Time `json:"terminating_at"`
		Users         []struct {
			ID             int    `json:"id"`
			Login          string `json:"login"`
			URL            string `json:"url"`
			Leader         bool   `json:"leader"`
			Occurrence     int    `json:"occurrence"`
			Validated      bool   `json:"validated"`
			ProjectsUserID int    `json:"projects_user_id"`
		} `json:"users"`
		Locked            bool      `json:"locked?"`
		Validated         bool      `json:"validated?"`
		Closed            bool      `json:"closed?"`
		RepoURL           string    `json:"repo_url"`
		RepoUUID          string    `json:"repo_uuid"`
		LockedAt          time.Time `json:"locked_at"`
		ClosedAt          time.Time `json:"closed_at"`
		ProjectSessionID  int       `json:"project_session_id"`
		ProjectGitlabPath string    `json:"project_gitlab_path"`
	} `json:"teams"`
}
