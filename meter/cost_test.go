package meter_test

import (
	"testing"
)

// Test for cost calculation functionality
func TestCostCalculation(t *testing.T) {
	// Example test case
	cost := CalculateCost(10, 0.5)
	expected := 5.0
	if cost != expected {
		t.Errorf("Expected %v, but got %v", expected, cost)
	}
}