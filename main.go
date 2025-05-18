package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID    int
	Title string
	Done  bool
}

var tasks []Task
var nextID int = 1

func LoadTasks() {
	// Placeholder: In a real app, load from file or DB
	tasks = []Task{}
}

func AddTask(title string) {
	tasks = append(tasks, Task{ID: nextID, Title: title, Done: false})
	nextID++
}

func ListTasks() []Task {
	return tasks
}

func MarkDone(id int) {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Done = true
			return
		}
	}
}

func DeleteTask(id int) {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return
		}
	}
}

func main() {
	LoadTasks()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n== To-Do List ==")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Mark Task as Done")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Print("Enter task: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)
			AddTask(title)
			fmt.Println("Task added!")
		case "2":
			tasks := ListTasks()
			if len(tasks) == 0 {
				fmt.Println("No tasks yet.")
			} else {
				for _, task := range tasks {
					status := " "
					if task.Done {
						status = "x"
					}
					fmt.Printf("[%s] %d: %s\n", status, task.ID, task.Title)
				}
			}
		case "3":
			fmt.Print("Enter task ID to mark as done: ")
			idStr, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idStr))
			MarkDone(id)
			fmt.Println("Task marked as done.")
		case "4":
			fmt.Print("Enter task ID to delete: ")
			idStr, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idStr))
			DeleteTask(id)
			fmt.Println("Task deleted.")
		case "5":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Try again.")
		}
	}
}
