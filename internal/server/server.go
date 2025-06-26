package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/yusnelgg/gohound/internal/scanner"
)

func Start() {
	http.HandleFunc("/api/scan", scanHandler)
	http.Handle("/", http.FileServer(http.Dir("static")))

	log.Println("🚀 Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func scanHandler(w http.ResponseWriter, r *http.Request) {
	domain := r.URL.Query().Get("domain")
	if domain == "" {
		http.Error(w, "Missing 'domain' query parameter", http.StatusBadRequest)
		return
	}

	results := scanner.ScanSubdomains(domain)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
