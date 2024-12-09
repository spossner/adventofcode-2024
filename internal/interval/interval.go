package interval

// Interval represents a range of number from From to To (not included).
type Interval struct {
	From, To int // To not included
}

func (i Interval) Len() int {
	return i.To - i.From
}
