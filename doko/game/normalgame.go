package game

func IsNormalTrump(k Card) bool {
	return k.Suit == Diamonds || k.Rank == Jack || k.Rank == Queen || k == TenOfHearts()
}

func NormalGameSuit(c Card) GameSuit {
	if IsNormalTrump(c) {
		return Trump
	}
	return c.Suit.AsNonTrump()
}

type NormalGameMode struct {
	parties    [NumPlayers]Party
	silentSolo bool
}

func NewNormalGame(dealtCards Cards) NormalGameMode {
	var parties [NumPlayers]Party
	silentSolo := false
	for _, p := range Players() {
		queensOfClubs := dealtCards[p].NumQueensOfClubs()
		if queensOfClubs >= 1 {
			parties[p] = ReParty
			if queensOfClubs == 2 {
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
