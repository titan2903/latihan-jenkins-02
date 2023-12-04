package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type HealthResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Code    int    `json:"code"`
}

func getEnvHandler(w http.ResponseWriter, r *http.Request) {
	envVars := os.Environ()

	// Convert environment variables to a map for easy JSON serialization
	envMap := make(map[string]string)
	for _, envVar := range envVars {
		pair := strings.SplitN(envVar, "=", 2)
		if len(pair) == 2 {
			envMap[pair[0]] = pair[1]
		}
	}

	// Convert the map to JSON
	envJSON, err := json.Marshal(envMap)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header and write the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.Write(envJSON)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// response body health check
	response := &HealthResponse{
		Message: "Health Check API",
		Status:  "Ok",
		Code:    http.StatusOK,
	}

	// Marshal the struct into JSON
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header and write the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}

func main() {
	http.HandleFunc("/getenv", getEnvHandler)
	http.HandleFunc("/health", healthCheckHandler)

	port := 5000
	fmt.Printf("Server is running on :%d\n", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
