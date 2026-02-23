package services

import (
	"fmt"
	"os"
)

func DeleteProjectFromServer(projectPath string, clonedProject string) {

	err := os.RemoveAll(projectPath)

	if err != nil {
		fmt.Println("can't remove compiled project folder", err)
	}

	err = os.RemoveAll(clonedProject)

	if err != nil {
		fmt.Println("can't remove cloned project folder")
	}
}
