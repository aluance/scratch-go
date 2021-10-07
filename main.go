package main

import (
	"fmt"

	"github.com/aluance/scratch-go/internal/tasks"
)

func main() {
	fmt.Println("Hello!")

	if err := tasks.CauseModify(); err != nil {
		panic("Something went wrong!")
	}

	// if err := tasks.DoSomething(); err != nil {
	// 	panic("Something went wrong!")
	// }
}
