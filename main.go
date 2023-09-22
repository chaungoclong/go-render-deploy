package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "success",
			"data":   "Hello, world",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	serveErr := http.ListenAndServe(":"+port, nil)
	if serveErr != nil {
		fmt.Println("Error: ", serveErr)
		return
	}
}
