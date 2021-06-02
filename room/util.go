package room

func UsersExcept(users []UserId, except UserId) []UserId {
	ans := make([]UserId, len(users)-1)
	i := 0
	for _, p := range users {
		if p != except {
			ans[i] = p
			i++
		}
	}
	return ans
}
