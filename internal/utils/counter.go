package utils

func Counter[C comparable](slice []C) map[C]int {
	cnt := make(map[C]int)
	for _, el := range slice {
		cnt[el]++
	}
	return cnt
}
