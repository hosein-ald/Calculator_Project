package handlers

import (
	"encoding/json"
	"net/http"
)

type MathRequest struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

type MathResponse struct {
	Result float64 `json:"result"`
}

// HandleSubtract handles POST /subtract
func HandleSubtract(w http.ResponseWriter, r *http.Request) {
	var req MathRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp := MathResponse{Result: req.A - req.B}
	writeJSON(w, resp)
}

// HandleAdd handles POST /add
func HandleAdd(w http.ResponseWriter, r *http.Request) {
	var req MathRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp := MathResponse{Result: req.A + req.B}
	writeJSON(w, resp)
}

// HandleMultiply handles POST /multiply
func HandleMultiply(w http.ResponseWriter, r *http.Request) {
	var req MathRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp := MathResponse{Result: req.A * req.B}
	writeJSON(w, resp)
}

// HandleDivide handles POST /divide
func HandleDivide(w http.ResponseWriter, r *http.Request) {
	var req MathRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.B == 0 {
		http.Error(w, "division by zero is not allowed", http.StatusBadRequest)
		return
	}

	resp := MathResponse{Result: req.A / req.B}
	writeJSON(w, resp)
}

// Helper to send JSON responses
func writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
