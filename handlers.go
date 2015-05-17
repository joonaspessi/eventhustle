package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to eventshuffle!")
}

func EventCreate(w http.ResponseWriter, r *http.Request) {
	var event Event
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &event); err != nil {
		w.Header().Set("Content-type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	t := CreateEvent(event)
	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

func EventIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	events := GetEvents()

	eventInfos := make([]EventInfo, len(events), len(events))

	for i, _ := range events {
		eventInfos[i] = EventInfo{Event: &events[i]}
	}

	if err := json.NewEncoder(w).Encode(eventInfos); err != nil {
		panic(err)
	}
}

func EventShow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	eventId := vars["eventId"]

	event, err := GetEvent(eventId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(event); err != nil {
		panic(err)
	}
}

func EventResults(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	eventId := vars["eventId"]

	event, err := GetEventResult(eventId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	event = event.GetResult()
	eventResult := EventResult{Event: &event, SuitableDays: &event.Votes}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(eventResult); err != nil {
		panic(err)
	}
}

func EventVote(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	eventId := vars["eventId"]

	var personVote PersonVote
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &personVote); err != nil {
		w.Header().Set("Content-type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	t, err := AddEventVote(eventId, &personVote)
	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}
