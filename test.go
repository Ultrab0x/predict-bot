package main

import (
	"math/big"
	"testing"
)

func TestPredictPrice(t *testing.T) {
	// Create a new instance of the prediction model
	model := NewPredictionModel()

	// Test the predict price function with different input values
	currentPrice := big.NewFloat(1)
	predictedPrice := model.PredictPrice(currentPrice)
	expectedPrice := big.NewFloat(1.2)
	if predictedPrice.Cmp(expectedPrice) != 0 {
		t.Errorf("Expected predictPrice to return %f, got %f", expectedPrice, predictedPrice)
	}

	currentPrice = big.NewFloat(2)
	predictedPrice = model.PredictPrice(currentPrice)
	expectedPrice = big.NewFloat(2.4)
	if predictedPrice.Cmp(expectedPrice) != 0 {
		t.Errorf("Expected predictPrice to return %f, got %f", expectedPrice, predictedPrice)
	}
}
