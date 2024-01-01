package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	// slice
	var tasknum uint = 0
	var tasks = []string{}
	// var uinput string

	fmt.Println("\nWelcome to CatTask, a Task Manager.")
	fmt.Printf("Save your tasks here, use exit to stop writing.\n\n\n\n\n")

	for {
		tasknum += 1
		fmt.Printf("Task %v: ", tasknum)
		// fmt.Scanln(&uinput)

		// copied from stackoverflow - allows for spaces in input
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan() // use `for scanner.Scan()` to keep reading
		line := scanner.Text()

		tasks = append(tasks, line)
		if line == "exit" {
			tasks = tasks[:len(tasks)-1] // removes last element ("exit")
			fmt.Println("All tasks: ", tasks)
			break
		}
		if tasknum == 10 {
			fmt.Println("\n\nMaximum task limit reached. Come back tomorrow!")
			time.Sleep(3 * time.Second)
			fmt.Println("All tasks: ", tasks)
			break
		}
	}
}


