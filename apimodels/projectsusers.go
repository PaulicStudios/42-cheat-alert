package apimodels

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
type Project struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	ParentID any    `json:"parent_id"`
}
type Versions struct {
	Large  string `json:"large"`
	Medium string `json:"medium"`
	Small  string `json:"small"`
	Micro  string `json:"micro"`
}
type Image struct {
	Link     string   `json:"link"`
	Versions Versions `json:"versions"`
}
type User struct {
	ID              int       `json:"id"`
	Email           string    `json:"email"`
	Login           string    `json:"login"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	UsualFullName   string    `json:"usual_full_name"`
	UsualFirstName  any       `json:"usual_first_name"`
	URL             string    `json:"url"`
	Phone           string    `json:"phone"`
	Displayname     string    `json:"displayname"`
	Kind            string    `json:"kind"`
	Image           Image     `json:"image"`
	Staff           bool      `json:"staff?"`
	CorrectionPoint int       `json:"correction_point"`
	PoolMonth       string    `json:"pool_month"`
	PoolYear        string    `json:"pool_year"`
	Location        any       `json:"location"`
	Wallet          int       `json:"wallet"`
	AnonymizeDate   string    `json:"anonymize_date"`
	DataErasureDate string    `json:"data_erasure_date"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	AlumnizedAt     any       `json:"alumnized_at"`
	Alumni          bool      `json:"alumni?"`
	Active          bool      `json:"active?"`
}
type Users struct {
	ID             int    `json:"id"`
	Login          string `json:"login"`
	URL            string `json:"url"`
	Leader         bool   `json:"leader"`
	Occurrence     int    `json:"occurrence"`
	Validated      bool   `json:"validated"`
	ProjectsUserID int    `json:"projects_user_id"`
}
type Teams struct {
	ID                int       `json:"id"`
	Name              string    `json:"name"`
	URL               string    `json:"url"`
	FinalMark         int       `json:"final_mark"`
	ProjectID         int       `json:"project_id"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	Status            string    `json:"status"`
	TerminatingAt     time.Time `json:"terminating_at"`
	Users             []Users   `json:"users"`
	Locked            bool      `json:"locked?"`
	Validated         bool      `json:"validated?"`
	Closed            bool      `json:"closed?"`
	RepoURL           string    `json:"repo_url"`
	RepoUUID          string    `json:"repo_uuid"`
	LockedAt          time.Time `json:"locked_at"`
	ClosedAt          time.Time `json:"closed_at"`
	ProjectSessionID  int       `json:"project_session_id"`
	ProjectGitlabPath string    `json:"project_gitlab_path"`
}
