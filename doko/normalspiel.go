package doko

type NormalspielModus struct{}

func (n NormalspielModus) SpielFarbe(k Karte) SpielFarbe {
	return NormalspielFarbe(k)
}

func (n NormalspielModus) Sticht(neu Karte, alt Karte) bool {
	return NormalspielSticht(neu, alt)
}

func NormalspielSticht(neu Karte, alt Karte) bool {
	if IsNormalspielTrumpf(alt) {
		return NormalspielFarbe(neu) == Trumpf && NormalspielStichwert(neu) > NormalspielStichwert(alt)
	}
	if IsNormalspielTrumpf(neu) {
		return true
	}
	if NormalspielFarbe(neu) != NormalspielFarbe(alt) {
		return false
	}
	return NormalspielStichwert(neu) > NormalspielStichwert(alt)
}

func IsNormalspielTrumpf(k Karte) bool {
	return k.Farbe == Karo || k.rank == Bube || k.rank == Dame || k.IsDulle()
}

func NormalspielFarbe(k Karte) SpielFarbe {
	if IsNormalspielTrumpf(k) {
		return Trumpf
	}
	return k.Farbe.AlsFehl()
}

func NormalspielStichwert(card Karte) int {
	switch {
	case card.IsDulle():
		return 1000
	case card.rank == Dame:
		return 200 + int(card.Farbe)
	case card.rank == Bube:
		return 100 + int(card.Farbe)
	default:
		return card.Value()
	}
}

var Normalspiel = NormalspielModus{}
