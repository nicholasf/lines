package producer

import (
	"io/ioutil"
	"strings"
	"time"
)

type Event struct {
	Sentence string
	When     time.Time
}

const (
	PERIOD = "."
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
	content := string(b[:])
	pieces := strings.Split(content, PERIOD)

	var sentence string

	for _, s := range pieces {

		// if there is no PERIOD, append to sentence and repeat
		if len(pieces) == 1 {
			sentence = sentence + s
			continue
		}

		// otherwise, the last piece becomes the new sentence
		// and all prior pieces become events
		sentence = pieces[len(pieces)-1]

		for i := 0; i < len(pieces)-1; i++ {
			events <- &Event{
				Sentence: pieces[i],
				When:     time.Now(),
			}
		}
	}

	events <- &Event{
		Sentence: "1 2 3",
		When:     time.Now(),
	}
}
