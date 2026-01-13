package main

import (
	"back"
	"fmt"
	"front"
)

func main() {
	var err = back.LoadTasks()
	if err != nil {
		fmt.Printf("error: %v\n", err.Error())
		return
	}

	front.ParseAndExecute()
}
