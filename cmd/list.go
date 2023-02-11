/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/m4salah/tri/todo"
	"github.com/spf13/cobra"
)

func listRun(cmd *cobra.Command, args []string) {
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', tabwriter.TabIndent)
	defer w.Flush()
	items, err := todo.ReadItems(datafile)
	if err != nil {
		log.Printf("%v", err)
	}
	sort.Sort(todo.ByPrio(items))
	fmt.Printf("%#v\n", items)
	for _, i := range items {
		fmt.Fprintln(w, i.Label()+i.PrettyP()+"\t"+i.Text)
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
