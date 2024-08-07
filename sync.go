package main

import (
	"fmt"
	"sync"
)

// EventEmitter is a simple event emitter system.
type EventEmitter struct {
	mu       sync.RWMutex
	listeners map[string][]func(data interface{})
}

// NewEventEmitter creates a new instance of EventEmitter.
func NewEventEmitter() *EventEmitter {
	return &EventEmitter{
		listeners: make(map[string][]func(data interface{})),
	}
}

// On registers a listener for a specific event type.
func (e *EventEmitter) On(eventType string, listener func(data interface{})) {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.listeners[eventType] = append(e.listeners[eventType], listener)
}

// Emit triggers all listeners registered for a specific event type.
func (e *EventEmitter) Emit(eventType string, data interface{}) {
	e.mu.RLock()
	defer e.mu.RUnlock()

	listeners, exists := e.listeners[eventType]
	if !exists {
		return
	}
	for _, listener := range listeners {
		go listener(data) // Run listeners in separate goroutines
	}
}

func main() {
	// Create a new event emitter
	emitter := NewEventEmitter()

	// Register listeners for the "message" event
	emitter.On("message", func(data interface{}) {
		fmt.Printf("Listener 1 received: %v\n", data)
	})

	emitter.On("message", func(data interface{}) {
		fmt.Printf("Listener 2 received: %v\n", data)
	})

	// Emit a "message" event
	emitter.Emit("message", "Hello, World!")

	// Emit another "message" event with different data
	emitter.Emit("message", "Event Emitter in Go")

	// Sleep to allow goroutines to complete
	// In a real application, consider using synchronization methods
	// or handling graceful shutdowns
	time.Sleep(1 * 
