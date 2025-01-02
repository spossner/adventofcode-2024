package interval

import (
	"github.com/spossner/aoc2024/internal/utils"
	"iter"
)

// Interval represents a range of number from From to To (not included).
type Interval struct {
	From, To int // To not included
}

// Len retrieves the length of the interval
func (i Interval) Len() int {
	return i.To - i.From
}

// Contains checks if the given value is within the interval
func (i Interval) Contains(value int) bool {
	return value >= i.From && value < i.To
}

// Overlaps checks whether or not this interval overlaps another one
func (i Interval) Overlaps(other Interval) bool {
	if i.From < other.From && i.To <= other.From {
		return false
	}

	if other.From < i.From && other.To <= i.From {
		return false
	}

	return true
}

// Intersect returns the Interval which is covered by this and the other interval
func (i Interval) Intersect(other Interval) Interval {
	if !i.Overlaps(other) {
		return Interval{}
	}

	_, from := utils.MinMax(i.From, other.From)
	to, _ := utils.MinMax(i.To, other.To)
	return Interval{from, to}
}

// All iterates all values in the interval
func (i Interval) All() iter.Seq[int] {
	return func(yield func(value int) bool) {
		for n := i.From; n < i.To; n++ {
			if !yield(n) {
				break
			}
		}
	}
}
