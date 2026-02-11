package services

import (
	"cdcat/types"
	"fmt"
	"math/rand/v2"
	"os"
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
