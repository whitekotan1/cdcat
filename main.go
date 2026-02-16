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
	cloudflareKeys := services.LoadEnv()

	cloudflareConfig := services.Initialize_R2(cloudflareKeys)

	openFile, openErr := os.Open("C:/IT/cdcat/index.html")
	if openErr != nil {
		fmt.Println("can't find file", openErr)
		return
	}
	defer openFile.Close()

	uploadErr := services.UploadFileToR2(cloudflareConfig, "cdcat", "index.html", openFile)
	if uploadErr != nil {
		fmt.Println("can't upload to r2")
	}

	fmt.Println(cloudflareConfig)
	http.HandleFunc("/", api.HandlePage)
	http.HandleFunc("/request", api.HandleRequest)

	err := http.ListenAndServe(":5000", nil)

	if err != nil {
		fmt.Println("error server runinig")
		os.Exit(1)
	}

}
