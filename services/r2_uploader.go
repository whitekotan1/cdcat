package services

import (
	"fmt"

	"os"

	"github.com/joho/godotenv"
)

func R2_service() {

	fmt.Println("for r2")

	loadFromEnv := godotenv.Load(".env")

	if loadFromEnv != nil {
		fmt.Println("cant load .env")
	}

	bucketName := os.Getenv("bucketName")
	accountId := os.Getenv("accountId")
	accessKeyId := os.Getenv("accessKeyId")
	accessKeySecret := os.Getenv("accessKeySecret")

	fmt.Println(bucketName, accountId, accessKeyId, accessKeySecret)

}
