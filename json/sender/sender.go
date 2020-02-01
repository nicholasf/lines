package sender

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/nicholasf/lines/json/producer"
)

func Push(events chan *producer.Event, url string) {

	for {
		e := <-events
		b, err := json.Marshal(*e)
		payload := bytes.NewBuffer(b)

		resp, err := http.Post("http://127.0.0.1:1323/json", "application/json", payload)

		if err != nil {
			log.Println("Error communicating with localhost:1323: ", err.Error())
			log.Println("On ", e.Sentence)
		}

		if resp != nil {
			if resp.StatusCode != 200 {
				log.Println("non 200 from receiver %d", resp.StatusCode)
			}
		}
	}
}
