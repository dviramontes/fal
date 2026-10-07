package concurrency

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"
)

func TestBroker(t *testing.T) {
	t.Run("basic pub/sub", func(t *testing.T) {
		broker := NewBroker()
		if broker == nil {
			t.Fatal("NewBroker returned nil")
		}
		defer broker.Close()

		sub := broker.Subscribe("AAPL")

		go func() {
			time.Sleep(10 * time.Millisecond)
			broker.Publish("AAPL", 150.0)
		}()

		select {
		case msg := <-sub:
			if msg.Topic != "AAPL" {
				t.Errorf("Topic = %s, want AAPL", msg.Topic)
			}
			if msg.Payload.(float64) != 150.0 {
				t.Errorf("Payload = %v, want 150.0", msg.Payload)
			}
		case <-time.After(time.Second):
			t.Error("Did not receive message")
		}
	})

	t.Run("multiple subscribers", func(t *testing.T) {
		broker := NewBroker()
		if broker == nil {
			t.Fatal("NewBroker returned nil")
		}
		defer broker.Close()

		sub1 := broker.Subscribe("AAPL")
		sub2 := broker.Subscribe("AAPL")

		go func() {
			time.Sleep(10 * time.Millisecond)
			broker.Publish("AAPL", 150.0)
		}()

		var wg sync.WaitGroup
		wg.Add(2)

		received := make([]bool, 2)

		go func() {
			defer wg.Done()
			select {
			case <-sub1:
				received[0] = true
			case <-time.After(time.Second):
			}
		}()

		go func() {
			defer wg.Done()
			select {
			case <-sub2:
				received[1] = true
			case <-time.After(time.Second):
			}
		}()

		wg.Wait()

		if !received[0] || !received[1] {
			t.Error("Not all subscribers received message")
		}
	})

	t.Run("unsubscribe", func(t *testing.T) {
		broker := NewBroker()
		if broker == nil {
			t.Fatal("NewBroker returned nil")
		}
		defer broker.Close()

		sub := broker.Subscribe("AAPL")
		broker.Unsubscribe("AAPL", sub)

		broker.Publish("AAPL", 150.0)

		select {
		case <-sub:
			t.Error("Received message after unsubscribe")
		case <-time.After(50 * time.Millisecond):
			// Expected
		}
	})
}

func TestCircuitBreaker(t *testing.T) {
	t.Run("closed state allows calls", func(t *testing.T) {
		cb := NewCircuitBreaker(3, 1000)
		if cb == nil {
			t.Fatal("NewCircuitBreaker returned nil")
		}

		ctx := context.Background()
		result, err := cb.Execute(ctx, func() (interface{}, error) {
			return "success", nil
		})

		if err != nil {
			t.Errorf("Execute returned error: %v", err)
		}
		if result != "success" {
			t.Errorf("Result = %v, want success", result)
		}
		if cb.State() != StateClosed {
			t.Errorf("State = %v, want StateClosed", cb.State())
		}
	})

	t.Run("opens after max failures", func(t *testing.T) {
		cb := NewCircuitBreaker(3, 1000)
		if cb == nil {
			t.Fatal("NewCircuitBreaker returned nil")
		}

		ctx := context.Background()
		testErr := errors.New("test error")

		// Fail 3 times
		for i := 0; i < 3; i++ {
			cb.Execute(ctx, func() (interface{}, error) {
				return nil, testErr
			})
		}

		if cb.State() != StateOpen {
			t.Errorf("State = %v, want StateOpen after %d failures", cb.State(), 3)
		}

		// Next call should fail immediately
		_, err := cb.Execute(ctx, func() (interface{}, error) {
			return "should not run", nil
		})

		if err == nil {
			t.Error("Execute should return error when circuit is open")
		}
	})

	t.Run("half-open after timeout", func(t *testing.T) {
		cb := NewCircuitBreaker(1, 50) // 50ms timeout
		if cb == nil {
			t.Fatal("NewCircuitBreaker returned nil")
		}

		ctx := context.Background()

		// Trip the breaker
		cb.Execute(ctx, func() (interface{}, error) {
			return nil, errors.New("fail")
		})

		if cb.State() != StateOpen {
			t.Fatal("Expected circuit to be open")
		}

		// Wait for timeout
		time.Sleep(60 * time.Millisecond)

		// Should transition to half-open
		if cb.State() != StateHalfOpen {
			t.Errorf("State = %v, want StateHalfOpen after timeout", cb.State())
		}

		// Successful call should close circuit
		cb.Execute(ctx, func() (interface{}, error) {
			return "success", nil
		})

		if cb.State() != StateClosed {
			t.Errorf("State = %v, want StateClosed after success in half-open", cb.State())
		}
	})
}
