package financial

import (
	"math"
	"testing"
)

func TestCalculateRebalance(t *testing.T) {
	t.Run("basic rebalance", func(t *testing.T) {
		positions := []Position{
			{Ticker: "AAPL", Value: 5000},
			{Ticker: "GOOGL", Value: 3000},
			{Ticker: "CASH", Value: 2000},
		}
		targets := map[string]float64{
			"AAPL":  0.4,
			"GOOGL": 0.4,
			"CASH":  0.2,
		}

		trades, err := CalculateRebalance(positions, targets, RebalanceConfig{})
		if err != nil {
			t.Fatalf("CalculateRebalance returned error: %v", err)
		}

		// Total portfolio = $10,000
		// Target AAPL = $4,000 (currently $5,000) -> sell $1,000
		// Target GOOGL = $4,000 (currently $3,000) -> buy $1,000

		for _, trade := range trades {
			switch trade.Ticker {
			case "AAPL":
				if math.Abs(trade.Amount-(-1000)) > 1 {
					t.Errorf("AAPL trade = %f, want -1000", trade.Amount)
				}
			case "GOOGL":
				if math.Abs(trade.Amount-1000) > 1 {
					t.Errorf("GOOGL trade = %f, want 1000", trade.Amount)
				}
			}
		}
	})

	t.Run("respects minimum trade size", func(t *testing.T) {
		positions := []Position{
			{Ticker: "AAPL", Value: 1000},
			{Ticker: "GOOGL", Value: 990},
		}
		targets := map[string]float64{
			"AAPL":  0.5,
			"GOOGL": 0.5,
		}

		trades, _ := CalculateRebalance(positions, targets, RebalanceConfig{
			MinTradeSize: 50,
		})

		// Difference is only $5 each way, below min trade size
		if len(trades) != 0 {
			t.Errorf("Expected no trades below minimum, got %d", len(trades))
		}
	})

	t.Run("invalid allocations error", func(t *testing.T) {
		positions := []Position{{Ticker: "AAPL", Value: 1000}}
		targets := map[string]float64{"AAPL": 0.5} // Only sums to 0.5

		_, err := CalculateRebalance(positions, targets, RebalanceConfig{})
		if err != ErrInvalidAllocation {
			t.Errorf("Expected ErrInvalidAllocation, got %v", err)
		}
	})
}

func TestOrderBook(t *testing.T) {
	t.Run("matches crossing orders", func(t *testing.T) {
		ob := NewOrderBook("AAPL")
		if ob == nil {
			t.Fatal("NewOrderBook returned nil")
		}

		// Add sell order at 100
		ob.AddOrder(Order{ID: "s1", Side: Sell, Price: 100.0, Quantity: 10, Timestamp: 1})

		// Add buy order at 100 - should match
		fills := ob.AddOrder(Order{ID: "b1", Side: Buy, Price: 100.0, Quantity: 5, Timestamp: 2})

		if len(fills) != 1 {
			t.Fatalf("Expected 1 fill, got %d", len(fills))
		}

		if fills[0].Quantity != 5 {
			t.Errorf("Fill quantity = %d, want 5", fills[0].Quantity)
		}
		if fills[0].Price != 100.0 {
			t.Errorf("Fill price = %f, want 100.0", fills[0].Price)
		}
	})

	t.Run("best bid and ask", func(t *testing.T) {
		ob := NewOrderBook("AAPL")
		if ob == nil {
			t.Fatal("NewOrderBook returned nil")
		}

		ob.AddOrder(Order{ID: "b1", Side: Buy, Price: 99.0, Quantity: 10, Timestamp: 1})
		ob.AddOrder(Order{ID: "b2", Side: Buy, Price: 98.0, Quantity: 10, Timestamp: 2})
		ob.AddOrder(Order{ID: "s1", Side: Sell, Price: 101.0, Quantity: 10, Timestamp: 3})
		ob.AddOrder(Order{ID: "s2", Side: Sell, Price: 102.0, Quantity: 10, Timestamp: 4})

		bid, ok := ob.BestBid()
		if !ok || bid != 99.0 {
			t.Errorf("BestBid() = %f, want 99.0", bid)
		}

		ask, ok := ob.BestAsk()
		if !ok || ask != 101.0 {
			t.Errorf("BestAsk() = %f, want 101.0", ask)
		}

		if ob.Spread() != 2.0 {
			t.Errorf("Spread() = %f, want 2.0", ob.Spread())
		}
	})

	t.Run("cancel order", func(t *testing.T) {
		ob := NewOrderBook("AAPL")
		if ob == nil {
			t.Fatal("NewOrderBook returned nil")
		}

		ob.AddOrder(Order{ID: "b1", Side: Buy, Price: 99.0, Quantity: 10, Timestamp: 1})

		if !ob.CancelOrder("b1") {
			t.Error("CancelOrder returned false for existing order")
		}

		_, ok := ob.BestBid()
		if ok {
			t.Error("BestBid should return false after cancelling only bid")
		}
	})
}

func TestVWAPCalculator(t *testing.T) {
	t.Run("calculates VWAP correctly", func(t *testing.T) {
		v := NewVWAPCalculator()
		if v == nil {
			t.Fatal("NewVWAPCalculator returned nil")
		}

		v.AddTrade(100.0, 100) // $10,000
		v.AddTrade(102.0, 200) // $20,400
		v.AddTrade(99.0, 100)  // $9,900

		// VWAP = (10000 + 20400 + 9900) / 400 = 40300 / 400 = 100.75
		expected := 100.75
		if math.Abs(v.VWAP()-expected) > 1e-9 {
			t.Errorf("VWAP() = %f, want %f", v.VWAP(), expected)
		}

		if v.TotalVolume() != 400 {
			t.Errorf("TotalVolume() = %d, want 400", v.TotalVolume())
		}
	})

	t.Run("reset clears state", func(t *testing.T) {
		v := NewVWAPCalculator()
		if v == nil {
			t.Fatal("NewVWAPCalculator returned nil")
		}

		v.AddTrade(100.0, 100)
		v.Reset()

		if v.TotalVolume() != 0 {
			t.Errorf("TotalVolume() after reset = %d, want 0", v.TotalVolume())
		}
	})
}
