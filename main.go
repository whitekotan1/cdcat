package main

import (
	"encoding/json"
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
	http.HandleFunc("/request", handleRequest)

	err := http.ListenAndServe(":5000", nil)

	if err != nil {
		fmt.Println("error server runinig")
		os.Exit(1)
	}

}

func handlePage(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "this is GET methd!", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "index.html")

}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "only post", http.StatusMethodNotAllowed)
		return
	}

	var response Response

	err := json.NewDecoder(r.Body).Decode(&response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
