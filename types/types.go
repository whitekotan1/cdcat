package types

type Request struct {
	RepoUrl string `json:"repoUrl"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
