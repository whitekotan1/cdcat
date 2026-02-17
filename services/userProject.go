package services

import (
	"cdcat/types"
	"fmt"
	"math/rand/v2"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

func CreateUserProject(request types.Request) types.UserProject {

	var userProject types.UserProject
	userProject.ID = rand.IntN(1000)
	userProject.GihubLink = request.RepoUrl
	userProject.UserID = 0

	err := os.Mkdir(strconv.Itoa(userProject.ID), 0755)

	if err != nil {
		fmt.Println("error", err)
	}

	return userProject

}

func CloneUserProject(userProject types.UserProject) error {

	projectDir := strconv.Itoa(userProject.ID)

	cloneProjectCmd := exec.Command("git", "clone", userProject.GihubLink, projectDir)

	cloneProjectOutput, err := cloneProjectCmd.CombinedOutput()

	if err != nil {
		fmt.Println("cant clone git repo", err)
	}

	fmt.Println("success", string(cloneProjectOutput))

	return nil

}

func BuildUserProject(userProject types.UserProject) {

	projectDir := strconv.Itoa(userProject.ID)

	absolutePath, err := filepath.Abs(projectDir)
	if err != nil {
		fmt.Println("path error", err)
		return
	}

	buildUserProject := exec.Command(
		"docker", "run", "--rm",
		"-v", absolutePath+":/app",
		"-w", "/app",
		"node:20",
		"sh", "-c", "npm install && npm run build",
	)

	buildResult, err := buildUserProject.CombinedOutput()
	if err != nil {
		fmt.Println("build unsuccessful", err)
	}

	fmt.Println(string(buildResult))
}
