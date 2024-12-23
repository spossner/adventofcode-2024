package counter

type Counter[T comparable] map[T]int

func NewCounter[T comparable](items ...T) Counter[T] {
	cnt := make(Counter[T])
	for _, item := range items {
		cnt[item]++
	}
	return cnt
}

func (c Counter[T]) Add(item T) {
	c[item]++
}

func (c Counter[T]) Total() int {
	total := 0
	for _, v := range c {
		total += v
	}
	return total
}
