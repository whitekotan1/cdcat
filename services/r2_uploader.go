package services

import (
	"encoding/json"
	"fmt"

	"os"

	"github.com/joho/godotenv"

	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func LoadEnv() {

	fmt.Println("for r2")

	loadFromEnv := godotenv.Load(".env")

	if loadFromEnv != nil {
		fmt.Println("cant load .env")
	}

	bucketName := os.Getenv("bucketName")
	accountId := os.Getenv("accountId")
	accessKeyId := os.Getenv("accessKeyId")
	accessKeySecret := os.Getenv("accessKeySecret")

	fmt.Println(bucketName, accountId, accessKeyId, accessKeySecret)

}

func Initialize_R2() {

	fmt.Println("for r2")

	loadFromEnv := godotenv.Load(".env")

	if loadFromEnv != nil {
		fmt.Println("cant load .env")
	}

	bucketName := os.Getenv("bucketName")
	accountId := os.Getenv("accountId")
	accessKeyId := os.Getenv("accessKeyId")
	accessKeySecret := os.Getenv("accessKeySecret")

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyId, accessKeySecret, "")),
		config.WithRegion("auto"),
	)

	if err != nil {
		fmt.Println("cant initialize config")
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountId))
	})

	listObjectsOutput, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: &bucketName,
	})

	if err != nil {
		fmt.Println("cat list objects")
	}

	for _, object := range listObjectsOutput.Contents {
		obj, _ := json.MarshalIndent(object, "", "\t")
		fmt.Println(string(obj))
	}

}
