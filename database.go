package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

var (
	IsDrop = true
)

var dbSession *mgo.Session
var db *mgo.Collection

func init() {

	fmt.Println("Opening database connection")

	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	if IsDrop {
		err = session.DB("eventhustle").DropDatabase()
		if err != nil {
			panic(err)
		}
	}

	db = session.DB("eventhustle").C("events")

	err = db.Insert(
		&Event{
			Name:  "Polttarit",
			Dates: []string{"10-10-2015", "11-11-2015"},
			Votes: []Vote{{"10-10-2014", []string{"jorma"}}},
		},
		&Event{
			Name:  "Häät",
			Dates: []string{"10-10-2014", "11-11-2015"},
		},
		&Event{
			Name:  "Hautajaiset",
			Dates: []string{"10-10-2014", "11-11-2015"},
		},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("database connection established")
}

func CreateEvent(t Event) Event {
	err := db.Insert(t)

	if err != nil {
		panic(err)
	}

	var result Event
	err = db.Find(t).One(&result)

	if err != nil {
		panic(err)
	}

	return result
}

func GetEvents() []Event {
	var results []Event
	err := db.Find(nil).All(&results)

	if err != nil {
		panic(err)
	}
	return results
}

func GetEvent(id string) Event {
	var result Event
	err := db.FindId(bson.ObjectIdHex(id)).One(&result)

	if err != nil {
		log.Fatal(err)
	}
	return result
}

func DBClose() {
	dbSession.Close()
}
