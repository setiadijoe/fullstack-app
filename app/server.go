package app

import (
	"fmt"

	"github.com/setiadijoe/fullstack/app/config"
	"github.com/setiadijoe/fullstack/app/controller"
	"github.com/setiadijoe/fullstack/app/seed"
)

var server = controller.Server{}

// Run ...
func Run(port, dbDriver, dbName, dbHost, dbPort, dbUser, dbPass string) {
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
func Migration(port, dbDriver, dbName, dbHost, dbPort, dbUser, dbPass string) {
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

	seed.Load(server.DB)
}
