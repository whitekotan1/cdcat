package main

import (
	"fmt"

	"net/http"
	"os"
)

type Request struct {
	RepoUrl string `json:"repoUrl"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	fmt.Println("hii")

	http.HandleFunc("/", handlePage)

	err := http.ListenAndServe(":5000", nil)

	if err != nil {
		fmt.Println("error server runinig")
		os.Exit(1)
	}

}

func handlePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")

}
