package main

type omit *struct{}

// Event info presentation
type EventInfo struct {
	*Event
	Dates omit `json:"dates,omitempty"`
	Votes omit `json:"votes,omitempty"`
}

// Event result presentation
type EventResult struct {
	*Event
	Dates        omit   `json:"dates,omitempty"`
	Votes        omit   `json:"votes,omitempty"`
	SuitableDays *Votes `json:"suitableDates"`
}
