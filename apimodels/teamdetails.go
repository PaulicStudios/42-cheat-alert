package apimodels

import "time"

type Team struct {
	ID                int            `json:"id"`
	Name              string         `json:"name"`
	URL               string         `json:"url"`
	FinalMark         int            `json:"final_mark"`
	ProjectID         int            `json:"project_id"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	Status            string         `json:"status"`
	TerminatingAt     time.Time      `json:"terminating_at"`
	Users             []Users        `json:"users"`
	Locked            bool           `json:"locked?"`
	Validated         bool           `json:"validated?"`
	Closed            bool           `json:"closed?"`
	RepoURL           string         `json:"repo_url"`
	RepoUUID          string         `json:"repo_uuid"`
	LockedAt          time.Time      `json:"locked_at"`
	ClosedAt          time.Time      `json:"closed_at"`
	ProjectSessionID  int            `json:"project_session_id"`
	ProjectGitlabPath string         `json:"project_gitlab_path"`
	ScaleTeams        []ScaleTeams   `json:"scale_teams"`
	TeamsUploads      []TeamsUploads `json:"teams_uploads"`
}
type Flag struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Positive  bool      `json:"positive"`
	Icon      string    `json:"icon"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type Correcteds struct {
	ID    int    `json:"id"`
	Login string `json:"login"`
	URL   string `json:"url"`
}
type Corrector struct {
	ID    int    `json:"id"`
	Login string `json:"login"`
	URL   string `json:"url"`
}
type Truant struct {
}
type ScaleTeams struct {
	ID                   int          `json:"id"`
	ScaleID              int          `json:"scale_id"`
	Comment              string       `json:"comment"`
	CreatedAt            time.Time    `json:"created_at"`
	UpdatedAt            time.Time    `json:"updated_at"`
	Feedback             string       `json:"feedback"`
	FinalMark            int          `json:"final_mark"`
	Flag                 Flag         `json:"flag"`
	BeginAt              time.Time    `json:"begin_at"`
	Correcteds           []Correcteds `json:"correcteds"`
	Corrector            Corrector    `json:"corrector"`
	Truant               Truant       `json:"truant"`
	FilledAt             time.Time    `json:"filled_at"`
	QuestionsWithAnswers []any        `json:"questions_with_answers"`
}
type TeamsUploads struct {
	ID        int       `json:"id"`
	FinalMark int       `json:"final_mark"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UploadID  int       `json:"upload_id"`
}
