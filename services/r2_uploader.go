package services

import (
	"context"
	"fmt"
	"io"
	"mime"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func UploadFolder(client *s3.Client, bucketName string, projectID string, rootpath string) error {

	return filepath.Walk(rootpath, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		relativePath, relErr := filepath.Rel(rootpath, path)
		if relErr != nil {
			return relErr
		}

		key := filepath.ToSlash(filepath.Join(projectID, relativePath))

		file, fileErr := OpenFile(path)
		if fileErr != nil {
			return fileErr
		}
		defer file.Close()

		return UploadFileToR2(client, bucketName, key, file)
	})
}

func OpenFile(fileName string) (*os.File, error) {

	openFile, openErr := os.Open(fileName)
	if openErr != nil {
		fmt.Println("can't find file", openErr)
		return nil, fmt.Errorf("can't read file %v", openErr)
	}

	return openFile, nil
}

func UploadFileToR2(client *s3.Client, bucketName string, key string, body io.Reader) error {

	fileExt := filepath.Ext(key)
	contentType := MimeTypifier(fileExt)
	input := &s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(key),
		Body:        body,
		ContentType: aws.String(contentType),
	}

	_, err := client.PutObject(context.TODO(), input)
	if err != nil {
		fmt.Println("can't input file to r2:", err)
		return err
	}

	fmt.Println("uploaded file:", key)
	return nil
}

func MimeTypifier(fileName string) string {

	fileExtension := mime.TypeByExtension(fileName)

	return fileExtension

}
