package http

import (
	"net/http"
	"encoding/json"
)

// MeterCost represents the meter cost configuration
type MeterCost struct {
	ID     string  `json:"id"`
	Amount float64 `json:"amount"`
}

// GetMeterCostHandler handles getting the meter cost configuration
func GetMeterCostHandler(w http.ResponseWriter, r *http.Request) {
	cost := MeterCost{ID: "1", Amount: 0.15} // Example data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cost)
}

// SetMeterCostHandler handles setting the meter cost configuration
func SetMeterCostHandler(w http.ResponseWriter, r *http.Request) {
	var cost MeterCost
	if err := json.NewDecoder(r.Body).Decode(&cost); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Here you would typically save the cost configuration
	w.WriteHeader(http.StatusNoContent)
}