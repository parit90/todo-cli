package cmd

import (
	"fmt"
	"strings"

	"github.com/parit90/todo-cli/task/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to the task list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}
		fmt.Printf("Task: %s Add to your task list", task)

	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
