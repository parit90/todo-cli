package cmd

import (
	"fmt"
	"os"

	"github.com/parit90/todo-cli/task/db"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all of your task",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hi.......")
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Here......")
			fmt.Println("Something went wront:", err.Error())
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("You dont have any task pending")
			return
		}

		fmt.Println("You have the following task:")
		for i, task := range tasks {
			fmt.Printf("%d, %s\n", i+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
