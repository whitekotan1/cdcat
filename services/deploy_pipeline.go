package services

import (
	"cdcat/types"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func BuildProjectPipeline(request types.Request) types.UserProject {
	userProject := CreateUserProject(request)

	cloneErr := CloneUserProject(userProject)
	if cloneErr != nil {
		fmt.Println("can't clone repo")
	}
	userProject.DistPath = BuildUserProject(userProject)

	return userProject
}

func DeployPipeline(distPath string, bucketName string, projectID string, cloudflareConfig *s3.Client) {

	err := UploadFolder(cloudflareConfig, bucketName, projectID, distPath)

	if err != nil {
		fmt.Printf("can't depploy %v\n", err)
		DeleteUserProject(distPath)
		return
	}
	DeleteUserProject(distPath)

}
