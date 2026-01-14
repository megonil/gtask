// Package back
package back

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
)

type task struct {
	Title string   `json:"title"`
	Mode  taskMode `json:"mode"`
}

var tasks = make([]task, 0)

const taskJSONFile = "tasks.json"

func LoadTasks() error {
	err := withFileDir(gtaskDir, typeDir)
	if err != nil {
		return err
	}

	err = withFileDir(tasksJSONPath, typeFile)
	if err != nil {
		return err
	}

	var data, errFile = os.ReadFile(tasksJSONPath)
	if errFile != nil {
		return errFile
	}

	var decodeErr = json.Unmarshal(data, &tasks)
	if decodeErr != nil {
		return decodeErr
	}

	return nil
}

func ListTasks() {
	if len(tasks) == 0 {
		fmt.Println("List is empty")
		return
	}
	for i := range tasks {
		fmt.Printf("%d. %s: %s\n", i+1, tasks[i].Title, tasks[i].Mode.String())
	}
}

func AddTask(title string) (int, error) {
	if idx := findByTitle(title); idx != -1 {
		return idx, errors.New("task with such title already existing")
	}
	var newTask = task{Title: title, Mode: modePlanned}

	tasks = append(tasks, newTask)
	return len(tasks), writeTasks()
}

func Remove(args []string) error {
	var toDelete = make(map[int]bool)

	for _, arg := range args {
		if num, err := strconv.Atoi(arg); err == nil {
			if num > 0 && num <= len(tasks) {
				toDelete[num-1] = true
				continue
			}
		}

		var idx = findByTitle(arg)
		if idx == -1 {
			return fmt.Errorf("no task with title %s", arg)
		}
		toDelete[idx] = true
	}

	if len(toDelete) == 0 {
		fmt.Println("nothing to do")
		return nil
	}

	var i int
	tasks = slices.DeleteFunc(tasks, func(t task) bool {
		var shouldDelete = toDelete[i]
		i++
		return shouldDelete
	})

	return writeTasks()
}

func MarkTasks(args []string, mode taskMode) error {
	var toMark = make(map[int]bool)

	for _, arg := range args {
		if num, err := strconv.Atoi(arg); err == nil {
			if num > 0 && num <= len(tasks) {
				toMark[num-1] = true
				continue
			}
		}

		var idx = findByTitle(arg)
		if idx == -1 {
			return fmt.Errorf("no task with title %s", arg)
		}
		toMark[idx] = true
	}

	for idx, b := range toMark {
		if b {
			tasks[idx].Mode = mode
		}
	}

	return writeTasks()
}
