package main

import (
	"fmt"
	"os"
	"tdsd/internal/tdsd"
)

func main() {
	projectName := os.Args[1]

	if projectName == "" {
		fmt.Println("No project name provided")
		os.Exit(1)
	}

	tdsd.SwitchProject(projectName)

	fmt.Println("Switched to project: " + projectName)
}
