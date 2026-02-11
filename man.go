package main

import (
	"fmt"
	"net/http"
)

type Request struct {
	R3AccountID string
}

type Response struct {
}

func main() {
	fmt.Println("hii")

	err := http.ListenAndServe(":5000", nil)

	if err != nil {
		fmt.Println("server runinig")
	}

}
