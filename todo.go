package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var tasks []Task
var nextID int = 1

const dataFile = "tasks.json"

// Load tasks from file
func LoadTasks() {
	file, err := os.Open(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			tasks = []Task{}
			return
		}
		fmt.Println("Error reading tasks:", err)
		return
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		fmt.Println("Error decoding tasks:", err)
		tasks = []Task{}
	}

	// Set nextID
	for _, task := range tasks {
		if task.ID >= nextID {
			nextID = task.ID + 1
		}
	}
}

// Save tasks to file
func SaveTasks() {
	file, err := os.Create(dataFile)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(tasks)
	if err != nil {
		fmt.Println("Error encoding tasks:", err)
	}
}

func AddTask(title string) {
	task := Task{ID: nextID, Title: title, Done: false}
	tasks = append(tasks, task)
	nextID++
	SaveTasks()
}

func ListTasks() []Task {
	return tasks
}

func MarkDone(id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Done = true
			SaveTasks()
			return
		}
	}
	fmt.Println("Task not found.")
}

func DeleteTask(id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			SaveTasks()
			return
		}
	}
	fmt.Println("Task not found.")
}
