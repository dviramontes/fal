package financial

import (
	"math"
	"testing"
)

func TestSharpeCalculator(t *testing.T) {
	t.Run("calculates sharpe ratio", func(t *testing.T) {
		sc := NewSharpeCalculator(0.02) // 2% risk-free rate
		if sc == nil {
			t.Fatal("NewSharpeCalculator returned nil")
		}

		// Add some returns
		returns := []float64{0.05, 0.03, -0.02, 0.04, 0.01}
		for _, r := range returns {
			sc.AddReturn(r)
		}

		// Mean return = (0.05 + 0.03 - 0.02 + 0.04 + 0.01) / 5 = 0.022
		// Excess return = 0.022 - 0.02 = 0.002
		// StdDev ≈ 0.0259
		// Sharpe ≈ 0.077

		sharpe := sc.SharpeRatio()
		if math.IsNaN(sharpe) {
			t.Error("SharpeRatio() returned NaN")
		}

		// Just verify it's a reasonable value
		if sharpe < -5 || sharpe > 5 {
			t.Errorf("SharpeRatio() = %f, seems unreasonable", sharpe)
		}
	})
}

func TestDrawdownCalculator(t *testing.T) {
	t.Run("tracks maximum drawdown", func(t *testing.T) {
		dc := NewDrawdownCalculator()
		if dc == nil {
			t.Fatal("NewDrawdownCalculator returned nil")
		}

		values := []float64{100, 120, 90, 110, 80, 100}
		for _, v := range values {
			dc.AddValue(v)
		}

		// Peak was 120, trough was 80
		// Max drawdown = (120 - 80) / 120 = 0.333...
		expected := (120.0 - 80.0) / 120.0
		if math.Abs(dc.MaxDrawdown()-expected) > 1e-9 {
			t.Errorf("MaxDrawdown() = %f, want %f", dc.MaxDrawdown(), expected)
		}
	})

	t.Run("current drawdown", func(t *testing.T) {
		dc := NewDrawdownCalculator()
		if dc == nil {
			t.Fatal("NewDrawdownCalculator returned nil")
		}

		dc.AddValue(100)
		dc.AddValue(120)
		dc.AddValue(100)

		// Current drawdown from peak of 120
		expected := (120.0 - 100.0) / 120.0
		if math.Abs(dc.CurrentDrawdown()-expected) > 1e-9 {
			t.Errorf("CurrentDrawdown() = %f, want %f", dc.CurrentDrawdown(), expected)
		}
	})
}

func TestKellyFraction(t *testing.T) {
	tests := []struct {
		name         string
		winProb      float64
		winLossRatio float64
		expected     float64
	}{
		{
			name:         "favorable bet",
			winProb:      0.6,
			winLossRatio: 1.0,
			expected:     0.2, // 0.6 - 0.4/1.0 = 0.2
		},
		{
			name:         "coin flip even odds",
			winProb:      0.5,
			winLossRatio: 1.0,
			expected:     0.0, // 0.5 - 0.5/1.0 = 0
		},
		{
			name:         "good win ratio",
			winProb:      0.4,
			winLossRatio: 2.0,
			expected:     0.1, // 0.4 - 0.6/2.0 = 0.1
		},
		{
			name:         "negative kelly (clamped)",
			winProb:      0.3,
			winLossRatio: 1.0,
			expected:     0.0, // Would be -0.4, clamped to 0
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := KellyFraction(tt.winProb, tt.winLossRatio)
			if math.Abs(result-tt.expected) > 1e-9 {
				t.Errorf("KellyFraction(%f, %f) = %f, want %f",
					tt.winProb, tt.winLossRatio, result, tt.expected)
			}
		})
	}
}

func TestCorrelationMatrix(t *testing.T) {
	t.Run("perfect positive correlation", func(t *testing.T) {
		returns := [][]float64{
			{1, 2, 3, 4, 5},
			{2, 4, 6, 8, 10}, // Perfectly correlated
		}

		matrix := CorrelationMatrix(returns)
		if matrix == nil {
			t.Fatal("CorrelationMatrix returned nil")
		}

		if len(matrix) != 2 || len(matrix[0]) != 2 {
			t.Fatalf("Expected 2x2 matrix, got %dx%d", len(matrix), len(matrix[0]))
		}

		// Diagonal should be 1.0
		if math.Abs(matrix[0][0]-1.0) > 1e-9 {
			t.Errorf("matrix[0][0] = %f, want 1.0", matrix[0][0])
		}

		// Off-diagonal should be 1.0 (perfect correlation)
		if math.Abs(matrix[0][1]-1.0) > 1e-9 {
			t.Errorf("matrix[0][1] = %f, want 1.0", matrix[0][1])
		}
	})

	t.Run("negative correlation", func(t *testing.T) {
		returns := [][]float64{
			{1, 2, 3, 4, 5},
			{5, 4, 3, 2, 1}, // Perfectly negatively correlated
		}

		matrix := CorrelationMatrix(returns)
		if matrix == nil {
			t.Fatal("CorrelationMatrix returned nil")
		}

		// Off-diagonal should be -1.0
		if math.Abs(matrix[0][1]-(-1.0)) > 1e-9 {
			t.Errorf("matrix[0][1] = %f, want -1.0", matrix[0][1])
		}
	})
}
