package financial

import (
	"math"
	"testing"
)

func TestSMA(t *testing.T) {
	prices := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := SMA(prices, 3)

	if result == nil {
		t.Fatal("SMA returned nil")
	}

	// First 2 values should be NaN (not enough data)
	if !math.IsNaN(result[0]) || !math.IsNaN(result[1]) {
		t.Error("First values should be NaN")
	}

	// SMA at index 2 = (1+2+3)/3 = 2
	if math.Abs(result[2]-2.0) > 1e-9 {
		t.Errorf("SMA[2] = %f, want 2.0", result[2])
	}

	// SMA at index 9 = (8+9+10)/3 = 9
	if math.Abs(result[9]-9.0) > 1e-9 {
		t.Errorf("SMA[9] = %f, want 9.0", result[9])
	}
}

func TestEMA(t *testing.T) {
	prices := []float64{10, 11, 12, 13, 14, 15}
	result := EMA(prices, 3)

	if result == nil {
		t.Fatal("EMA returned nil")
	}

	// EMA should start from index period-1
	// k = 2/(3+1) = 0.5
	// EMA[2] = SMA[2] = (10+11+12)/3 = 11
	// EMA[3] = 13 * 0.5 + 11 * 0.5 = 12
	// EMA[4] = 14 * 0.5 + 12 * 0.5 = 13

	if math.IsNaN(result[2]) {
		t.Error("EMA[2] should not be NaN")
	}
}

func TestRSI(t *testing.T) {
	// Prices that go up then down
	prices := []float64{44, 44.34, 44.09, 43.61, 44.33, 44.83, 45.10, 45.42, 45.84, 46.08,
		45.89, 46.03, 45.61, 46.28, 46.28, 46.00, 46.03, 46.41, 46.22, 45.64}

	result := RSI(prices, 14)

	if result == nil {
		t.Fatal("RSI returned nil")
	}

	// RSI should be between 0 and 100
	for i, v := range result {
		if !math.IsNaN(v) && (v < 0 || v > 100) {
			t.Errorf("RSI[%d] = %f, should be between 0 and 100", i, v)
		}
	}
}

func TestBollingerBands(t *testing.T) {
	prices := []float64{20, 21, 22, 21, 20, 21, 22, 23, 22, 21}
	bb := CalculateBollingerBands(prices, 5, 2.0)

	if bb.Middle == nil || bb.Upper == nil || bb.Lower == nil {
		t.Fatal("Bollinger Bands returned nil slices")
	}

	// After period, middle should equal SMA
	for i := 4; i < len(prices); i++ {
		if math.IsNaN(bb.Middle[i]) {
			continue
		}
		// Upper should be above middle
		if bb.Upper[i] <= bb.Middle[i] {
			t.Errorf("Upper[%d] = %f should be > Middle[%d] = %f",
				i, bb.Upper[i], i, bb.Middle[i])
		}
		// Lower should be below middle
		if bb.Lower[i] >= bb.Middle[i] {
			t.Errorf("Lower[%d] = %f should be < Middle[%d] = %f",
				i, bb.Lower[i], i, bb.Middle[i])
		}
	}
}

func TestSimpleReturns(t *testing.T) {
	prices := []float64{100, 110, 99, 105}
	returns := SimpleReturns(prices)

	if returns == nil {
		t.Fatal("SimpleReturns returned nil")
	}

	if len(returns) != 3 {
		t.Fatalf("Expected 3 returns, got %d", len(returns))
	}

	// Return 0: (110-100)/100 = 0.10
	if math.Abs(returns[0]-0.10) > 1e-9 {
		t.Errorf("returns[0] = %f, want 0.10", returns[0])
	}

	// Return 1: (99-110)/110 = -0.10
	if math.Abs(returns[1]-(-0.10)) > 1e-9 {
		t.Errorf("returns[1] = %f, want -0.10", returns[1])
	}
}

func TestLogReturns(t *testing.T) {
	prices := []float64{100, 110, 99, 105}
	returns := LogReturns(prices)

	if returns == nil {
		t.Fatal("LogReturns returned nil")
	}

	if len(returns) != 3 {
		t.Fatalf("Expected 3 returns, got %d", len(returns))
	}

	// Return 0: ln(110/100) ≈ 0.0953
	expected := math.Log(110.0 / 100.0)
	if math.Abs(returns[0]-expected) > 1e-9 {
		t.Errorf("returns[0] = %f, want %f", returns[0], expected)
	}
}

func TestCumulativeReturns(t *testing.T) {
	prices := []float64{100, 110, 99, 120}
	returns := CumulativeReturns(prices)

	if returns == nil {
		t.Fatal("CumulativeReturns returned nil")
	}

	// Cumulative return at end: (120-100)/100 = 0.20
	if math.Abs(returns[len(returns)-1]-0.20) > 1e-9 {
		t.Errorf("Final cumulative return = %f, want 0.20", returns[len(returns)-1])
	}
}
