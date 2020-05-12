package match

func StandardSonderspiele() Sonderspiele {
	return MakeSonderspiele(append(AllFarbsolos(), VorbehaltHochzeit{}))
}
