# KarlchenCloud · [![CircleCI](https://circleci.com/gh/supermihi/karlchencloud.svg?style=shield)](https://circleci.com/gh/supermihi/karlchencloud)

KarlchenCloud is an open source client/server implementation of the German card game
[Doppelkopf](https://en.wikipedia.org/wiki/Doppelkopf).


We provide a web application that can be used without any installation.
Using the open and simple [API](api/README.md), however, other clients an be developed as well.
## Usage
Sorry - no working version available yet.

### Command line game with bots
You can debug the game with bots (that just play random cards).

#### Interactive Client (plus bots)
To start a game on the command line, you need three terminals and start in the
following order:
1. Server terminal:
   `CONSTANT_INVITE_CODE=1234 CONSTANT_TABLE_ID=5678 go run cmd/server/main.go`
1. Client terminal:
   `go run cmd/cli_client/main.go`
1. Input `c` to create a table.
1. Bot terminal:
   `INVITE_CODE=1234 TABLE_ID=5678 go run cmd/bot_client/main.go`
1. Play the game in the client terminal.

#### Bots Only
This is very similar, but only requires two terminals (in order):
1. Server terminal:
   `CONSTANT_INVITE_CODE=1234 CONSTANT_TABLE_ID=5678 go run cmd/server/main.go`
1. Bot terminal:
   `INVITE_CODE=1234 TABLE_ID=5678 INIT_TABLE=1 NUM_BOTS=4 go run cmd/bot_client/main.go`
1. "Enjoy" bots at work.


## Development
See [Development](Development.md).
