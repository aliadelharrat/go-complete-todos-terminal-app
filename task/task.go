package task

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"

	"github.com/aliadelharrat/task16/output"
)

type Task struct {
	ID          int
	Description string
	Completed   bool
}

func AddTask(description string, currentTasks *[]Task, currentNextID *int) {
	newTask := Task{
		ID:          *currentNextID,
		Description: description,
		Completed:   false,
	}
	*currentTasks = append(*currentTasks, newTask)
	*currentNextID++
	output.PrintSuccess(fmt.Sprintf("Added task: '%s'", newTask.Description))
}

func ViewTasks(currentTasks []Task, mode string) {
	var filteredTasks []Task
	switch mode {
	case "all":
		filteredTasks = currentTasks
	case "pending":
		for _, task := range currentTasks {
			if !task.Completed {
				filteredTasks = append(filteredTasks, task)
			}
		}
	case "completed":
		for _, task := range currentTasks {
			if task.Completed {
				filteredTasks = append(filteredTasks, task)
			}
		}
	default:
		output.PrintError(fmt.Sprintf("filter option '%s' not defined", mode))
		return
	}

	if len(filteredTasks) == 0 {
		if mode == "all" {
			output.PrintSuccess("No tasks yet!")
		} else {
			output.PrintSuccess(fmt.Sprintf("No %s tasks found", mode))
		}
		return
	}

	fmt.Println("====== Your Tasks ======")
	for _, task := range filteredTasks {
		status := " "
		if task.Completed {
			status = "X"
		}
		fmt.Printf("[%s]: %d: %s\n", status, task.ID, task.Description)
	}
	fmt.Println("========================")
}

func EditTask(idToEdit int, newDescription string, currentTasks *[]Task) {
	found := false

	for i, task := range *currentTasks {
		if idToEdit == task.ID {
			(*currentTasks)[i].Description = newDescription
			output.PrintSuccess(fmt.Sprintf("Task %d updated to: '%s'", idToEdit, newDescription))
			found = true
			break
		}
	}

	if !found {
		output.PrintError(fmt.Sprintf("Error: Task with ID %d not found for editing.", idToEdit))
	}
}

func CompleteTask(idToComplete int, currentTasks *[]Task) {
	found := false

	for i, task := range *currentTasks {
		if idToComplete == task.ID {
			(*currentTasks)[i].Completed = true
			output.PrintSuccess(fmt.Sprintf("Task %d ('%s') marked as complete", task.ID, task.Description))
			found = true
			break
		}
	}

	if !found {
		output.PrintError(fmt.Sprintf("Error: Task with ID %d not found.", idToComplete))
	}
}

func ClearCompletedTasks(currentTasks *[]Task) {
	var pendingTasks []Task
	clearedCount := 0

	for _, task := range *currentTasks {
		if !task.Completed {
			pendingTasks = append(pendingTasks, task)
		} else {
			clearedCount++
		}
	}
	if clearedCount == 0 {
		output.PrintSuccess(fmt.Sprintf("No completed tasks to clear."))
	} else {
		*currentTasks = pendingTasks
		output.PrintSuccess(fmt.Sprintf("%d completed task(s) cleared.", clearedCount))
	}
}

func DeleteTask(idToDelete int, currentTasks *[]Task) {
	found := false

	for i, task := range *currentTasks {
		if idToDelete == task.ID {
			*currentTasks = slices.Delete(*currentTasks, i, i+1)
			output.PrintSuccess(fmt.Sprintf("Task %d ('%s') deleted.", task.ID, task.Description))
			found = true
			break
		}
	}

	if !found {
		output.PrintError(fmt.Sprintf("Error: Task with ID %d not found.", idToDelete))
	}
}

func LoadTasksFromFile(filename string, tasks *[]Task, nextID *int) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	err = json.Unmarshal(data, tasks)
	if err != nil {
		return err
	}

	var highestID int
	for _, task := range *tasks {
		if task.ID > highestID {
			highestID = task.ID
		}
	}
	*nextID = highestID + 1

	return nil
}

func SaveTasksToFile(filename string, tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
