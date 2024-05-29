package main

import (
	"github.com/PaulicStudios/42-cheat-alert/cheatalert"
	"github.com/PaulicStudios/42-cheat-alert/db"
	"github.com/PaulicStudios/42-cheat-alert/requests"
	"github.com/PaulicStudios/42-cheat-alert/telegram"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db.ConnectDB()
	requests.GetTokenSetClient()
	cheatalert.UpdateProjects()
	telegram.Init_telegram()
}
