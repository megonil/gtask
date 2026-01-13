package back

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"slices"
)

type fileDirType int

const (
	typeDir fileDirType = iota
	typeFile
)

func getConfigDir() string {
	var dir, err = os.UserConfigDir()
	if err != nil {
		panic("Cannot get config dir")
	}

	return dir
}

var configDir = getConfigDir()
var gtaskDir = filepath.Join(configDir, "gtask")
var tasksJSONPath = filepath.Join(gtaskDir, taskJSONFile)

func withFileDir(fileDir string, kind fileDirType) error {
	if _, err := os.Stat(fileDir); errors.Is(err, os.ErrNotExist) {
		if kind == typeFile {
			file, err := os.Create(fileDir)
			if err != nil {
				return err
			}

			file.WriteString("[]")
			file.Close()
		} else {
			os.Mkdir(fileDir, 0755)
		}

	}

	return nil
}

func writeTasks() error {
	var data, err = json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(tasksJSONPath, data, 0644)
}

func checkString(want string, found string) bool {
	return false
}

func findByTitle(title string) int {
	return slices.IndexFunc(tasks, func(t task) bool { return t.Title == title })
}
