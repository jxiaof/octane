package utils

import (
	"fmt"
	"sync"
	"time"
)

// ProgressBar represents a simple progress bar
type ProgressBar struct {
	total   int
	current int
	mu      sync.Mutex
}

// NewProgressBar initializes a new ProgressBar
func NewProgressBar(total int) *ProgressBar {
	return &ProgressBar{total: total}
}

// Increment increases the current progress by one
func (pb *ProgressBar) Increment() {
	pb.mu.Lock()
	defer pb.mu.Unlock()
	pb.current++
	pb.render()
}

// render displays the current progress
func (pb *ProgressBar) render() {
	percentage := float64(pb.current) / float64(pb.total) * 100
	barLength := 50
	filledLength := int(float64(barLength) * percentage / 100)
	bar := fmt.Sprintf("█%s", string(make([]rune, filledLength)))
	empty := fmt.Sprintf("%s", string(make([]rune, barLength-filledLength)))
	fmt.Printf("\r[%s%s] %.2f%%", bar, empty, percentage)
	if pb.current >= pb.total {
		fmt.Println("\nDone!")
	}
}

// SimulateProgress simulates progress for demonstration purposes
func SimulateProgress(pb *ProgressBar) {
	for i := 0; i < pb.total; i++ {
		time.Sleep(100 * time.Millisecond) // Simulate work
		pb.Increment()
	}
}
