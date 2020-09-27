package main

//go install .
import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/parit90/todo-cli/task/cmd"
	"github.com/parit90/todo-cli/task/db"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	must(db.Init(dbPath))

	must(cmd.RootCmd.Execute())
}

func must(err error) {
	fmt.Println("main error...", err)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
