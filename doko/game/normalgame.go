package game

func IsNormalTrump(k Card) bool {
	return k.Suit == Karo || k.Rank == Bube || k.Rank == Dame || k == Dulle()
}

func NormalGameSuit(c Card) GameSuit {
	if IsNormalTrump(c) {
		return Trumpf
	}
	return c.Suit.AsFehl()
}

type NormalGameMode struct {
	parties    [NumPlayers]Party
	silentSolo bool
}

func NewNormalGame(dealtCards Cards) NormalGameMode {
	var parties [NumPlayers]Party
	silentSolo := false
	for _, p := range Players() {
		alte := dealtCards[p].NumAlte()
		if alte >= 1 {
			parties[p] = ReParty
			if alte == 2 {
				silentSolo = true
			}
		} else {
			parties[p] = ContraParty
		}
	}
	return NormalGameMode{parties, silentSolo}
}

func (n NormalGameMode) IsSilentSolo() bool {
	return n.silentSolo
}

func (NormalGameMode) OnCompletedTrick(Trick, int) {
}

func (NormalGameMode) Klaerungsstich() int {
	return -1
}

func (n NormalGameMode) GameSuit(k Card) GameSuit {
	return NormalGameSuit(k)
}

func (n NormalGameMode) PartyOf(p Player) Party {
	return n.parties[p]
}

func (n NormalGameMode) Type() AnnouncedGameType {
	return NormalGameType
}
