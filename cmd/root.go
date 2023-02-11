/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	// "github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tri",
	Short: "Tri is a todo list to help you manage your tasks",
	Long: `
This app will help you manage your tasks, and will show you what you have to do next, 
and prioritize your tasks.
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var (
	datafile string
	cfgFile  string
)

func initConfig() {
	viper.SetConfigName(".tri") // name of config file (without extension)
	viper.AddConfigPath("$HOME")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Using config file: ", viper.ConfigFileUsed())
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	home, err := os.UserHomeDir()
	if err != nil {
		log.Println("Unable to detact home directory. Please set data using --datafile flag")
	}
	filePath := home + string(os.PathSeparator) + ".tridos.json"

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVar(&datafile, "datafile",
		filePath,
		"data file store todos")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config",
		"",
		"config file (default is $HOME/.tri.yml)")
}
