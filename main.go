package main

import (
	"cdcat/api"
	"cdcat/services"
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("hii")
	cloudflareConfig := services.LoadEnv()

	services.Initialize_R2(cloudflareConfig)
	http.HandleFunc("/", api.HandlePage)
	http.HandleFunc("/request", api.HandleRequest)

	err := http.ListenAndServe(":5000", nil)

	if err != nil {
		fmt.Println("error server runinig")
		os.Exit(1)
	}

}
