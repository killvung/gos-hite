package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadFromDotEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		// By the way, you don't "throw", you "log"
		log.Fatalf("Error loading .env file")
	}
}

func getEnv(key string) string {
	return os.Getenv(key)
}

func main() {
	fmt.Println("Hello world")
	value := os.Getenv("env")
	if value == "" || value == "dev" {
		loadFromDotEnv()
	}
	println(getEnv("ENV"))
	println(getEnv("SECRET"))
}
