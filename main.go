package main

import (
	"log"
	"os"
	"time"

	"github.com/PaulicStudios/42-cheat-alert/cheatalert"
	"github.com/PaulicStudios/42-cheat-alert/db"
	"github.com/PaulicStudios/42-cheat-alert/requests"
	"github.com/PaulicStudios/42-cheat-alert/telegram"
	"github.com/joho/godotenv"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	args := os.Args
	godotenv.Load()
	db.ConnectDB()
	requests.GetTokenSetClient()

	if len(args) < 2 {
		println("Usage: ./42-cheat-alert <update|timer|list>")
		return
	}

	switch args[1] {
	case "update":
		go (func() {
			telegram.Init_telegram()
		})()
		cheatalert.UpdateProjectsLastMonth()
	case "timer":
		log.Println("Starting timer")
		ticker := time.NewTicker(2 * time.Minute)
		go func() {
			for range ticker.C {
				cheatalert.UpdateProjects()
			}
		}()
		select {}
		// telegram.Init_telegram()
	case "list":
		teams := db.GetAllCheaterTeams()
		for _, team := range teams {
			println(team.Name)
		}
	case "test":
		println(requests.TestGetTeamDetails(12).Name)
	case "pisciner":
		cheatalert.GetPisciners()
	default:
		println("Usage: ./42-cheat-alert <update|timer|list|pisciner>")
	}
}
