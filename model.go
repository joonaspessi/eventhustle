package main

import (
	"errors"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type PersonVote struct {
	Name  string   `json:"name"`
	Votes []string `json:"votes"`
}

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

func (v *Event) GetResult() Event {
	result := *v

	mostVotes := 0
	for _, vote := range v.Votes {
		if len(vote.People) > mostVotes {
			result.Votes = []Vote{vote}
			mostVotes = len(vote.People)
		}
	}

	return result
}

func (v *Event) getVoteDay(date string) (*Vote, error) {
	var resultVote *Vote
	var err error

	for i, vote := range v.Votes {
		if vote.Date == date {
			resultVote = &v.Votes[i]
		}
	}

	if resultVote == nil {
		err = errors.New("Vote day not found")
	}

	return resultVote, err
}

func (v *Event) AddVote(date string, name string) *Event {

	vote, err := v.getVoteDay(date)

	if err != nil {
		err = errors.New("not found")
	} else {
		vote.People = append(vote.People, name)
	}

	return v
}
