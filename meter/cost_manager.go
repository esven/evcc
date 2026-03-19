// CostManager is responsible for calculating and managing costs related to various operations.

package meter

import (
	"errors"
)

type CostManager struct {
	costs map[string]float64
}

// NewCostManager creates a new CostManager instance.
func NewCostManager() *CostManager {
	return &CostManager{costs: make(map[string]float64)}
}

// SetCost sets the cost for a specified operation.
func (cm *CostManager) SetCost(operation string, cost float64) {
	cm.costs[operation] = cost
}

// GetCost retrieves the cost for a specified operation.
func (cm *CostManager) GetCost(operation string) (float64, error) {
	cost, exists := cm.costs[operation]
	if !exists {
		return 0, errors.New("cost not set for operation: " + operation)
	}
	return cost, nil
}

// CalculateTotalCost calculates the total cost for multiple operations.
func (cm *CostManager) CalculateTotalCost(operations []string) (float64, error) {
	totalCost := 0.0
	for _, op := range operations {
		cost, err := cm.GetCost(op)
		if err != nil {
			return 0, err
		}
		totalCost += cost
	}
	return totalCost, nil
}