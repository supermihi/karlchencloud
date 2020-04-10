package doko

import "fmt"

type Farbe int

const (
	Karo Farbe = iota
	Herz
	Pik
	Kreuz
)

func (s Farbe) String() string {
	return [...]string{"Karo", "Herz", "Pik", "Kreuz"}[s]
}

type Symbol int

const (
	Neun Symbol = iota
	Bube
	Dame
	Koenig
	Zehn
	Ass
)

func (s Symbol) String() string {
	return [...]string{"Neun", "Bube", "Dame", "Koenig", "Zehn", "Ass"}[s]
}

func (s Symbol) Value() int {
	return [...]int{0, 2, 3, 4, 10, 11}[s]
}

type Karte struct {
	Farbe Farbe
	rank  Symbol
}

func (c Karte) IsDulle() bool {
	return c.Farbe == Herz && c.rank == Zehn
}

func (c Karte) String() string {
	return fmt.Sprintf("%s %s", c.Farbe, c.rank)
}

func (c Karte) Value() int {
	return c.rank.Value()
}

func Deck() []Karte {
	var ans [48]Karte
	pos := 0
	for suit := Karo; suit <= Kreuz; suit++ {
		for rank := Neun; rank <= Ass; rank++ {
			ans[pos] = Karte{suit, rank}
			ans[pos+1] = Karte{suit, rank}
			pos += 2
		}
	}
	return ans[:]
}
