# KarlchenCloud · [![CircleCI](https://circleci.com/gh/supermihi/karlchencloud.svg?style=shield)](https://circleci.com/gh/supermihi/karlchencloud)

KarlchenCloud is an open source client/server implementation of the German card game
[Doppelkopf](https://en.wikipedia.org/wiki/Doppelkopf).


We provide a web application that can be used without any installation.
Using the open and simple [API](api/README.md), however, other clients an be developed as well.
## Usage
### Docker
- optional: [build the frontend](frontend/README.md)
- run `docker-compose build`
- run `docker-compose up`
- optional: if you built the frontend, visit http://localhost:5051 in a browser

### Command line game with bots
You can debug the game with bots (that just play random cards).

#### Interactive Client (plus bots)
To start a game on the command line, you need three terminals and start in the
following order:
1. Server terminal:
   `go run ./cmd/server`
1. Client terminal:
   `go run ./cmd/client --name "Karlchen Müller" --email karlchen@mueller.de --password 12345`
1. Follow the instructions
1. Bot terminal:
   `go run ./cmd/client bot --num-bots 3 --delay 2000`
1. Play the game in the client terminal.

#### Bots Only
This is very similar, but only requires two terminals (in order):
1. Server terminal:
   `go run ./cmd/server`
1. Bot terminal:
   `go run ./cmd/client bot --num-bots 3 --delay 250 --owner`
1. "Enjoy" bots at work.


## Development
See [Development](Development.md).
