package types

import "github.com/aws/aws-sdk-go-v2/service/s3"

type Request struct {
	RepoUrl string `json:"repoUrl"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type UserProject struct {
	ID        int    `json:"id"`
	GihubLink string `json:"githublink"`
	UserID    int    `json:"userid"`
	DistPath  string `json:"distpath"`
}

type R2Config struct {
	BucketName      string
	AccountID       string
	AccessKeyID     string
	AccessKeySecret string
}

type R2Client struct {
	CloudflareCfg *s3.Client
}
