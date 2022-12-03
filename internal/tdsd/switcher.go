package tdsd

import (
	"os"
)

func SwitchProject(name string) {
	reader := NewFinder()

	config := reader.GetConfig(name)

	replaceFile(config.Path)
}

func replaceFile(filePath string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	origin, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(homeDir+"/.kube/config", origin, 0644)

	if err != nil {
		panic(err)
	}
}
