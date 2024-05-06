package main

import (
	"os"
	"pepper-app/internal/database"
	"pepper-app/internal/routes"
	"pepper-app/internal/utils"
)

func init() {
	// Load environment variables
	utils.LoadEnvVariables()

	// Establish database connection and sync tables
	database.EstablishConnection()
}

func main() {
	r := routes.SetupRouter()
	err := r.Run(":" + os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}
	defer database.DB.Close()
}
