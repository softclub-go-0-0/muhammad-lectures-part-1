package main

import "fmt"

func main() {
	var tasksNum, correct, mistakes int
	var taskName string
	var isSolved bool
	fmt.Scan(&tasksNum)
	tasks := make(map[string]bool)
	for i := 0; i < tasksNum; i++ {
		fmt.Scan(&taskName)
		fmt.Scan(&isSolved)
		tasks[taskName] = isSolved
		if isSolved {

			correct++

		}
		if !isSolved {
			mistakes++
		}
	}
	fmt.Println("Corrects:", correct)
	fmt.Println("Mistakes:", mistakes)
}
