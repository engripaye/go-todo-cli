package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	Title string
	Done  bool
}

var tasks []Task

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go add <task>")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task description")
			return
		}
		title := os.Args[2]
		tasks = append(tasks, Task{Title: title, Done: false})
		saveTasks()
		fmt.Println("✅ Task added today:", title)
	}
}

func saveTasks() {
	file, _ := os.Create("tasks.json")
	defer file.Close()
	json.NewEncoder(file).Encode(tasks)

}
