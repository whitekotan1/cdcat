package services

import (
	"encoding/json"
	"fmt"

	"os"

	"cdcat/types"

	"github.com/joho/godotenv"

	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func LoadEnv() types.R2Config {

	fmt.Println("for r2")

	loadFromEnv := godotenv.Load(".env")

	if loadFromEnv != nil {
		fmt.Println("cant load .env")
	}

	r2_cfg := types.R2Config{

		BucketName:      os.Getenv("bucketName"),
		AccountID:       os.Getenv("accountId"),
		AccessKeyID:     os.Getenv("accessKeyId"),
		AccessKeySecret: os.Getenv("accessKeySecret"),
	}

	return r2_cfg

}

func Initialize_R2(r2_cfg types.R2Config) *s3.Client {

	bucketName := r2_cfg.BucketName
	accountId := r2_cfg.AccountID
	accessKeyId := r2_cfg.AccessKeyID
	accessKeySecret := r2_cfg.AccessKeySecret

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

	return client

}
