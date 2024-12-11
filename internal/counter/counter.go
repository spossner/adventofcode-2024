package counter

type Counter[T comparable] map[T]int

func NewCounter[T comparable](slice []T) Counter[T] {
	cnt := make(Counter[T])
	for _, el := range slice {
		cnt[el]++
	}
	return cnt
}

func (c *Counter[T]) Total() int {
	total := 0
	for _, v := range *c {
		total += v
	}
	return total
}
