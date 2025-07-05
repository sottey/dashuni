package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/sottey/dashuni/internal/converter"
	"github.com/sottey/dashuni/internal/model"
)

// ConversionRequest represents the JSON payload for the convert endpoint
type ConversionRequest struct {
	Body    string `json:"body"`
	Mapping string `json:"mapping"`
}

// ConversionResponse represents the JSON response with the converted config
type ConversionResponse struct {
	Result string `json:"result"`
}

// NewRouter sets up all routes for the web server
func NewRouter() http.Handler {
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Post("/convert", handleConvert)
	})

	// Static file serving for the web UI can go here in the future
	// r.Handle("/*", http.FileServer(http.Dir("./static")))

	return r
}

// handleConvert is the HTTP handler for /api/convert
func handleConvert(w http.ResponseWriter, r *http.Request) {
	var req ConversionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Invalid JSON payload: %v", err), http.StatusBadRequest)
		return
	}

	// Validate mapping file exists
	if _, err := os.Stat(req.Mapping); err != nil {
		http.Error(w, fmt.Sprintf("Mapping file does not exist: %v", err), http.StatusBadRequest)
		return
	}

	// Parse canonical dashuni.json into model.Site
	site := &model.Site{}
	if err := json.Unmarshal([]byte(req.Body), site); err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse canonical JSON: %v", err), http.StatusBadRequest)
		return
	}

	// Apply the mapping template
	result, err := converter.ApplyTemplate(site, req.Mapping)
	if err != nil {
		http.Error(w, fmt.Sprintf("Template render error: %v", err), http.StatusInternalServerError)
		return
	}

	// Return the result in JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ConversionResponse{
		Result: result,
	})
}
