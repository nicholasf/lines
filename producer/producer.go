package producer

import (
	"fmt"
	"io/ioutil"
	"time"
)

type Event struct {
	Sentence string
	When     time.Time
}

const (
	PARSE_DISTANCE = 10000
	PERIOD         = "."
)

/*
	Accepts the name of a flat text file.
	When it finds a PERIOD it then has a sentence to build an Event.
	Sends an Event with the resulting sentence through the Event channel.

	Returns an err if
		* the file doesn't exist

	Will need some kind of recovery if the file contains endless text without a PERIOD
*/

func Read(p string) (chan *Event, error) {
	content, err := ioutil.ReadFile(p)

	if err != nil {
		return nil, err
	}

	var events = make(chan *Event, 1)

	go parse(content, events)
	return events, nil
}

func parse(b []byte, events chan *Event) {
	fmt.Printf("parsing ....")
}
