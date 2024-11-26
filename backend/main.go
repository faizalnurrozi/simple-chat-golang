package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Middleware untuk menambahkan header CORS
func enableCors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
}

func main() {
	hub := NewHub()
	go hub.Run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		enableCors(w, r)
		ServeWs(hub, w, r)
	})

	http.HandleFunc("/themes", func(w http.ResponseWriter, r *http.Request) {
		enableCors(w, r)

		// Data tema dalam array
		themes := []string{"Technology", "Education", "Entertainment", "Sports"}

		// Tetapkan header untuk JSON
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Encode data menjadi JSON dan kirimkan ke response
		err := json.NewEncoder(w).Encode(map[string][]string{"themes": themes})
		if err != nil {
			http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
			fmt.Println("Error encoding JSON:", err)
		}
	})

	fmt.Println("Server running at :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
