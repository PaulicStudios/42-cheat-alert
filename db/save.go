package db

import (
	"log"
	"time"

	"github.com/PaulicStudios/42-cheat-alert/apimodels"
	"github.com/PaulicStudios/42-cheat-alert/models"
)

func SaveUser(user *models.User) {
	result := db.Save(user)
	if result.Error != nil {
		log.Printf("Error saving user: %v", result.Error)
	}
}

func SaveApiUser(user *apimodels.User) {
	SaveUser(&models.User{
		ID:              user.ID,
		Email:           user.Email,
		Login:           user.Login,
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		UsualFullName:   user.UsualFullName,
		Phone:           user.Phone,
		Displayname:     user.Displayname,
		Kind:            user.Kind,
		Staff:           user.Staff,
		CorrectionPoint: user.CorrectionPoint,
		PoolMonth:       user.PoolMonth,
		PoolYear:        user.PoolYear,
		Wallet:          user.Wallet,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
		Alumni:          user.Alumni,
		Active:          user.Active,
	})
}

func SaveApiDetailUser(user *apimodels.UserDetailed) {
	var currentProject *string
	if len(user.ProjectsUsers) > 0 {
		currentProject = &user.ProjectsUsers[0].Project.Name
	} else {
		currentProject = nil
	}

	SaveUser(&models.User{
		ID:              user.ID,
		Email:           user.Email,
		Login:           user.Login,
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		UsualFullName:   user.UsualFullName,
		Phone:           user.Phone,
		Displayname:     user.Displayname,
		Kind:            user.Kind,
		Staff:           user.Staff,
		CorrectionPoint: user.CorrectionPoint,
		PoolMonth:       user.PoolMonth,
		PoolYear:        user.PoolYear,
		Wallet:          user.Wallet,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
		Alumni:          user.Alumni,
		Active:          user.Active,
		XP:              &user.CursusUsers[0].Level,
		CurrentProject:  currentProject,
	})
}

func SaveApiTeam(user *apimodels.User, team *apimodels.Teams) bool {
	teamModel := models.Team{
		ID:                team.ID,
		Name:              team.Name,
		FinalMark:         team.FinalMark,
		ProjectID:         team.ProjectID,
		CreatedAt:         &team.CreatedAt,
		UpdatedAt:         &team.UpdatedAt,
		Status:            team.Status,
		Locked:            team.Locked,
		Validated:         team.Validated,
		Closed:            team.Closed,
		RepoURL:           team.RepoURL,
		RepoUUID:          team.RepoUUID,
		ProjectSessionID:  team.ProjectSessionID,
		ProjectGitlabPath: team.ProjectGitlabPath,
	}
	if team.LockedAt.IsZero() {
		teamModel.LockedAt = nil
	} else {
		teamModel.LockedAt = &team.LockedAt
	}
	db.Model(&models.User{ID: user.ID}).Association("Teams").Append(&teamModel)
	var existingTeam models.Team
	db.Where("id = ?", team.ID).First(&existingTeam)

	if existingTeam.UpdatedAt == nil || time.Duration(team.UpdatedAt.Sub(*existingTeam.UpdatedAt)) > 0 {
		db.Save(&teamModel)
		return true
	}
	return false
}

func SaveApiTeamDetails(team *apimodels.Team) {
	for _, apiScaleTeam := range team.ScaleTeams {
		scaleTeam := models.ScaleTeam{
			ID:        apiScaleTeam.ID,
			TeamID:    team.ID,
			ScaleID:   apiScaleTeam.ScaleID,
			Comment:   apiScaleTeam.Comment,
			CreatedAt: apiScaleTeam.CreatedAt,
			UpdatedAt: apiScaleTeam.UpdatedAt,
			Feedback:  apiScaleTeam.Feedback,
			FinalMark: apiScaleTeam.FinalMark,
			BeginAt:   apiScaleTeam.BeginAt,
		}
		if !apiScaleTeam.FilledAt.IsZero() {
			scaleTeam.FilledAt = &apiScaleTeam.FilledAt
		}
		db.Save(&scaleTeam)

		flag := models.Flag{
			ScaleID:   apiScaleTeam.ID,
			ID:        apiScaleTeam.Flag.ID,
			Name:      apiScaleTeam.Flag.Name,
			Positive:  apiScaleTeam.Flag.Positive,
			Icon:      apiScaleTeam.Flag.Icon,
			CreatedAt: apiScaleTeam.Flag.CreatedAt,
			UpdatedAt: apiScaleTeam.Flag.UpdatedAt,
		}
		db.Save(&flag)
	}

	for _, teamUpload := range team.TeamsUploads {
		moulinette := models.Moulinette{
			ID:        teamUpload.ID,
			TeamID:    team.ID,
			FinalMark: teamUpload.FinalMark,
			Comment:   teamUpload.Comment,
			CreatedAt: teamUpload.CreatedAt,
			UploadID:  teamUpload.UploadID,
		}
		db.Save(&moulinette)
	}
}
