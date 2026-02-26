package main

import (
	"cdcat/api"
	"cdcat/services"
	"cdcat/storage"
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("hii")
	storage.InitDB()
	cloudflareKeys := services.LoadEnv()

	cloudflareConfig := services.Initialize_R2(cloudflareKeys)

	R2Client := api.R2Client{
		CloudflareCfg: cloudflareConfig,
	}

	fmt.Println(R2Client)

	fmt.Println(cloudflareConfig)
	http.HandleFunc("/", api.HandlePage)
	http.HandleFunc("/request", R2Client.HandleRequest)

	err := http.ListenAndServe(":5000", nil)

	if err != nil {
		fmt.Println("error server runinig")
		os.Exit(1)
	}

}
