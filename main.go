package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/aliadelharrat/go-complete-todos-terminal-app/input"
	"github.com/aliadelharrat/go-complete-todos-terminal-app/output"
	"github.com/aliadelharrat/go-complete-todos-terminal-app/task"
)

const tasksFileName = "tasks.json"

func main() {
	var tasks []task.Task
	nextID := 1

	err := task.LoadTasksFromFile(tasksFileName, &tasks, &nextID)
	if err != nil {
		output.PrintError(fmt.Sprintln("Error loading tasks: ", err))
		fmt.Println()
	}

	for {
		fmt.Print("What do you wanna do? (options: add, view all, view pending, view completed, edit, complete, clear completed, delete, exit): ")
		command, err := input.GetInput()
		if err != nil {
			output.PrintError("Couldn't get command input. Please try again.")
			continue
		}
		switch command {
		case "add":
			fmt.Print("Type your description: ")
			description, err := input.GetInput()
			if err != nil {
				output.PrintError("Couldn't get description. Task not added.")
				continue
			}
			task.AddTask(description, &tasks, &nextID)
		case "view all":
			task.ViewTasks(tasks, "all")
		case "view pending":
			task.ViewTasks(tasks, "pending")
		case "view completed":
			task.ViewTasks(tasks, "completed")
		case "edit":
			fmt.Print("Enter ID of task to edit: ")
			IDStr, err := input.GetInput()
			if err != nil {
				output.PrintError("Couldn't get ID for edit. Please try again.")
				continue
			}
			id, err := strconv.Atoi(IDStr)
			if err != nil {
				output.PrintError("Invalid ID format.")
				continue
			}
			fmt.Print("Enter new description: ")
			newDescription, err := input.GetInput()
			if err != nil {
				output.PrintError("Couldn't get new description. Task not edited.")
				continue
			}
			task.EditTask(id, newDescription, &tasks)
		case "complete":
			fmt.Print("Enter ID of task to mark as complete: ")
			IDStr, err := input.GetInput()
			if err != nil {
				output.PrintError("Couldn't get ID for complete. Please try again.")
				continue
			}
			id, err := strconv.Atoi(IDStr)
			if err != nil {
				output.PrintError("Invalid ID format.")
				continue
			}
			task.CompleteTask(id, &tasks)
		case "clear completed":
			task.ClearCompletedTasks(&tasks)
		case "delete":
			fmt.Print("Enter ID of task to delete: ")
			IDStr, err := input.GetInput()
			if err != nil {
				output.PrintError("Couldn't get ID for delete. Please try again.")
			}
			id, err := strconv.Atoi(IDStr)
			if err != nil {
				output.PrintError("Invalid ID format.")
				continue
			}
			task.DeleteTask(id, &tasks)
		case "exit":
			if err := task.SaveTasksToFile(tasksFileName, tasks); err != nil {
				output.PrintError(fmt.Sprint("Error saving tasks: ", err))
				fmt.Println()
			}
			output.PrintSuccess("Thanks for your visit 😁")
			os.Exit(0)
		default:
			output.PrintError(fmt.Sprintf("Command '%s' is unknown", command))
		}
	}
}
