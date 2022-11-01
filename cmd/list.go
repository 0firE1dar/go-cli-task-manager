package cmd

import (
	"fmt"
	"taskmanager/database"

	cobra "github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all tasks",
	Long: `task list
	will list all tasks`,
	Run: func(cmd *cobra.Command, args []string) {
	  tasks := database.ListTasks()
	  if len(tasks) == 0 {
		fmt.Println("There are currently no tasks in the list")
	  }
	  if len(tasks) > 0 {
		fmt.Println("You have the following tasks:",)
		for i, task := range tasks {
		  fmt.Printf("%d. %s\n",i+1 ,task.Value)
		}
	  }
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}