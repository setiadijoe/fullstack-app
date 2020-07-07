package cmd

/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import (
	"fmt"

	"github.com/setiadijoe/fullstack/app"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "run the service",
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetString("port")
		dbDriver, _ := cmd.Flags().GetString("dbDriver")
		dbHost, _ := cmd.Flags().GetString("dbHost")
		dbName, _ := cmd.Flags().GetString("dbName")
		dbUser, _ := cmd.Flags().GetString("dbUser")
		dbPass, _ := cmd.Flags().GetString("dbPass")
		dbPort, _ := cmd.Flags().GetString("dbPort")
		app.Run(port, dbDriver, dbName, dbHost, dbPort, dbUser, dbPass)
		fmt.Println("serve called")
	},
}

var migrationCmd = &cobra.Command{
	Use:   "migrate",
	Short: "run a migration and seeder",
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetString("port")
		dbDriver, _ := cmd.Flags().GetString("dbDriver")
		dbHost, _ := cmd.Flags().GetString("dbHost")
		dbName, _ := cmd.Flags().GetString("dbName")
		dbUser, _ := cmd.Flags().GetString("dbUser")
		dbPass, _ := cmd.Flags().GetString("dbPass")
		dbPort, _ := cmd.Flags().GetString("dbPort")
		app.Migration(port, dbDriver, dbName, dbHost, dbPort, dbUser, dbPass)
		fmt.Println("run migration")
	},
}
