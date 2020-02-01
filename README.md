# Introduction

This is a play project so I can experiment with Istio and various write protocols for service communication.

A flat text file can be fed into the initial service, which converts each line of the text file into an Event

```
type Event struct {
	Sentence string    `json:"sentence"`
	When     time.Time `json:"when"`
}
```

The initiating service turns it into JSON, then sends it onto the next service. 
