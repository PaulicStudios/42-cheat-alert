package main

import (
	"github.com/PaulicStudios/42-cheat-alert/db"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db.ConnectDB()
	// requests.GetTokenSetClient()
}
