package experiment

import "time"

// Observation represents the outcome of a candidate that has run.
type Observation struct {
	Duration   time.Duration
	Error      error
	Success    bool
	Name       string
	Panic      interface{}
	Value      interface{}
	CleanValue interface{}
}
