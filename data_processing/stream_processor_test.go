package dataprocessing

import (
	"math"
	"testing"
)

func TestWindowAggregator(t *testing.T) {
	t.Run("aggregates trades into OHLCV", func(t *testing.T) {
		wa := NewWindowAggregator(1000) // 1 second windows
		if wa == nil {
			t.Fatal("NewWindowAggregator returned nil")
		}

		// Add trades in first window (0-999ms)
		wa.AddTrade(Trade{Ticker: "AAPL", Price: 100.0, Volume: 100, Timestamp: 0})
		wa.AddTrade(Trade{Ticker: "AAPL", Price: 105.0, Volume: 200, Timestamp: 500})
		wa.AddTrade(Trade{Ticker: "AAPL", Price: 95.0, Volume: 150, Timestamp: 800})

		// Add trade in next window to complete first
		wa.AddTrade(Trade{Ticker: "AAPL", Price: 102.0, Volume: 100, Timestamp: 1500})

		windows := wa.GetCompletedWindows()
		if len(windows) < 1 {
			t.Fatal("Expected at least one completed window")
		}

		w := windows[0]
		if w.Open != 100.0 {
			t.Errorf("Open = %f, want 100.0", w.Open)
		}
		if w.High != 105.0 {
			t.Errorf("High = %f, want 105.0", w.High)
		}
		if w.Low != 95.0 {
			t.Errorf("Low = %f, want 95.0", w.Low)
		}
		if w.Close != 95.0 {
			t.Errorf("Close = %f, want 95.0", w.Close)
		}
		if w.Volume != 450 {
			t.Errorf("Volume = %d, want 450", w.Volume)
		}
	})
}

func TestRollingStats(t *testing.T) {
	t.Run("calculates mean correctly", func(t *testing.T) {
		rs := NewRollingStats(3)
		if rs == nil {
			t.Fatal("NewRollingStats returned nil")
		}

		rs.Add(1.0)
		rs.Add(2.0)
		rs.Add(3.0)

		if math.Abs(rs.Mean()-2.0) > 1e-9 {
			t.Errorf("Mean() = %f, want 2.0", rs.Mean())
		}
	})

	t.Run("rolling window drops old values", func(t *testing.T) {
		rs := NewRollingStats(3)
		if rs == nil {
			t.Fatal("NewRollingStats returned nil")
		}

		rs.Add(1.0)
		rs.Add(2.0)
		rs.Add(3.0)
		rs.Add(4.0) // Should drop 1.0

		// Mean should now be (2+3+4)/3 = 3.0
		if math.Abs(rs.Mean()-3.0) > 1e-9 {
			t.Errorf("Mean() = %f, want 3.0", rs.Mean())
		}
	})

	t.Run("calculates std dev", func(t *testing.T) {
		rs := NewRollingStats(4)
		if rs == nil {
			t.Fatal("NewRollingStats returned nil")
		}

		rs.Add(2.0)
		rs.Add(4.0)
		rs.Add(4.0)
		rs.Add(4.0)
		rs.Add(5.0)
		rs.Add(5.0)
		rs.Add(7.0)
		rs.Add(9.0)

		// Last 4 values: 5, 5, 7, 9
		// Mean = 6.5, Variance = 2.917, StdDev ≈ 1.708
		stdDev := rs.StdDev()
		if math.Abs(stdDev-1.708) > 0.01 {
			t.Errorf("StdDev() = %f, want ~1.708", stdDev)
		}
	})

	t.Run("tracks min and max", func(t *testing.T) {
		rs := NewRollingStats(3)
		if rs == nil {
			t.Fatal("NewRollingStats returned nil")
		}

		rs.Add(5.0)
		rs.Add(2.0)
		rs.Add(8.0)

		if rs.Min() != 2.0 {
			t.Errorf("Min() = %f, want 2.0", rs.Min())
		}
		if rs.Max() != 8.0 {
			t.Errorf("Max() = %f, want 8.0", rs.Max())
		}

		rs.Add(10.0) // Drops 5.0

		if rs.Min() != 2.0 {
			t.Errorf("Min() = %f, want 2.0", rs.Min())
		}
		if rs.Max() != 10.0 {
			t.Errorf("Max() = %f, want 10.0", rs.Max())
		}

		rs.Add(6.0) // Drops 2.0

		if rs.Min() != 6.0 {
			t.Errorf("Min() = %f, want 6.0 after dropping 2.0", rs.Min())
		}
	})
}

func TestDeduplicator(t *testing.T) {
	t.Run("detects duplicates", func(t *testing.T) {
		d := NewDeduplicator(1000) // 1 second window
		if d == nil {
			t.Fatal("NewDeduplicator returned nil")
		}

		if d.IsDuplicate(1, 100) {
			t.Error("First occurrence should not be duplicate")
		}
		if !d.IsDuplicate(1, 200) {
			t.Error("Second occurrence should be duplicate")
		}
	})

	t.Run("cleanup removes old entries", func(t *testing.T) {
		d := NewDeduplicator(1000)
		if d == nil {
			t.Fatal("NewDeduplicator returned nil")
		}

		d.IsDuplicate(1, 100)
		d.Cleanup(2000) // 2 seconds later

		// Same seq should not be duplicate after cleanup
		if d.IsDuplicate(1, 2100) {
			t.Error("Should not be duplicate after cleanup")
		}
	})
}
