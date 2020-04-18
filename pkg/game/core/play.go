package core

type ActionType int

type PlayCardRequest struct {
	Player Player
	Card   Card
}

type PlayCardResult int

const (
	CardPlayed PlayCardResult = iota
	InvalidCard
	InvalidPlayer
	GameFinished
)

func (g Game) ProcessPlayCardRequest(r PlayCardRequest) PlayCardResult {
	if g.IsFinished() {
		return GameFinished
	}
	if r.Player != g.WhoseTurn() {
		return InvalidPlayer
	}
	if !g.TryPlayCard(r.Player, r.Card) {
		return InvalidCard
	}
	return CardPlayed
}
