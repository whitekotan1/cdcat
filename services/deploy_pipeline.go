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
	BuildUserProject(userProject)

	return userProject
}

func DeployPipeline(filePath string, bucketName string, fileName string, cloudflareConfig *s3.Client) {

	openFile, openErr := OpenFile(filePath)
	if openErr != nil {
		fmt.Println("can't find file", openErr)
		return
	}
	defer openFile.Close()

	uploadErr := UploadFileToR2(cloudflareConfig, bucketName, fileName, openFile)
	if uploadErr != nil {
		fmt.Println("can't upload to r2")
	}

}
