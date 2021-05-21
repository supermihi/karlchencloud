package server

import "math/rand"

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandomLetters(n int) string {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(res)
}

func usersExcept(users []UserId, except UserId) []UserId {
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

func getStringWithDefault(value *string, def func() string) string {
	if value != nil && *value != "" {
		return *value
	}
	return def()
}
