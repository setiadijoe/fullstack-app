package app

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/setiadijoe/fullstack/app/config"
	"github.com/setiadijoe/fullstack/app/controller"
	"github.com/setiadijoe/fullstack/app/seed"
)

var server = controller.Server{}

// Run ...
func Run(port, dbDriver, dbName, dbHost, dbPort, dbUser, dbPass string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	}

	dbConfig := config.GetConfig()
	if port != "" {
		dbConfig.APPPort = port
	}
	if dbDriver != "" {
		dbConfig.DBDriver = dbDriver
	}
	if dbName != "" {
		dbConfig.DBName = dbName
	}
	if dbHost != "" {
		dbConfig.DBHost = dbHost
	}
	if dbPort != "" {
		dbConfig.DBPort = dbPort
	}
	if dbUser != "" {
		dbConfig.DBUser = dbUser
	}
	if dbPass != "" {
		dbConfig.DBPassword = dbPass
	}

	server.Initialize(dbConfig.DBDriver, dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBPort, dbConfig.DBHost, dbConfig.DBName)
	server.Run(fmt.Sprintf(":%s", dbConfig.APPPort))
}

// Migration ...
func Migration() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}
	dbConfig := config.GetConfig()

	server.Initialize(dbConfig.DBDriver, dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBPort, dbConfig.DBHost, dbConfig.DBName)

	seed.Load(server.DB)
}
