package types

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
}

type R2Config struct {
	BucketName      string
	AccountID       string
	AccessKeyID     string
	AccessKeySecret string
}
