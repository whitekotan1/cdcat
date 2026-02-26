package types

type Request struct {
	RepoUrl  string `json:"repoUrl"`
	RepoType string `json:"repoType"`
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

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Plan  string `json:"plan"`
}
