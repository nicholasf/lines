package sender

import (
	"fmt"

	"github.com/nicholasf/lines/producer"
)

func Push(e *producer.Event, url string) {
	fmt.Println(e.Sentence)
}
