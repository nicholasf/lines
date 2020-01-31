# Introduction

This is a simple parsing tool that works with any text file to provide JSON events to a pet project of mine called [JSONator](https://github.com/nicholasf/jsonator). 

## Behaviour

The parser takes two environment vars - DATA_FILE and TARGET_URL.

### DATA_FILE

Any plain file. The program will split on full stop symbol and then build a JSON payload using the sentence.

For example, for a flat file resembling the following.

```
Stately, plump Buck Mulligan came from the stairhead, bearing a bowl of
lather on which a mirror and a razor lay crossed. A yellow
dressinggown, ungirdled, was sustained gently behind him on the mild
morning air. He held the bowl aloft and intoned:

—_Introibo ad altare Dei_.
```

The following JSON packets would be made:

{ 
    "order": 0,
    "sentence": "Stately, plump Buck Mulligan came from the stairhead, bearing a bowl of
lather on which a mirror and a razor lay crossed.",
    "time": "<current timestamp>",
}

* order - the position of the sentence in the file
* sentence - the sentence
* time - current UTC timestamp on the machine
* run - the number of times the program has looped through this sentence

### TARGET_URL

The endpoint URL where to post this JSON event.

## Caveat

Obviously, this program will blow up if you give it too large a file to swallow. It is primitive but all that I need for my purposes.