package game

import "fmt"

type Player int

const (
	Player1 Player = iota
	Player2
	Player3
	Player4
	NoPlayer
)

func (p Player) String() string {
	return fmt.Sprintf("Player%d", int(p)+1)
}

func (p Player) NextPlayer() Player {
	switch p {
	case Player1:
		return Player2
	case Player2:
		return Player3
	case Player3:
		return Player4
	case Player4:
		return Player1
	default:
		panic("not a valid player")
	}
}

func Players() [NumPlayers]Player {
	return [...]Player{Player1, Player2, Player3, Player4}
}

// return players in order, starting with the given player
func PlayersFrom(player Player) [NumPlayers]Player {
	var ans [NumPlayers]Player
	for i := 0; i < NumPlayers; i++ {
		ans[i] = player.NthNext(i)
	}
	return ans
}
func (p Player) NthNext(i int) Player {
	r := p
	for i > 0 {
		r = r.NextPlayer()
		i -= 1
	}
	return r
}
