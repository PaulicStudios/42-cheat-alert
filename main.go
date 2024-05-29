package main

import (
	"github.com/PaulicStudios/42-cheat-alert/cheatalert"
	"github.com/PaulicStudios/42-cheat-alert/db"
	"github.com/PaulicStudios/42-cheat-alert/requests"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db.ConnectDB()
	requests.GetTokenSetClient()
	// ticker := time.NewTicker(10 * time.Second)
	// go func() {
	// 	for range ticker.C {
	// 		cheatalert.UpdateProjects()
	// 	}
	// }()
	// go func() {
	// 	telegram.Init_telegram()
	// }()
	cheatalert.UpdateProjects()
	// teams := db.GetAllCheaterTeams()
	// for _, team := range teams {
	// 	println(team.Name)
	// }
}
