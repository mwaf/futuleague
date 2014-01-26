# FutuLeague

Is a personal tournament and ranking system for sports video games (e.g. FIFA14, NHL14). Is is bulit with go and was originally started during a TDD hackaton at Futurice's Berlin office.

## Requirements

* [Go](http://golang.org/doc/install)

### Dependencies

* github.com/mattn/go-sqlite3
* github.com/gorilla/mux

## Building and running

* `$ go build`
* `$ ./futuleague`

See `$ ./futuleague -h` for additional options.

## Running tests

* `$ go test`

## Terminology

There is some ambigiouty in the terms so below are some definitions.

- *Game* - the actual video game, e.g. FIFA14
- *Player* - human being playing a _game_, e.g. Jon
- *Team* - one or more (usually a pair of) _players_ playing on the same side in a _game_, e.g. Jon + Martin
- *Match* - a single matched played in a _game_ between two _teams_ with a _club_, e.g. Jon (Finland) vs. Martin (Germany): 2-2
- *Rating* - a rating showing how good a _team_ is compared to other _teams_, e.g. 4.37
- *Club* - sports club or international team within the _game_ playing in a _league_, e.g. Team Finland
- *Stars* - the number of stars a _club_ has in a _game_, e.g. 4.5/5
- *League* - sports league within the _game_ that _clubs_ are associated with, e.g. Bundesliga
- *Tournament* - a tournament created by FutuLeague in which three or more _teams_ compete against each other

