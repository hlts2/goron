package goron

// DefaultSpecCount represents the default spec count
const DefaultSpecCount = 5

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
		days    []Spec
		months  []Spec
		weeks   []Spec
	}
)

// NewSchedule returns *Schedule object
func NewSchedule() *Schedule {
	return &Schedule{
		minutes: make([]Spec, 0, DefaultSpecCount),
		hours:   make([]Spec, 0, DefaultSpecCount),
		days:    make([]Spec, 0, DefaultSpecCount),
		months:  make([]Spec, 0, DefaultSpecCount),
		weeks:   make([]Spec, 0, DefaultSpecCount),
	}
}

func parse(spec []string) (*Schedule, error) {
	return nil, nil
}
