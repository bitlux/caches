package util

import (
	"fmt"
	"maps"
	"os"
	"slices"
	"strings"
)

// This API is experimental.

// ModeSelector
// reads argv[1]
// It is not safe for concurrent access from multiple goroutines.
type ModeSelector struct {
	m map[string]func() error
}

// NewModeSelector creates a ModeSelector.
func NewModeSelector() *ModeSelector {
	return &ModeSelector{m: map[string]func() error{}}
}

// Add defines a new mode.
func (ms *ModeSelector) Add(name string, f func() error) {
	ms.m[name] = f
}

func (ms *ModeSelector) usage(bad string) {
	modes := slices.Collect(maps.Keys(ms.m))
	fmt.Printf("Invalid mode: '%s'. Options are: %s\n", bad, strings.Join(modes, ", "))
	os.Exit(1)
}

func (ms *ModeSelector) Run() error {
	if len(os.Args) < 2 {
		fmt.Printf("Mode not supplied")
		os.Exit(1)
	}
	name := os.Args[1]
	f, ok := ms.m[name]
	if !ok {
		ms.usage(name)
	}
	return f()
}
