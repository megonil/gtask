// Package front
package front

import (
	"back"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func ParseAndExecute() {
	var rootCmd = &cobra.Command{Use: "gtask"}

	var addCmd = &cobra.Command{
		Use:   "add [task title]",
		Short: "Add new task",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var task = args[0]
			var err = back.AddTask(task)
			if err != nil {
				fmt.Printf("error: %v\n", err.Error())
				return
			}
		},
	}

	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "lists all tasks",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("List of tasks")
			back.ListTasks()
		},
	}

	var removeCmd = &cobra.Command{
		Use:   "remove <index> or <title>",
		Short: "remove taks with <index> or <title>",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var task = args[0]
			var num, err = strconv.Atoi(task)
			if err == nil {
				err = back.RemoveTaskByID(num)
				if err != nil {
					fmt.Printf("error: %v\n", err.Error())
				}
			} else {
				err = back.RemoveTask(task)
				if err != nil {
					fmt.Printf("error: %v\n", err.Error())
				}
			}
		},
	}

	rootCmd.AddCommand(addCmd, listCmd, removeCmd)
	rootCmd.Execute()
}
