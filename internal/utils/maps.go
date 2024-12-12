package utils

func Contains[T comparable, V any](m map[T]V, key T) bool {
	if _, ok := m[key]; ok {
		return true
	}
	return false
}
