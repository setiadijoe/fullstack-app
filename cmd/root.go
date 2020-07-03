package cmd

/*
Copyright © 2020 NAME HERE <setiadi_y56@yahoo.co.id>

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
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var userLicense string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fullstack",
	Short: "fullstack is a CRUD sample of blogs applications",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(migrationCmd)
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringP("port", "p", "", "port for running")
	rootCmd.PersistentFlags().StringP("dbDriver", "", "", "name of db driver")
	rootCmd.PersistentFlags().StringP("dbHost", "", "", "host for db")
	rootCmd.PersistentFlags().StringP("dbPort", "", "", "port for running db")
	rootCmd.PersistentFlags().StringP("dbName", "", "", "name of db used")
	rootCmd.PersistentFlags().StringP("dbUser", "", "root", "username of db")
	rootCmd.PersistentFlags().StringP("dbPass", "", "admin", "password of db")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "YONATHAN SETIADI <setiadi_y56@yahoo.co.id>")
	viper.SetDefault("license", "apache")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".fullstack" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".fullstack")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
