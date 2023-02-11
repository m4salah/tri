/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/m4salah/tri/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	doneOpt  bool
	allOpt   bool
	queryOpt string
)

func listRun(cmd *cobra.Command, args []string) {
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', tabwriter.TabIndent)
	defer w.Flush()
	items, err := todo.ReadItems(viper.GetString("datafile"))
	if err != nil {
		log.Printf("Error while readin datafile, %v", err)
	}
	sort.Sort(todo.ByPrio(items))

	for _, i := range items {
		if allOpt || i.Done == doneOpt {
			if len(queryOpt) > 0 {
				if strings.Contains(i.Text, queryOpt) {
					fmt.Println(i)
				}
			} else {
				fmt.Println(i)
			}

		}
	}
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todo items",
	Long:  `Listing the todo items`,
	Run:   listRun,
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVar(&doneOpt, "done", false, "Show 'Done' todos")
	listCmd.Flags().BoolVar(&allOpt, "all", false, "Show All todos")
	listCmd.Flags().StringVarP(&queryOpt, "query", "q", "", "Search query")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
