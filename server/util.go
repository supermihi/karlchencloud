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

func stringsExcept(strings []string, except string) []string {
	ans := make([]string, len(strings)-1)
	i := 0
	for _, p := range strings {
		if p != except {
			ans[i] = p
			i++
		}
	}
	return ans
}
