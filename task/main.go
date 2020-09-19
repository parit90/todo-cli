package main

//go install .
import "github.com/parit90/todo-cli/task/cmd"

func main() {
	cmd.RootCmd.Execute()
}
