package users

func IdsExcept(users []Id, except Id) []Id {
	ans := make([]Id, len(users)-1)
	i := 0
	for _, p := range users {
		if p != except {
			ans[i] = p
			i++
		}
	}
	return ans
}
