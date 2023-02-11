/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/m4salah/tri/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var priority int

func addRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("datafile"))
	if err != nil {
		log.Println("The datafile is not found, we will create new datafile")
	}
	for _, arg := range args {
		item := todo.Item{Text: arg, Priority: priority}
		items = append(items, item)
	}
	todo.SaveItems(viper.GetString("datafile"), items)
	fmt.Println("Saved")
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add your task",
	Long:  `this command allows you to add your task.`,
	Run:   addRun,
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "priority of the task accept: 1, 2, 3")
}
