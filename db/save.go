package db

import (
	"log"

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

func SaveApiTeam(user *apimodels.User, team *apimodels.Teams) {
	db.Model(&models.User{ID: user.ID}).Association("Teams").Append(&models.Team{
		ID:                team.ID,
		Name:              team.Name,
		FinalMark:         team.FinalMark,
		ProjectID:         team.ProjectID,
		CreatedAt:         team.CreatedAt,
		UpdatedAt:         team.UpdatedAt,
		Status:            team.Status,
		Locked:            team.Locked,
		Validated:         team.Validated,
		Closed:            team.Closed,
		RepoURL:           team.RepoURL,
		RepoUUID:          team.RepoUUID,
		LockedAt:          team.LockedAt,
		ProjectSessionID:  team.ProjectSessionID,
		ProjectGitlabPath: team.ProjectGitlabPath,
	})
}

// func SaveTeams(team []models.Team) {
// 	db.Save(&team)
// }
