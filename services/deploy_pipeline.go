package services

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

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
