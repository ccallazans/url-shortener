package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const (
	ERROR_LOAD_ENV_VARIABLES = "Error loading .env file"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(ERROR_LOAD_ENV_VARIABLES)
	}
}

func main() {
	router := gin.Default()
}
