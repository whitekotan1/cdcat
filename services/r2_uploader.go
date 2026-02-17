package services

import (
	"fmt"

	"os"

	"io"

	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func OpenFile(fileName string) (*os.File, error) {
	openFile, openErr := os.Open(fileName)
	if openErr != nil {
		fmt.Println("can't find file", openErr)
		return nil, fmt.Errorf("can't read file %s", openErr)
	}
	return openFile, nil

}

func UploadFileToR2(client *s3.Client, bucketName string, key string, body io.Reader) error {
	input := &s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(key),
		Body:        body,
		ContentType: aws.String("application/octet-stream"),
	}

	_, err := client.PutObject(context.TODO(), input)

	if err != nil {
		fmt.Println("can't input file to r2")
	}
	fmt.Printf("uploaded file %s", key)
	return nil

}
