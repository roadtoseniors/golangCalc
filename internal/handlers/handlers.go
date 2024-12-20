package handlers

import (
	"encoding/json"
	"net/http"
)

type CalculateRequest struct {
	Expression string `json:"expression"`
}

type CalcResponse struct {
	Result float64 `json:"result,omitempty"`
	Error  string  `json:"error,omitempty"`
}

func MathHandler(resp http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(resp, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(resp, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		return
	}

	defer r.Body.Close()

	var req CalculateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeJSON(resp, CalcResponse{Error: "Invalid JSON data"}, http.StatusBadRequest)
		return
	}

	result, err := calc.Calc(req.Expression)
	if err != nil {
		writeJSON(resp, CalcResponse{Error: "Expression is not valid"}, http.StatusUnprocessableEntity)
		return
	}

	writeJSON(resp, CalcResponse{Result: result}, http.StatusOK)
}

func writeJSON(w http.ResponseWriter, data CalcResponse, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
