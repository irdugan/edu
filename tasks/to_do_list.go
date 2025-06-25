package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	const (
		create = 1
		read   = 2
		update = 3
		delete = 4
	)

	fmt.Println("Welcome to the TO DO List CLI app!")
	fmt.Println()
	fmt.Println("Enter your command (create - 1, read - 2, update - 3, delete - 4):")

	//var updatedTask string
	var taskNumber int
	tasks := make([]string, 0, 1)
	taskNumberToUpdate := 0
	taskNumberToDelete := 0

	for {
		fmt.Scan(&taskNumber)
		switch taskNumber {
		case create:
			fmt.Println("Enter task name:")
			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')
			tasks = append(tasks, input)
			fmt.Println("Enter your command (create - 1, read - 2, update - 3, delete - 4):")
			break
		case read:
			for i, values := range tasks {
				fmt.Printf("%d: %s\n", i+1, values)
			}
			fmt.Println("Enter your command (create - 1, read - 2, update - 3, delete - 4):")
			break
		case update:
			fmt.Println("Enter task number to update:")
			fmt.Scan(&taskNumberToUpdate)
			for i := 0; i < len(tasks); i++ {
				if taskNumberToUpdate == i+1 {
					fmt.Println("Enter new task:")
					readerNewTask := bufio.NewReader(os.Stdin)
					inputNewTask, _ := readerNewTask.ReadString('\n')
					tasks[i] = inputNewTask
					fmt.Println("Updated task", i+1, "saved successfully!")
					fmt.Println("Enter your command (create - 1, read - 2, update - 3, delete - 4):")
					break
				}
			}
		case delete:
			fmt.Println("Enter task number to remove:")
			fmt.Scan(&taskNumberToDelete)
			for i := 0; i < len(tasks); i++ {
				if taskNumberToDelete == i+1 {
					tasks = append(tasks[:i], tasks[i+1:]...)
					fmt.Println("Task", i+1, "removed successfully!")
					fmt.Println("Enter your command (create - 1, read - 2, update - 3, delete - 4):")
					break
				} else if taskNumberToDelete > len(tasks) {
					fmt.Println("Invalid command! Please, try again!")
					fmt.Println("Enter your command (create - 1, read - 2, update - 3, delete - 4):")
					break
				}
			}
		}
	}

}
