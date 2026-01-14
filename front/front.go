// Package front
package front

import (
	"back"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func ParseAndExecute() {
	var rootCmd = &cobra.Command{Use: "gtask"}

	var addCmd = &cobra.Command{
		Use:   "add [task title]",
		Short: "Add new task",
		Args:  cobra.MaximumNArgs(255),
		Run: func(cmd *cobra.Command, args []string) {
			var taskTitle = strings.Join(args, " ")
			var idx, err = back.AddTask(taskTitle)
			if err != nil {
				fmt.Printf("error: %v\n", err.Error())
				return
			}

			fmt.Printf("Added task (ID = %d)\n", idx)
		},
	}

	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "lists all tasks",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			back.ListTasks()
		},
	}

	var removeCmd = &cobra.Command{
		Use:   "remove [<index> or <title>]",
		Short: "remove taks with <index> or <title>",
		Args:  cobra.MaximumNArgs(255),
		Run: func(cmd *cobra.Command, args []string) {
			back.Remove(args)
		},
	}

	var markCmd = &cobra.Command{
		Use:  "mark [mode] [<index> or <title>]",
		Args: cobra.MaximumNArgs(255),
		Run: func(cmd *cobra.Command, args []string) {
			var mode, err = back.TaskModeFromString(args[0])
			if err != nil {
				fmt.Printf("error: %v\n", err.Error())
				return
			}

			back.MarkTasks(args[1:], mode)
		},
	}

	rootCmd.AddCommand(addCmd, listCmd, removeCmd, markCmd)
	rootCmd.Execute()
}
