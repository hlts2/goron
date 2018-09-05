package goron

// const value for date
const (
	Minute = iota
	Hour
	Day
	Month
	Week
)

type (
	// Spec is spec for goron
	Spec struct {
		min      int
		max      int
		wildcard string
	}

	// Schedule represents time schedule
	Schedule struct {
		minutes []Spec
		hours   []Spec
		day     []Spec
		Month   []Spec
		Week    []Spec
	}
)

func parse(spec []string) (Schedule, error) {
	return Schedule{}, nil
}
