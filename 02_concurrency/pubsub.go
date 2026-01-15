package concurrency

import "context"

// Exercise 6: Pub/Sub Message Broker
// Implement a simple in-memory pub/sub system for distributing market updates.
// This pattern is core to event-driven architectures.
//
// Requirements:
// - Multiple subscribers can subscribe to topics
// - Publishers can publish messages to topics
// - Messages are delivered to all subscribers of that topic
// - Subscribers can unsubscribe
// - Thread-safe operations
//
// Example:
//   broker := NewBroker()
//   sub := broker.Subscribe("AAPL")
//   broker.Publish("AAPL", PriceUpdate{Price: 150.0})
//   msg := <-sub // receives the update

type Message struct {
	Topic   string
	Payload interface{}
}

type Broker struct {
	// TODO: Add fields
}

func NewBroker() *Broker {
	// TODO: Implement this
	return nil
}

func (b *Broker) Subscribe(topic string) <-chan Message {
	// TODO: Implement this
	return nil
}

func (b *Broker) Unsubscribe(topic string, ch <-chan Message) {
	// TODO: Implement this
}

func (b *Broker) Publish(topic string, payload interface{}) {
	// TODO: Implement this
}

func (b *Broker) Close() {
	// TODO: Implement this
}

// Exercise 7: Circuit Breaker
// Implement a circuit breaker pattern for external API calls.
// Essential for building resilient distributed systems.
//
// States:
// - Closed: requests pass through, failures are counted
// - Open: requests fail immediately, no calls to external service
// - Half-Open: allow limited requests to test if service recovered
//
// Example:
//   cb := NewCircuitBreaker(5, time.Second*30) // 5 failures, 30s timeout
//   result, err := cb.Execute(func() (interface{}, error) {
//       return apiClient.FetchData()
//   })

type CircuitState int

const (
	StateClosed CircuitState = iota
	StateOpen
	StateHalfOpen
)

type CircuitBreaker struct {
	// TODO: Add fields
}

func NewCircuitBreaker(maxFailures int, timeout int64) *CircuitBreaker {
	// TODO: Implement this
	return nil
}

func (cb *CircuitBreaker) Execute(ctx context.Context, fn func() (interface{}, error)) (interface{}, error) {
	// TODO: Implement this
	return nil, nil
}

func (cb *CircuitBreaker) State() CircuitState {
	// TODO: Implement this
	return StateClosed
}
