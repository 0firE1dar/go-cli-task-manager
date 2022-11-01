package cmd

import (
	"log"
	"strconv"
	"taskmanager/database"

	cobra "github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "task do 1",
	Long: `task do [task id]
	will mark the task as completed`,
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				log.Fatalf("Failed to parse command args: %q", err)
			} else {
				ids = append(ids, id)
			}
		}
		
		tasks := database.ListTasks()

		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				log.Fatalf("Id is out of range: %v", id)
				continue
			}
			task := tasks[id-1]
			database.DeleteTask(task.Key)
		}
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}