package main

import (
	api "cdcat/handlers"
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("hii")

	http.HandleFunc("/", api.HandlePage)
	http.HandleFunc("/request", api.HandleRequest)

	err := http.ListenAndServe(":5000", nil)

	if err != nil {
		fmt.Println("error server runinig")
		os.Exit(1)
	}

}
