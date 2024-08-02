package apimodels

import "time"

type UserDetailed struct {
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
	Location        string    `json:"location,omitempty"`
	Wallet          int       `json:"wallet,omitempty"`
	AnonymizeDate   time.Time `json:"anonymize_date,omitempty"`
	DataErasureDate time.Time `json:"data_erasure_date,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
	AlumnizedAt     any       `json:"alumnized_at,omitempty"`
	Alumni          bool      `json:"alumni?,omitempty"`
	Active          bool      `json:"active?,omitempty"`
	Groups          []any     `json:"groups,omitempty"`
	CursusUsers     []struct {
		Grade  any     `json:"grade,omitempty"`
		Level  float64 `json:"level,omitempty"`
		Skills []struct {
			ID    int     `json:"id,omitempty"`
			Name  string  `json:"name,omitempty"`
			Level float64 `json:"level,omitempty"`
		} `json:"skills,omitempty"`
		BlackholedAt any       `json:"blackholed_at,omitempty"`
		ID           int       `json:"id,omitempty"`
		BeginAt      time.Time `json:"begin_at,omitempty"`
		EndAt        time.Time `json:"end_at,omitempty"`
		CursusID     int       `json:"cursus_id,omitempty"`
		HasCoalition bool      `json:"has_coalition,omitempty"`
		CreatedAt    time.Time `json:"created_at,omitempty"`
		UpdatedAt    time.Time `json:"updated_at,omitempty"`
		User         struct {
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
			Location        string    `json:"location,omitempty"`
			Wallet          int       `json:"wallet,omitempty"`
			AnonymizeDate   time.Time `json:"anonymize_date,omitempty"`
			DataErasureDate time.Time `json:"data_erasure_date,omitempty"`
			CreatedAt       time.Time `json:"created_at,omitempty"`
			UpdatedAt       time.Time `json:"updated_at,omitempty"`
			AlumnizedAt     any       `json:"alumnized_at,omitempty"`
			Alumni          bool      `json:"alumni?,omitempty"`
			Active          bool      `json:"active?,omitempty"`
		} `json:"user,omitempty"`
		Cursus struct {
			ID        int       `json:"id,omitempty"`
			CreatedAt time.Time `json:"created_at,omitempty"`
			Name      string    `json:"name,omitempty"`
			Slug      string    `json:"slug,omitempty"`
			Kind      string    `json:"kind,omitempty"`
		} `json:"cursus,omitempty"`
	} `json:"cursus_users,omitempty"`
	ProjectsUsers []struct {
		ID            int    `json:"id,omitempty"`
		Occurrence    int    `json:"occurrence,omitempty"`
		FinalMark     int    `json:"final_mark,omitempty"`
		Status        string `json:"status,omitempty"`
		Validated     bool   `json:"validated?,omitempty"`
		CurrentTeamID int    `json:"current_team_id,omitempty"`
		Project       struct {
			ID       int    `json:"id,omitempty"`
			Name     string `json:"name,omitempty"`
			Slug     string `json:"slug,omitempty"`
			ParentID any    `json:"parent_id,omitempty"`
		} `json:"project,omitempty"`
		CursusIds   []int     `json:"cursus_ids,omitempty"`
		MarkedAt    time.Time `json:"marked_at,omitempty"`
		Marked      bool      `json:"marked,omitempty"`
		RetriableAt time.Time `json:"retriable_at,omitempty"`
		CreatedAt   time.Time `json:"created_at,omitempty"`
		UpdatedAt   time.Time `json:"updated_at,omitempty"`
	} `json:"projects_users,omitempty"`
	LanguagesUsers []struct {
		ID         int       `json:"id,omitempty"`
		LanguageID int       `json:"language_id,omitempty"`
		UserID     int       `json:"user_id,omitempty"`
		Position   int       `json:"position,omitempty"`
		CreatedAt  time.Time `json:"created_at,omitempty"`
	} `json:"languages_users,omitempty"`
	Achievements []struct {
		ID           int    `json:"id,omitempty"`
		Name         string `json:"name,omitempty"`
		Description  string `json:"description,omitempty"`
		Tier         string `json:"tier,omitempty"`
		Kind         string `json:"kind,omitempty"`
		Visible      bool   `json:"visible,omitempty"`
		Image        string `json:"image,omitempty"`
		NbrOfSuccess any    `json:"nbr_of_success,omitempty"`
		UsersURL     string `json:"users_url,omitempty"`
	} `json:"achievements,omitempty"`
	Titles          []any `json:"titles,omitempty"`
	TitlesUsers     []any `json:"titles_users,omitempty"`
	Partnerships    []any `json:"partnerships,omitempty"`
	Patroned        []any `json:"patroned,omitempty"`
	Patroning       []any `json:"patroning,omitempty"`
	ExpertisesUsers []any `json:"expertises_users,omitempty"`
	Roles           []any `json:"roles,omitempty"`
	Campus          []struct {
		ID       int    `json:"id,omitempty"`
		Name     string `json:"name,omitempty"`
		TimeZone string `json:"time_zone,omitempty"`
		Language struct {
			ID         int       `json:"id,omitempty"`
			Name       string    `json:"name,omitempty"`
			Identifier string    `json:"identifier,omitempty"`
			CreatedAt  time.Time `json:"created_at,omitempty"`
			UpdatedAt  time.Time `json:"updated_at,omitempty"`
		} `json:"language,omitempty"`
		UsersCount         int    `json:"users_count,omitempty"`
		VogsphereID        int    `json:"vogsphere_id,omitempty"`
		Country            string `json:"country,omitempty"`
		Address            string `json:"address,omitempty"`
		Zip                string `json:"zip,omitempty"`
		City               string `json:"city,omitempty"`
		Website            string `json:"website,omitempty"`
		Facebook           string `json:"facebook,omitempty"`
		Twitter            string `json:"twitter,omitempty"`
		Active             bool   `json:"active,omitempty"`
		Public             bool   `json:"public,omitempty"`
		EmailExtension     string `json:"email_extension,omitempty"`
		DefaultHiddenPhone bool   `json:"default_hidden_phone,omitempty"`
	} `json:"campus,omitempty"`
	CampusUsers []struct {
		ID        int       `json:"id,omitempty"`
		UserID    int       `json:"user_id,omitempty"`
		CampusID  int       `json:"campus_id,omitempty"`
		IsPrimary bool      `json:"is_primary,omitempty"`
		CreatedAt time.Time `json:"created_at,omitempty"`
		UpdatedAt time.Time `json:"updated_at,omitempty"`
	} `json:"campus_users,omitempty"`
}
