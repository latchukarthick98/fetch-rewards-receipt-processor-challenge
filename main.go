/**
*	Created by Lakshman Karthik Ramkumar (latchukarthick98) on 09/26/2023
 */

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Helper function to retrieve env variable
func getEnvVar(key string) string {

	// Load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	// Initialize Gin
	r := gin.Default()

	// Get PORT number from .env file
	port := getEnvVar("PORT")

	// routes.InitRouter(r)
	fmt.Printf("Running on Port %s \n", port)
	// Initialze Server on Port 3001
	r.Run(":" + port)
}
