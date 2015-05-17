package main

import (
	"gopkg.in/mgo.v2/bson"
)

type Vote struct {
	Date   string   `json:"date"`
	People []string `json:"people"`
}

type Votes []Vote

type Event struct {
	Id    bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name  string        `json:"name"`
	Dates []string      `json:"dates"`
	Votes Votes         `json:"votes,omitempty"`
}

type Events []Event

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
