package App

import (
	"github.com/gin/demo/internal/application"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestEventDispatcher(t *testing.T) {
	dispatcher := application.NewEventDispatcher()

	// Channel to track handler execution order
	order := make(chan string, 3)

	// Sequential handler
	//dispatcher.Subscribe("TestEvent", func(event interface{}) {
	//	time.Sleep(2 * time.Second)
	//	order <- "Sequential Handler"
	//}, true)
	//
	//// Concurrent handler
	//dispatcher.Subscribe("TestEvent", func(event interface{}) {
	//	order <- "Concurrent Handler 1"
	//}, false)
	//
	//// Another concurrent handler
	//dispatcher.Subscribe("TestEvent", func(event interface{}) {
	//	order <- "Concurrent Handler 2"
	//}, false)

	// Dispatch the event
	dispatcher.Dispatch("TestEvent", nil)

	// Close the order channel after all handlers are finished
	go func() {
		time.Sleep(3 * time.Second)
		close(order)
	}()

	// Assert order of execution
	expectedOrder := []string{"Concurrent Handler 1", "Concurrent Handler 2", "Sequential Handler"}
	var actualOrder []string
	for entry := range order {
		actualOrder = append(actualOrder, entry)
	}

	assert.ElementsMatch(t, expectedOrder, actualOrder)
}
