package main

import (
	"cdcat/api"
	"cdcat/services"
	"cdcat/types"
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("hii")
	cloudflareKeys := services.LoadEnv()

	cloudflareConfig := services.Initialize_R2(cloudflareKeys)

	R2Client := types.R2Client{
		CloudflareCfg: cloudflareConfig,
	}

	fmt.Println(R2Client)

	//services.DeployPipeline("C:/IT/cdcat/dist/index.html", "cdcat", "myFIle.html", )

	fmt.Println(cloudflareConfig)
	http.HandleFunc("/", api.HandlePage)
	http.HandleFunc("/request", api.HandleRequest)

	err := http.ListenAndServe(":5000", nil)

	if err != nil {
		fmt.Println("error server runinig")
		os.Exit(1)
	}

}
