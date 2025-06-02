package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome To GoTasks")

	var taskCount int
	var tasks []string

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("How many tasks you want to add: ")
	fmt.Scanln(&taskCount)

	for i := range taskCount {
		fmt.Printf("Enter Task %d: ", i+1)

		taskName, _ := reader.ReadString('\n') // The underscore _ is used to ignore a returned value, Which is where `err` would hold any error encountered while reading
		taskName = strings.TrimSpace(taskName) // remove newline and spaces
		tasks = append(tasks, taskName)        // append adds new element at last always
	}

	fmt.Println("\nYour Tasks")

	// forEach like loop
	for i, task := range tasks {
		fmt.Printf("%d. %s\n", i+1, task)
	}
}
