package game

type Player int

const (
	Player1 Player = iota
	Player2
	Player3
	Player4
)

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
func (p Player) NthNext(i int) Player {
	r := p
	for i > 0 {
		r = r.NextPlayer()
		i -= 1
	}
	return r
}
