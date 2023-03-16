package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	ID          int
	Title       string
	Description string
	DueDate     time.Time
	Completed   bool
}

type TaskManager struct {
	Tasks []Task
}

func (tm *TaskManager) AddTask(task Task) {
	tm.Tasks = append(tm.Tasks, task)
}

func (tm *TaskManager) DeleteTask(taskID int) {
	for i, task := range tm.Tasks {
		if task.ID == taskID {
			tm.Tasks = append(tm.Tasks[:i], tm.Tasks[i+1:]...)
			break
		}
	}
}

func (tm *TaskManager) ListTasks() {
	// make a dynamic formating based on the size of the longest string in each column
	//
	// ID    Title       Description     Due Date   Completed
	// 1     Task 1      Description 1   2019-01-01 false
	// 2     Task 2      Description 2   2019-01-02 true

	fmt.Println("ID\tTitle\tDescription\tDue Date\tCompleted")
	for _, task := range tm.Tasks {
		fmt.Printf("%d\t%s\t%s\t%s\t%v\n", task.ID, task.Title, task.Description, task.DueDate.Format("2006-01-02"), task.Completed)
	}
}

func main() {

	fmt.Println("Welcome to the Task Manager!")
	fmt.Println("--------------------------------------------")

	fmt.Println("Commands: add, delete, list, quit (or exit)")

	taskManager := TaskManager{}

	// Read task data from file
	file, err := os.Open("tasks.json")
	if err == nil {
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&taskManager)
		if err != nil {
			log.Fatal("Error decoding task data:", err)
		}
		file.Close()
	} else if !os.IsNotExist(err) {
		log.Fatal("Error opening task data file:", err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		if line == "" {
			continue
		}
		args := strings.Split(line, " ")
		command := args[0]
		switch command {
		case "add":
			if len(args) < 4 {
				fmt.Println("Usage: add <title> <description> <due-date>")
				continue
			}
			dueDate, err := time.Parse("2006-01-02", args[3])
			if err != nil {
				fmt.Println("Invalid due date format. Please use yyyy-mm-dd.")
				continue
			}
			task := Task{
				ID:          len(taskManager.Tasks) + 1,
				Title:       args[1],
				Description: args[2],
				DueDate:     dueDate,
			}
			taskManager.AddTask(task)
			fmt.Println("Task added successfully!")
		case "delete":
			if len(args) < 2 {
				fmt.Println("Usage: delete <id>")
				continue
			}
			taskID, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("Invalid task ID. Please enter a number.")
				continue
			}
			taskManager.DeleteTask(taskID)
			fmt.Println("Task deleted successfully!")
		case "list":
			taskManager.ListTasks()
		case "quit", "exit":
			fmt.Println("Goodbye!")
			// Write task data to file
			file, err := os.Create("tasks.json")
			if err == nil {
				encoder := json.NewEncoder(file)
				err = encoder.Encode(&taskManager)
				if err != nil {
					log.Fatal("Error encoding task data:", err)
				}
				file.Close()
				return // Exit program
			} else {
				log.Fatal("Error creating task data file:", err)
			}
		default:
			fmt.Println("Unknown command. Please try again.")
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
