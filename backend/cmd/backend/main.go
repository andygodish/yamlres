// This should go in cmd/backend/main.go
package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/andygodish/yamlres/backend/internal/routes"
	"github.com/andygodish/yamlres/backend/internal/service"
)

func main() {
	// Check for primary resume file
	primaryPath := "./config/resume.yaml"
	fallbackPath := "./config/sample-resume.yaml"

	// Determine which resume file to use
	resumeFilePath := primaryPath
	if _, err := os.Stat(primaryPath); os.IsNotExist(err) {
		log.Printf("Primary resume file not found at %s, using fallback", primaryPath)
		resumeFilePath = fallbackPath
	}

	// Create absolute path
	absPath, err := filepath.Abs(resumeFilePath)
	if err != nil {
		log.Fatalf("Failed to get absolute path for resume file: %v", err)
	}

	log.Printf("Using resume file: %s", absPath)

	// Create resume service
	resumeService := service.NewResumeService(absPath)

	// Initialize router
	router := routes.InitRoutes(resumeService)

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the server
	log.Printf("Starting server on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
