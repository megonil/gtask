// Package back
package back

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"slices"
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
	for i := range tasks {
		fmt.Printf("%d. %s: %s\n", i+1, tasks[i].Title, tasks[i].Mode.String())
	}
}

func AddTask(title string) error {
	var newTask = task{Title: title, Mode: modePlanned}

	tasks = append(tasks, newTask)
	return writeTasks()
}

func RemoveTask(title string) error {
	idx := slices.IndexFunc(tasks, func(t task) bool { return t.Title == title })
	if idx == -1 {
		return errors.New("no task with such title")
	}

	tasks = remove(tasks, idx)
	if tasks == nil {
		return errors.New("cannot remove task")
	}

	return writeTasks()
}

func RemoveTaskByID(idx int) error {
	if idx < 0 || idx >= len(tasks) {
		return errors.New("invalid index")
	}

	tasks = remove(tasks, idx)
	return writeTasks()
}

func MarkTask(title string, mode taskMode) error {
	var idx = findByTitle(title)
	if idx == -1 {
		return errors.New("no task with such title")
	}
	tasks[idx].Mode = mode

	return writeTasks()
}
