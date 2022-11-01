package cmd

import (
	"fmt"
	"strings"
	"taskmanager/database"

	cobra "github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a task to the list",
	Long: `task add [content]
	will add a new task to the list with the content as string`,
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		database.CreateTask(task)
	  fmt.Printf("The following task was added: %s", task)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}